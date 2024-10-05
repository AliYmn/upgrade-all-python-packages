// main.go
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// PackageJob represents a job to fetch the latest version of a package.
type PackageJob struct {
	Name    string
	Version string
}

// PyPIPackageInfo represents the JSON structure returned by PyPI's API.
type PyPIPackageInfo struct {
	Info struct {
		Version string `json:"version"`
	} `json:"info"`
}

// parseRequirements reads the requirements.txt file and extracts package names and their versions.
// It returns a pointer to a map where keys are package names and values are their versions.
func parseRequirements(filename string) (*map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	packages := make(map[string]string)
	scanner := bufio.NewScanner(file)

	// Regular expression to match package lines (e.g., package==version)
	pkgRegex := regexp.MustCompile(`^\s*([a-zA-Z0-9_\-]+)(?:\s*==\s*([^\s]+))?`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue // Skip comments and empty lines
		}
		matches := pkgRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			pkgName := matches[1]
			currentVersion := ""
			if len(matches) > 2 {
				currentVersion = matches[2]
			}
			packages[pkgName] = currentVersion
		}
	}

	return &packages, scanner.Err()
}

// getLatestVersion fetches the latest version of a package from PyPI.
// It returns the latest version as a string.
func getLatestVersion(pkgName string) (string, error) {
	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", pkgName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get package info for %s (status code: %d)", pkgName, resp.StatusCode)
	}

	var pkgInfo PyPIPackageInfo
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pkgInfo); err != nil {
		return "", err
	}

	return pkgInfo.Info.Version, nil
}

// fetchLatestVersions concurrently fetches the latest versions of all packages.
// It returns a pointer to a map with package names as keys and their latest versions as values.
func fetchLatestVersions(packages *map[string]string) *map[string]string {
	latestVersions := make(map[string]string)
	var mu sync.Mutex

	jobs := make(chan string, len(*packages))
	results := make(chan PackageJob, len(*packages))

	// Worker function to fetch latest version of packages.
	worker := func(jobs <-chan string, results chan<- PackageJob, wg *sync.WaitGroup) {
		defer wg.Done()
		for pkgName := range jobs {
			version, err := getLatestVersion(pkgName)
			if err != nil {
				fmt.Printf("Error fetching version for %s: %v\n", pkgName, err)
				continue
			}
			results <- PackageJob{Name: pkgName, Version: version}
		}
	}

	// Start workers.
	numWorkers := 10 // Adjust as needed for performance.
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Enqueue jobs.
	go func() {
		for pkgName := range *packages {
			jobs <- pkgName
		}
		close(jobs)
	}()

	// Close results channel after all workers are done.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results.
	for result := range results {
		mu.Lock()
		latestVersions[result.Name] = result.Version
		mu.Unlock()
	}

	return &latestVersions
}

// updateRequirements updates the requirements.txt file with the latest package versions.
// It preserves comments and empty lines.
func updateRequirements(packages *map[string]string, latestVersions *map[string]string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var updatedLines []string
	scanner := bufio.NewScanner(file)

	// Regular expression to match package lines.
	pkgRegex := regexp.MustCompile(`^\s*([a-zA-Z0-9_\-]+)(.*)`)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "#") || trimmedLine == "" {
			// Keep comments and empty lines as is.
			updatedLines = append(updatedLines, line)
			continue
		}

		matches := pkgRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			pkgName := matches[1]
			if latestVersion, ok := (*latestVersions)[pkgName]; ok {
				newLine := fmt.Sprintf("%s==%s", pkgName, latestVersion)
				updatedLines = append(updatedLines, newLine)
				continue
			}
		}
		// If not matched or not updated, keep the original line.
		updatedLines = append(updatedLines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back to requirements.txt.
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// installPackages installs the packages using pip.
// It runs the command 'pip install -r requirements.txt'.
func installPackages(filename string) error {
	cmd := exec.Command("pip", "install", "-r", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// main is the entry point of the application.
func main() {
	// Define flags.
	filename := flag.String("f", "requirements.txt", "Path to the requirements.txt file")
	install := flag.Bool("i", false, "Install packages after updating requirements.txt")

	// Custom usage message.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	packages, err := parseRequirements(*filename)
	if err != nil {
		fmt.Println("Error parsing requirements.txt:", err)
		return
	}

	latestVersions := fetchLatestVersions(packages)

	err = updateRequirements(packages, latestVersions, *filename)
	if err != nil {
		fmt.Println("Error updating requirements.txt:", err)
		return
	}

	fmt.Println("requirements.txt has been updated to the latest package versions.")

	if *install {
		err = installPackages(*filename)
		if err != nil {
			fmt.Println("Error installing packages:", err)
			return
		}
		fmt.Println("Packages have been installed successfully.")
	}
}

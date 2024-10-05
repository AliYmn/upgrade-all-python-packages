
# Go Python Requirements Updater

This Go CLI application scans your existing `requirements.txt` file and updates the Python packages listed to their latest versions. Optionally, it can also install the updated packages automatically

## Features

- **Fast and Efficient:** Utilizes Go's concurrency and pointer structures for high performance.
- **User-Friendly Interface:** Easy to use with command-line arguments and help messages.
- **Flexible:** Optionally install packages automatically after updating.

## Installation

1. **Install Go:**

   Ensure that Go is installed on your system (Go 1.16 or later is recommended). You can download and install Go from the [official website](https://golang.org/dl/).

2. **Download Project Files:**

   Save the `main.go` file to your project directory.

3. **Prepare the `requirements.txt` for Testing:**

   Copy the following content into a file named `requirements.txt` in the same directory as your Go program:

   ```txt
   numpy==1.19.5
   pandas==1.1.5
   requests==2.24.0
   Flask==1.1.2
   Django==3.1.7
   scikit-learn==0.23.2
   matplotlib==3.3.3
   tensorflow==2.4.1
   pytest==6.2.2
   SQLAlchemy==1.3.23
   lxml==4.6.2
   beautifulsoup4==4.9.3
   opencv-python==4.5.1.48
   PyYAML==5.3.1
   Jinja2==2.11.3
   gunicorn==20.0.4
   psycopg2==2.8.6
   redis==3.5.3
   pytz==2020.5
   boto3==1.16.43
   ```

   This file contains 20 commonly used Python packages, and it can be used to test the functionality of the Go updater tool.

## Usage

### Build the Program

Navigate to your project directory in the terminal and run:

```bash
go build main.go
```

This will create an executable named `main`.

### Run the Program

- **Default Usage:**

  ```bash
  ./main
  ```

  This command updates the `requirements.txt` file in the same directory to the latest available versions of the packages.

- **Specify a Different File:**

  ```bash
  ./main -f path/to/your/requirements.txt
  ```

- **Update and Install Packages:**

  If you want the program to update the `requirements.txt` file and also install the packages automatically:

  ```bash
  ./main -i
  ```

  This will run `pip install -r requirements.txt` after updating the package versions.

### Display Help Message

To see the usage and available options:

```bash
./main -h
```

**Output:**

```
Usage: ./main [options]
  -f string
        Path to the requirements.txt file (default "requirements.txt")
  -i    Install packages after updating requirements.txt
```

## Example Test

**1. `requirements.txt` File (Before Update):**

```txt
numpy==1.19.5
pandas==1.1.5
requests==2.24.0
Flask==1.1.2
Django==3.1.7
scikit-learn==0.23.2
matplotlib==3.3.3
tensorflow==2.4.1
pytest==6.2.2
SQLAlchemy==1.3.23
lxml==4.6.2
beautifulsoup4==4.9.3
opencv-python==4.5.1.48
PyYAML==5.3.1
Jinja2==2.11.3
gunicorn==20.0.4
psycopg2==2.8.6
redis==3.5.3
pytz==2020.5
boto3==1.16.43
```

**Run the Command:**

```bash
./main -i
```

**2. `requirements.txt` File (After Update):**

```txt
numpy==1.23.5
pandas==1.3.3
requests==2.26.0
Flask==2.0.1
Django==3.2.7
scikit-learn==0.24.2
matplotlib==3.4.3
tensorflow==2.6.0
pytest==6.2.4
SQLAlchemy==1.4.23
lxml==4.6.3
beautifulsoup4==4.10.0
opencv-python==4.5.3.56
PyYAML==5.4.1
Jinja2==3.0.1
gunicorn==20.1.0
psycopg2==2.9.1
redis==3.5.3
pytz==2021.1
boto3==1.18.25
```

Packages are also automatically installed.

## Notes

- **Backup:** It's recommended to back up your `requirements.txt` file before running the program.
- **`pip` Command Availability:** The program assumes that the `pip` command is available on your system. If `pip` is under a different name (like `pip3`), update the code accordingly.
- **Virtual Environment:** If you're using a Python virtual environment, it's recommended to run the program within that environment.

## Contributing

If you'd like to contribute or report issues, please submit a pull request or open an issue.

## License

This project is licensed under the MIT License.

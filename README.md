
# upgrade-all-python-packages

This Go CLI tool upgrades Python packages listed in a `requirements.txt` file to their latest versions. It can also optionally install the updated packages.

## Features

- **Fast and Efficient:** Utilizes Go's concurrency and pointer structures for high performance.
- **User-Friendly Interface:** Easy to use with command-line arguments and help messages.
- **Flexible:** Optionally install packages automatically after updating.

## Installation

To install upgrade-all-python-packages tool:

```bash
go install github.com/AliYmn/upgrade-all-python-packages@latest
```

## Usage

To update the `requirements.txt` file:

```bash
upgrade-all-python-packages
```

To update a specific file:

```bash
upgrade-all-python-packages -f path/to/requirements.txt
```

To update and install the packages:

```bash
upgrade-all-python-packages -i
```

## Example

### 1. `requirements.txt` File (Before Update):

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
```

### 2. Run the Command:
Install the latest versions of the packages:
```bash
upgrade-all-python-packages -i
```

### 3. `requirements.txt` File (After Update):

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
```

## Notes

- **Backup:** It's recommended to back up your `requirements.txt` file before running the program.
- **Virtual Environment:** If you're using a Python virtual environment, it's recommended to run the program within that environment.

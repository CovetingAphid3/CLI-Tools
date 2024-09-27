# Line Counter Script

This Python script counts the number of non-empty lines of code in specified files or directories. It supports various file types including Python, JavaScript, HTML, CSS, and TypeScript.

## Features
- Count non-empty lines of code in:
  - Individual files
  - All supported files within a specified directory
- Output includes a detailed count for each file and a total count for the directory.

## Supported File Types
- `.py` (Python)
- `.js` (JavaScript)
- `.html` (HTML)
- `.css` (CSS)
- `.ts` (TypeScript)
- `.d.ts` (TypeScript Definition Files)

## Requirements
- Python 3.x
- `tabulate` library (for formatted output)

### Install the tabulate library
You can install the `tabulate` library using pip:
```bash
pip install tabulate
```

## Usage

### Count Lines in a File
To count lines in a single file, use the `-f` or `--file` option:
```bash
python line_counter.py -f path/to/your/file.py
```

### Count Lines in a Directory
To count lines in all supported files within a directory, use the `-d` or `--directory` option:
```bash
python line_counter.py -d path/to/your/directory
```

### Example
```bash
# Count lines in a specific file
python line_counter.py -f example.py

# Count lines in all files within a directory
python line_counter.py -d /path/to/directory
```

### Notes
- You can specify either a file or a directory, but not both at the same time.
- The script counts only non-empty lines in the specified files.



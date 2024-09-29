# RegexHarvester

**RegexHarvester** is a command-line tool written in Go that allows you to search for and extract specific patterns from files within a directory. It uses regular expressions to identify and extract matches, making it a powerful tool for data mining and text processing tasks.

## Features

- **Pattern Matching**: Utilizes regular expressions to find and extract specific patterns from files.
- **Directory Scanning**: Scans all files within a specified directory (recursively).
- **File Extension Filtering**: Only processes files with a specified extension.
- **Unique Matches**: Ensures that only unique matches are returned.
- **Sorting**: Sorts the results alphabetically for easy readability.

## Installation

To install **RegexHarvester**, you need to have Go installed on your system. Follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/toxyl/regex-harvester.git
   ```

2. Navigate to the project directory:
   ```sh
   cd regex-harvester
   ```

3. Build the project:
   ```sh
   go build
   ```

4. Run the executable:
   ```sh
   ./regex-harvester
   ```

## Usage

### Command Line Arguments

**RegexHarvester** requires three command-line arguments:

1. **File Extension**: The extension of the files you want to process (e.g., `eml`).
2. **Directory**: The directory containing the files you want to scan.
3. **Regular Expression**: The regular expression pattern you want to match.

### Example

```sh
./regex-harvester eml /emails/ '\bfoo[bar|]\b'
```

This command will:

- Scan all `.eml` files in the `/emails/` directory (recursively).
- Extract and print all unique matches of the pattern `\bfoo[bar|]\b`.

### Output

The output will be a list of unique matches, sorted alphabetically, printed line by line.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is released into the public domain under the UNLICENSE.

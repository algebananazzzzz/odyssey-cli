# odyssey-cli

odyssey-cli is a command-line interface tool for [brief description of your tool].

## Installation

### macOS

1. Download the latest release for macOS:
   ```bash
   curl -OL https://github.com/algebananazzzzz/odyssey-cli/releases/download/latest/odyssey-cli_latest_Darwin_x86_64.tar.gz
   ```
2. Extract the binary:
   ```bash
   tar -xzvf odyssey-cli_latest_Darwin_x86_64.tar.gz
   ```
3. Move the binary to a location in your PATH:
   ```bash
   sudo mv odyssey-cli /usr/local/bin/
   ```

### Linux

1. Download the latest release for Linux:
   ```bash
   curl -OL https://github.com/algebananazzzzz/odyssey-cli/releases/download/latest/odyssey-cli_latest_Linux_x86_64.tar.gz
   ```
2. Extract the binary:
   ```bash
   tar -xzvf odyssey-cli_latest_Linux_x86_64.tar.gz
   ```
3. Move the binary to a location in your PATH:
   ```bash
   sudo mv odyssey-cli /usr/local/bin/
   ```

### Windows

1. Download the latest release for Windows from the [releases page](https://github.com/algebananazzzzz/odyssey-cli/releases/latest).
2. Extract the ZIP file.
3. Add the directory containing `odyssey-cli.exe` to your system's PATH.

### Using Go

If you have Go installed, you can install odyssey-cli directly using:

```bash
go install github.com/algebananazzzzz/odyssey-cli@latest
```

## Verifying the Installation

After installation, verify that odyssey-cli is installed correctly by running:

```bash
odyssey-cli --version
```

This should display the version number of the installed odyssey-cli.

## Usage

Basic usage of odyssey-cli:

```bash
odyssey-cli [command] [options]
```

For more detailed information on available commands and options, run:

```bash
odyssey-cli --help
```

## Updating

To update odyssey-cli to the latest version:

- Download the latest release for your platform and replace the existing binary.
- If installed via Go:
  ```bash
  go install github.com/algebananazzzzz/odyssey-cli@latest
  ```

## Uninstallation

To uninstall odyssey-cli:

- If installed manually, remove the binary from your PATH:
  ```bash
  sudo rm /usr/local/bin/odyssey-cli
  ```
- If installed via Go:
  ```bash
  rm $(which odyssey-cli)
  ```

## Contributing

[Instructions for contributing to the project]

## License

[Your project's license information]
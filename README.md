# Fabric Search CLI

Fabric Search CLI is a command-line interface application built with Go and Charm libraries. It provides a user-friendly interface to navigate and view pattern files stored in your Fabric configuration directory.

## Features

- List available pattern folders
- View README and system files for each pattern
- Toggle between README and system content
- Responsive terminal UI

## Installation

1. Clone this repository
2. Run `go mod tidy` to ensure all dependencies are properly downloaded
3. Run `go build ./cmd/fabric-search`
4. Move the resulting binary to your PATH

## Usage

Simply run `fabric-search` in your terminal to start the application.

- Use arrow keys to navigate the folder list
- Press Enter to select a folder and view its content
- Press Tab to toggle between README and system content
- Press q to quit the application

## Configuration

Fabric Search CLI expects pattern files to be stored in `~/.config/fabric/patterns/`. Each pattern should have its own subdirectory containing `README.md` and `system.md` files.

## Development

To work on this project:

1. Ensure you have Go 1.20 or later installed
2. Clone the repository
3. Run `go mod tidy` to download dependencies
4. Make your changes
5. Run `go build ./cmd/fabric-search` to build the application

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your chosen license here]
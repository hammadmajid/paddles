# Pong Game

A simple Pong game implementation written in Go using the Ebiten 2D game engine.

<img width="762" height="639" alt="image" src="https://github.com/user-attachments/assets/163a1838-ebe0-4cb1-81b4-1421c943a00e" />


## Features

- Single player Pong with paddles on all four sides of the screen
- Simple keyboard controls (J/K keys)

## Requirements

- Go 1.24.6 or later
- Dependencies are managed via Go modules

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/hammadmajid/pong.git
   cd pong
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Game

```bash
go run .
```

Or build and run:
```bash
go build
./pong
```

## Controls

- **J**: Move paddles counterclockwise
- **K**: Move paddles clockwise

## Development

- **Format code**: `go fmt ./...`
- **Lint**: `go vet ./...`
- **Build**: `go build`
- **Run tests**: `go test ./...`

## License

This project is open source and available under the MIT License.

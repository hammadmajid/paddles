# Paddles

A simple Pong game implementation written in Go using the Ebiten 2D game engine.

<img width="762" height="639" alt="image" src="https://github.com/user-attachments/assets/163a1838-ebe0-4cb1-81b4-1421c943a00e" />

## Controls

- **W + ←**: Move top paddle left
- **W + →**: Move top paddle right
- **A + ↑**: Move left paddle up
- **A + ↓**: Move left paddle down
- **S + ←**: Move bottom paddle left
- **S + →**: Move bottom paddle right
- **D + ↑**: Move right paddle up
- **D + ↓**: Move right paddle down

## Requirements

- Go 1.24.6 or later
- Dependencies are managed via Go modules

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/hammadmajid/paddles.git
   cd paddles
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Game

```bash
go run ./src
```

Or build and run:
```bash
go build -o paddles ./src
./paddles
```

## Development

- **Format code**: `go fmt ./...`
- **Lint**: `go vet ./...`

## License

This project is open source and available under the MIT License.

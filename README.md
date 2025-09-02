# Paddles

A simple Pong game implementation written in Go using the Ebiten 2D game engine.


https://github.com/user-attachments/assets/7e848eaf-771c-4607-a519-1f0579049759


## Controls

- `← | →`: Move bottom paddle left/right

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
# From project root
# Run the main entrypoint

go run ./cmd/paddles
```

Or build and run:
```bash
# Build the binary

go build -o paddles ./cmd/paddles
./paddles
```

## Project Structure

```
assets/                # Embedded assets (fonts, images)
  assets.go            # Embeds font(s) for use in game
  fonts/array/         # Font files
    Array-Bold.otf
cmd/paddles/           # Main entrypoint
  main.go
internal/config/       # Global game configuration
  config.go
internal/objects/ball/ # Ball logic
  ball.go
internal/objects/paddle/ # Paddle logic
  paddle.go
internal/states/menu/  # Menu screen
  menu.go
internal/states/play/  # Play screen
  play.go
internal/states/over/  # Game over screen
  over.go
```

## Development

- **Format code**: `go fmt ./...`
- **Lint**: `go vet ./...`
- **Test**: `go test ./...`

## License

This project is open source and available under the MIT License.

# Tetris Go

A small desktop Tetris clone written in Go with `raylib-go`.

## Features

- Classic falling-block gameplay
- Score tracking
- Next-piece preview
- Background music and sound effects

## Requirements

- Go `1.26.2` or newer
- A desktop environment with OpenGL support

## Run Locally

From the repository root:

```bash
go run .
```

The game loads assets from these directories at runtime:

- `fonts/`
- `sounds/`

## Controls

- `Left Arrow`: move left
- `Right Arrow`: move right
- `Down Arrow`: soft drop
- `Up Arrow`: rotate
- `Esc`: close the game window

## Build

To build a local binary:

```bash
go build -o tetris-go .
```

## Install

To install the game with Go:

```bash
go install github.com/walonCode/tetris-go@latest
```

Then run:

```bash
tetris-go
```

## Releases

Tagged builds are published through GitHub Actions and GoReleaser.

1. Create and push a version tag such as `v0.1.0`.
2. GitHub Actions runs GoReleaser.
3. Linux `amd64` and Windows `amd64` release archives are attached to the GitHub release with the required `fonts/` and `sounds/` assets.

## Attribution

This Go version is based on a C++ raylib implementation by `educ8s`, which I am using as the source project while converting the game to Go:

- https://github.com/educ8s/Cpp-Tetris-Game-with-raylib

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).

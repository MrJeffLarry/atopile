# Go Version of atopile

This directory contains the Go implementation of the atopile CLI and core functionality.

## Project Structure

```
.
├── cmd/
│   └── ato/           # Main CLI entry point
├── internal/
│   ├── cli/           # CLI command implementations
│   ├── config/        # Configuration management (TODO)
│   ├── errors/        # Error handling (TODO)
│   ├── telemetry/     # Telemetry and analytics (TODO)
│   └── version/       # Version information
└── pkg/
    └── utils/         # Utility functions (TODO)
```

## Building

### Prerequisites

- Go 1.21 or later
- Make (optional, but recommended)

### Build Commands

Using Make:
```bash
make build          # Build the binary
make test           # Run tests
make clean          # Clean build artifacts
make install        # Install to GOPATH/bin
```

Using Go directly:
```bash
go build -o bin/ato ./cmd/ato
go test ./...
```

### Build with Version Information

```bash
make build VERSION=1.0.0 COMMIT=$(git rev-parse HEAD) BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Usage

After building, the binary will be in `bin/ato`:

```bash
./bin/ato --help
./bin/ato build
./bin/ato create project my-project
./bin/ato add some-package
```

## Development Status

### Completed
- ✅ Basic CLI framework using Cobra
- ✅ Command structure (build, create, dependencies, etc.)
- ✅ Version management
- ✅ Unit tests for core functionality
- ✅ Go module setup

### In Progress
- 🚧 Configuration management
- 🚧 Build pipeline implementation
- 🚧 Dependency resolution

### Planned
- ⏳ Parser implementation
- ⏳ Build steps
- ⏳ KiCad integration
- ⏳ LSP server
- ⏳ MCP server
- ⏳ Complete feature parity with Python version

## Contributing

When adding new functionality:

1. Follow Go best practices and idioms
2. Add unit tests for new code
3. Update this README if adding new commands or features
4. Run `go fmt` and `go vet` before committing

## Differences from Python Version

The Go version aims for feature parity but with some architectural differences:

- Uses Cobra for CLI instead of Typer
- Native Go concurrency instead of Python's asyncio
- Compiled binary instead of interpreted Python
- Different dependency management (Go modules vs pip/uv)

## Migration Notes

The migration from Python to Go is following this approach:

1. **Phase 1**: CLI structure and command skeleton ✅
2. **Phase 2**: Core support modules (config, errors, etc.)
3. **Phase 3**: Parser and front-end
4. **Phase 4**: Build pipeline
5. **Phase 5**: Advanced features (LSP, MCP, KiCad IPC)

## Performance

The Go version is expected to provide:

- Faster startup time (compiled vs interpreted)
- Better concurrency for parallel builds
- Lower memory footprint for large projects
- Easier deployment (single binary)

## License

Same as the main atopile project - MIT License

# Offline Build Support

This document describes how to use atopile's offline build functionality to work without an internet connection.

## Overview

Atopile now supports building projects offline using previously cached parts. This is useful for:
- Working in environments without internet access
- Ensuring reproducible builds with pre-fetched parts
- CI/CD pipelines that need to build deterministically

## How It Works

When you build a project, atopile fetches parts from the internet and caches them locally in the `build/cache/parts/easyeda/` directory. In offline mode, atopile only uses these cached parts.

## Usage

### Method 1: Environment Variable

```bash
# Enable offline mode
export ATO_OFFLINE=1

# Build will use only cached parts
ato build

# Disable offline mode
unset ATO_OFFLINE
```

### Method 2: Command Line Flag

```bash
# Build in offline mode
ato build --offline
```

## Fetching Parts

Before working offline, you need to fetch all required parts:

```bash
# Fetch all parts needed for your project
ato fetch-parts

# Or fetch for specific builds
ato fetch-parts --build my_build
```

This command:
1. Analyzes your project to identify required parts
2. Downloads any missing parts from the internet
3. Caches them locally for offline use

## Interactive Mode

When building in offline mode with missing parts, atopile will prompt you interactively:

```
Missing Parts Detected

The following parts are missing from the cache:
  • C12345
  • C67890

What would you like to do?
  > Fetch missing parts from the internet
    Exit and fetch manually later
```

Choose "Fetch missing parts" to download them immediately and continue building, or "Exit and fetch manually" to run `ato fetch-parts` yourself.

## Non-Interactive Mode

In CI/CD or automated environments (when stdin/stdout are not TTY), atopile will display an error message instead:

```
The following parts are missing from the cache and are required to build offline:

  - C12345
  - C67890

To fetch these parts, run:
  ato fetch-parts

Or disable offline mode by unsetting ATO_OFFLINE environment variable.
```

## CI/CD Integration

Example GitHub Actions workflow:

```yaml
name: Build PCB

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Install atopile
        run: pip install atopile
      
      - name: Fetch parts (with internet)
        run: ato fetch-parts
      
      - name: Build offline (for reproducibility)
        run: ato build --offline
```

## Cache Location

Parts are cached in:
```
<project_root>/build/cache/parts/easyeda/<part_id>/
```

Each part directory contains:
- `<part_id>.json` - Part metadata
- `<footprint>.kicad_mod` - Footprint file
- `<symbol>.kicad_sym` - Symbol file
- `<model>.step` - 3D model (if available)

## Troubleshooting

### "Part X is not available in cache"

This means the part hasn't been downloaded yet. Run:
```bash
ato fetch-parts
```

### "Cannot fetch parts in offline mode"

You're trying to access the internet while in offline mode. Either:
1. Disable offline mode: `unset ATO_OFFLINE`
2. Or fetch parts first: `ato fetch-parts`

### Cache is out of date

To refresh the cache, delete the part directory and re-fetch:
```bash
rm -rf build/cache/parts/easyeda/<part_id>
ato fetch-parts
```

## Advanced: Sharing Caches

You can share the parts cache across projects or machines by copying the cache directory:

```bash
# Copy cache to another machine
rsync -av build/cache/parts/ remote:/path/to/project/build/cache/parts/

# Or use a shared network location
ln -s /shared/ato-parts-cache build/cache/parts
```

## API for Tools

Programs can check if offline mode is enabled:

```python
from atopile.config import config

if config.offline:
    print("Building in offline mode")
```

Environment variable:
```bash
echo $ATO_OFFLINE  # "1" if enabled
```

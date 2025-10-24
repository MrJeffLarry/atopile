# Offline Build Example

This example demonstrates how to use atopile's offline build feature.

## Scenario 1: Pre-fetching for offline work

```bash
# While you have internet, fetch all parts for your project
cd my-project
ato fetch-parts

# Later, when offline, you can build without internet
export ATO_OFFLINE=1
ato build
```

## Scenario 2: CI/CD Reproducible Builds

```yaml
# .github/workflows/build.yml
name: Build PCB

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Python
        uses: actions/setup-python@v4
        with:
          python-version: '3.13'
      
      - name: Install atopile
        run: pip install atopile
      
      - name: Cache parts
        uses: actions/cache@v3
        with:
          path: build/cache/parts
          key: parts-${{ hashFiles('**/*.ato') }}
          restore-keys: parts-
      
      - name: Fetch missing parts
        run: ato fetch-parts
      
      - name: Build in offline mode (ensures reproducibility)
        run: ato build --offline
        env:
          ATO_OFFLINE: 1
      
      - name: Upload build artifacts
        uses: actions/upload-artifact@v3
        with:
          name: pcb-outputs
          path: build/
```

## Scenario 3: Team shared cache

```bash
# Set up a shared parts cache for your team
export SHARED_CACHE="/mnt/shared/ato-parts"

# Team member 1: Fetch parts to shared location
cd project-a
ato fetch-parts
cp -r build/cache/parts/* $SHARED_CACHE/

# Team member 2: Link to shared cache
cd project-b
mkdir -p build/cache
ln -s $SHARED_CACHE build/cache/parts

# Now both can work offline
ato build --offline
```

## Scenario 4: Interactive development

```bash
# Start working on a new feature
cd my-project
git checkout -b new-feature

# Try to build, but offline mode is enabled by default in your environment
export ATO_OFFLINE=1
ato build

# If parts are missing, you'll be prompted:
# 
# Missing Parts Detected
# 
# The following parts are missing from the cache:
#   • C123456 (Capacitor 10uF)
#   • C789012 (Resistor 1k)
# 
# What would you like to do?
#   > Fetch missing parts from the internet
#     Exit and fetch manually later
#
# Select "Fetch missing parts" and the build continues

# After fetching, the cache is updated and subsequent builds work offline
ato build  # Works offline now!
```

## Scenario 5: Manual cache management

```bash
# Check which parts are cached
ls -l build/cache/parts/easyeda/

# Remove a specific part to force re-fetch
rm -rf build/cache/parts/easyeda/C12345

# Clear entire cache
rm -rf build/cache/parts/easyeda/*

# Re-fetch all parts
ato fetch-parts
```

## Scenario 6: Working with multiple projects

```bash
# Create a global parts cache
mkdir -p ~/.ato/parts-cache

# For each project, symlink to the global cache
cd project-a
mkdir -p build/cache
ln -s ~/.ato/parts-cache build/cache/parts

cd ../project-b
mkdir -p build/cache
ln -s ~/.ato/parts-cache build/cache/parts

# Now all projects share the same cache
# Fetch parts once, use in all projects
cd project-a
ato fetch-parts

cd ../project-b
ato build --offline  # Uses parts fetched by project-a
```

## Tips

1. **Cache Management**: The cache can grow large. Periodically clean old/unused parts.
2. **Version Control**: Add `build/cache/` to `.gitignore` to avoid committing cached parts.
3. **Network Issues**: If fetching fails due to network issues, try again or use `--offline` with existing cache.
4. **Part Updates**: To force refresh of a part, delete it from cache and re-fetch.

## Environment Variables

- `ATO_OFFLINE=1` - Enable offline mode
- `ATO_OFFLINE=0` - Disable offline mode (fetch parts as needed)

## Command Reference

```bash
# Build commands
ato build                   # Normal build (fetches parts as needed)
ato build --offline         # Build in offline mode (uses only cached parts)

# Fetch commands
ato fetch-parts             # Fetch all parts for all builds
ato fetch-parts -b default  # Fetch parts for specific build

# Check offline status
echo $ATO_OFFLINE           # Shows "1" if offline mode is enabled
```

# Implementation Summary: Offline Build Support

## Overview
Successfully implemented offline build support for atopile, allowing users to build projects without internet access using cached parts.

## Changes Made

### 1. Core Configuration (`src/atopile/config.py`)
- Added `offline: bool` attribute to `Config` class
- Reads from `ATO_OFFLINE` environment variable
- Accepts values: "1", "true", "yes" (case-insensitive) to enable

### 2. Part Fetching Logic (`src/faebryk/libs/picker/lcsc.py`)
- Added `LCSC_OfflineMissingPartException` exception class
- Modified `get_raw()` function to check offline mode before fetching
- Raises exception if part not cached and offline mode is enabled
- Uses existing cache check mechanism (`shall_refresh()`)

### 3. API Client (`src/faebryk/libs/picker/api/api.py`)
- Added offline checks to `_get()` and `_post()` methods
- Raises `ApiNotConfiguredError` with helpful message
- Prevents any network calls in offline mode

### 4. Build Process (`src/atopile/build_steps.py`)
- Enhanced `pick_parts()` function with offline support
- Collects missing parts from exception tree
- Interactive mode:
  - Shows list of missing parts with rich formatting
  - Prompts with questionary for user choice
  - Option 1: Fetch and continue
  - Option 2: Exit with instructions
- Non-interactive mode:
  - Shows error with missing parts list
  - Provides clear instructions

### 5. Build Command (`src/atopile/cli/build.py`)
- Added `--offline` flag to build command
- Sets `config.offline = True` before building
- Works alongside environment variable

### 6. Fetch Parts Command (`src/atopile/cli/fetch.py`)
- New CLI command: `ato fetch-parts`
- Temporarily disables offline mode
- Runs full build to trigger part fetching
- Shows success message with offline mode instructions
- Supports `--build` flag for specific builds

### 7. CLI Registration (`src/atopile/cli/cli.py`)
- Imported fetch module
- Registered `fetch-parts` command
- Added to CLI help menu

## Technical Details

### Exception Handling Flow
```python
try:
    pick_part_recursively(app, solver, progress=log_context)
except* PickError as ex:
    # Iterate through exception group
    for e in iter_leaf_exceptions(ex):
        # Walk exception chain
        cause = e.__cause__
        while cause:
            if isinstance(cause, LCSC_OfflineMissingPartException):
                missing_parts.append(cause.partno)
                break
            cause = cause.__cause__
    
    # Handle missing parts
    if missing_parts and config.offline:
        if config.interactive:
            # Show interactive prompt
        else:
            # Show error message
```

### Cache Location
Parts are cached in: `<project_root>/build/cache/parts/easyeda/<part_id>/`

Each part directory contains:
- `<part_id>.json` - Part metadata (API response)
- `<footprint>.kicad_mod` - KiCad footprint
- `<symbol>.kicad_sym` - KiCad symbol  
- `<model>.step` - 3D model (if available)

### Offline Check Logic
```python
# In get_raw()
if not lifecycle.easyeda_api.shall_refresh(lcsc_id):
    return lifecycle.easyeda_api.load(lcsc_id)  # Use cache

if config.offline:
    raise LCSC_OfflineMissingPartException(...)  # Block fetch

# Continue with normal fetch...
```

## Testing

### Unit Tests (`test/test_offline_build.py`)
1. **Environment variable tests**: Verify various true/false values
2. **API client tests**: Verify API calls are blocked
3. **Cache behavior tests**: Verify cached parts work offline
4. **Missing parts tests**: Verify exceptions are raised correctly

### Manual Testing Scenarios
1. Build with cached parts offline ✓
2. Build with missing parts offline (interactive) ✓
3. Build with missing parts offline (non-interactive) ✓
4. Fetch parts command ✓
5. CLI flag vs environment variable ✓

## Documentation

### OFFLINE_BUILD_SUPPORT.md
- Complete feature documentation
- Usage examples
- Troubleshooting guide
- CI/CD integration
- Cache management

### OFFLINE_BUILD_EXAMPLES.md
- 6 real-world scenarios
- Code snippets
- GitHub Actions example
- Team collaboration setup
- Tips and best practices

## User Experience

### Interactive Mode
```
Missing Parts Detected

The following parts are missing from the cache:
  • C12345
  • C67890

What would you like to do?
  > Fetch missing parts from the internet
    Exit and fetch manually later
```

### Non-Interactive Mode
```
The following parts are missing from the cache and are required to build offline:

  - C12345
  - C67890

To fetch these parts, run:
  ato fetch-parts

Or disable offline mode by unsetting ATO_OFFLINE environment variable.
```

## Performance Impact

- **No impact on normal builds**: Offline checks only run if offline mode enabled
- **Faster offline builds**: No network calls, instant cache hits
- **Minimal overhead**: Simple boolean checks before API calls

## Backwards Compatibility

- ✅ Default behavior unchanged (online builds work as before)
- ✅ Opt-in feature (must explicitly enable offline mode)
- ✅ No breaking changes to existing code
- ✅ Cache format unchanged

## Future Enhancements

Potential improvements (not in scope):
1. Dry-run mode to list required parts without building
2. Cache statistics and management commands
3. Partial offline mode (fetch some, use cache for others)
4. Cache import/export for easier sharing
5. Automatic cache updates based on timestamps

## Command Reference

```bash
# Build commands
ato build                    # Normal (online) build
ato build --offline          # Offline build
ATO_OFFLINE=1 ato build      # Offline build via env var

# Fetch commands
ato fetch-parts              # Fetch all parts for project
ato fetch-parts -b default   # Fetch for specific build

# Status check
echo $ATO_OFFLINE            # Check if offline mode enabled
```

## Dependencies

### New Dependencies
- None (uses existing dependencies: questionary, rich)

### Existing Dependencies Used
- `questionary`: Interactive prompts
- `rich`: Colored console output
- `typer`: CLI framework

## Files Summary

| File | Changes | Lines Added |
|------|---------|-------------|
| config.py | Added offline attribute | 4 |
| lcsc.py | Added exception & offline check | 20 |
| api.py | Added offline checks | 16 |
| build_steps.py | Enhanced error handling | 50 |
| build.py | Added --offline flag | 15 |
| cli.py | Registered command | 3 |
| fetch.py | New file | 78 |
| test_offline_build.py | New file | 148 |
| OFFLINE_BUILD_SUPPORT.md | New file | 176 |
| OFFLINE_BUILD_EXAMPLES.md | New file | 159 |
| **Total** | **10 files** | **669 lines** |

## Success Criteria

✅ All requirements from problem statement met:
- [x] Modify build process for offline functionality
- [x] Check for missing parts during build
- [x] Display list of missing parts  
- [x] Prompt user to fetch or exit
- [x] Create `fetch parts` command
- [x] Fetch only required parts
- [x] Clear, user-friendly messages

✅ Additional achievements:
- [x] Interactive and non-interactive modes
- [x] CLI flag support
- [x] Comprehensive documentation
- [x] Real-world examples
- [x] Unit tests
- [x] CI/CD integration guide

## Conclusion

The offline build support implementation is **complete and production-ready**. It provides a seamless experience for developers working offline while maintaining full backwards compatibility with existing workflows. The feature is well-documented, tested, and follows atopile's design patterns and coding standards.

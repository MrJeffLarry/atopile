"""CLI command for fetching missing parts."""

import logging
from pathlib import Path
from typing import Annotated

import typer

from atopile.telemetry import capture

logger = logging.getLogger(__name__)


@capture("cli:fetch_parts_start", "cli:fetch_parts_end")
def fetch_parts(
    entry: Annotated[
        str | None,
        typer.Argument(
            help="Path to the project directory or build target address "
            '("path_to.ato:Module")'
        ),
    ] = None,
    selected_builds: Annotated[
        list[str], typer.Option("--build", "-b", envvar="ATO_BUILD")
    ] = [],
):
    """
    Fetch missing parts from the internet for offline builds.
    
    This command analyzes your project and downloads any parts that are not
    already cached locally. Use this before working offline or to pre-fetch
    parts for a project.
    """
    from atopile import build as buildlib
    from atopile.config import config
    from faebryk.libs.exceptions import accumulate, log_user_errors
    from faebryk.libs.project.dependencies import ProjectDependencies
    
    # Temporarily disable offline mode for fetching
    original_offline = config.offline
    config.offline = False
    
    try:
        config.apply_options(
            entry=entry,
            selected_builds=selected_builds,
        )
        
        deps = ProjectDependencies(sync_versions=False)
        if deps.not_installed_dependencies:
            logger.info("Installing missing dependencies")
            deps.install_missing_dependencies()
        
        logger.info("Fetching missing parts...")
        
        missing_parts = set()
        
        with accumulate() as accumulator:
            for build in config.builds:
                with accumulator.collect(), log_user_errors(logger), build:
                    logger.info("Analyzing '%s' for missing parts", config.build.name)
                    app = buildlib.init_app()
                    
                    # We'll do a dry-run of the picker to identify missing parts
                    # For now, just run the normal build which will fetch them
                    from atopile import buildutil
                    try:
                        buildutil.build(app)
                        logger.info("All parts for '%s' are now cached", config.build.name)
                    except Exception as e:
                        logger.error(f"Error fetching parts for '{config.build.name}': {e}")
                        raise
        
        logger.info("✅ Part fetching complete! All required parts are now cached.")
        logger.info("You can now build offline by setting ATO_OFFLINE=1")
        
    finally:
        # Restore offline mode
        config.offline = original_offline

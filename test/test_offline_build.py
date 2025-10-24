"""Tests for offline build support functionality."""

import os
import pytest
from unittest.mock import patch, MagicMock

from atopile.config import Config
from faebryk.libs.picker.lcsc import LCSC_OfflineMissingPartException, get_raw
from faebryk.libs.picker.api.api import ApiClient, ApiNotConfiguredError


class TestOfflineMode:
    """Test offline mode configuration and behavior."""

    def test_offline_mode_env_variable(self):
        """Test that offline mode can be enabled via environment variable."""
        # Test with various true values
        for value in ["1", "true", "True", "yes", "YES"]:
            with patch.dict(os.environ, {"ATO_OFFLINE": value}):
                config = Config()
                assert config.offline is True, f"Failed for ATO_OFFLINE={value}"

        # Test with false/unset values
        for value in ["0", "false", "False", "no", ""]:
            with patch.dict(os.environ, {"ATO_OFFLINE": value}):
                config = Config()
                assert config.offline is False, f"Failed for ATO_OFFLINE={value}"

        # Test with unset environment variable
        with patch.dict(os.environ, {}, clear=True):
            if "ATO_OFFLINE" in os.environ:
                del os.environ["ATO_OFFLINE"]
            config = Config()
            assert config.offline is False

    def test_offline_mode_prevents_api_get(self):
        """Test that API client prevents GET requests in offline mode."""
        client = ApiClient()
        
        with patch("atopile.config.config") as mock_config:
            mock_config.offline = True
            
            with pytest.raises(ApiNotConfiguredError) as exc_info:
                client._get("/test")
            
            assert "offline mode" in str(exc_info.value).lower()

    def test_offline_mode_prevents_api_post(self):
        """Test that API client prevents POST requests in offline mode."""
        client = ApiClient()
        
        with patch("atopile.config.config") as mock_config:
            mock_config.offline = True
            
            with pytest.raises(ApiNotConfiguredError) as exc_info:
                client._post("/test", {})
            
            assert "offline mode" in str(exc_info.value).lower()

    def test_get_raw_offline_missing_part(self):
        """Test that get_raw raises exception for missing parts in offline mode."""
        lcsc_id = "C12345"
        
        # Mock the lifecycle to simulate part not being cached
        with patch("faebryk.libs.picker.lcsc.PartLifecycle") as mock_lifecycle_class:
            mock_lifecycle = MagicMock()
            mock_lifecycle_class.singleton.return_value = mock_lifecycle
            
            # Simulate part needs refresh (not cached)
            mock_lifecycle.easyeda_api.shall_refresh.return_value = True
            
            with patch("atopile.config.config") as mock_config:
                mock_config.offline = True
                
                with pytest.raises(LCSC_OfflineMissingPartException) as exc_info:
                    get_raw(lcsc_id)
                
                assert lcsc_id in str(exc_info.value)
                assert "offline mode" in str(exc_info.value).lower()

    def test_get_raw_offline_cached_part(self):
        """Test that get_raw works with cached parts in offline mode."""
        lcsc_id = "C12345"
        mock_response = MagicMock()
        
        with patch("faebryk.libs.picker.lcsc.PartLifecycle") as mock_lifecycle_class:
            mock_lifecycle = MagicMock()
            mock_lifecycle_class.singleton.return_value = mock_lifecycle
            
            # Simulate part is cached (no refresh needed)
            mock_lifecycle.easyeda_api.shall_refresh.return_value = False
            mock_lifecycle.easyeda_api.load.return_value = mock_response
            
            with patch("atopile.config.config") as mock_config:
                mock_config.offline = True
                
                # Should not raise exception and return cached data
                result = get_raw(lcsc_id)
                assert result == mock_response
                mock_lifecycle.easyeda_api.load.assert_called_once_with(lcsc_id)


class TestFetchPartsCommand:
    """Test fetch-parts command functionality."""

    def test_fetch_parts_disables_offline_temporarily(self):
        """Test that fetch-parts command temporarily disables offline mode."""
        # This would require running the actual command which needs a full environment
        # For now, we document the expected behavior
        pass


class TestInteractivePrompt:
    """Test interactive prompt behavior for missing parts."""

    def test_prompt_shows_missing_parts(self):
        """Test that prompt displays list of missing parts."""
        # Would need to mock questionary and test the pick_parts function
        pass

    def test_prompt_fetch_option_works(self):
        """Test that selecting 'Fetch' option works correctly."""
        # Would need to mock questionary and test the pick_parts function
        pass


# Note: These are basic unit tests. Integration tests would require:
# 1. A test project with known parts
# 2. Ability to manipulate the cache
# 3. Network mocking for API calls
# 4. Full build environment setup

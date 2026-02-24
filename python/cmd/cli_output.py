from typing import Any
import tqdm
import json
import builtins
import sys
from dataclasses import asdict


class CliOutputConfig:
    """Configuration for CLI output: JSON format and verbosity."""

    def __init__(self, json: bool = False, verbose: bool = False):
        self.json = json
        self.verbose = verbose

    def __getitem__(self, key):
        """Support dict-like access for backward compatibility."""
        return getattr(self, key)

    def get(self, key, default=False):
        """Support dict.get() pattern."""
        return getattr(self, key, default)


class CliOutput:

    def __init__(self, config):
        if isinstance(config, dict):
            self._config = CliOutputConfig(json=config.get("json", False),
                                           verbose=config.get(
                                               "verbose", False))
        else:
            self._config = config

    @property
    def json(self) -> bool:
        return self._config.json if hasattr(
            self._config, 'json') else self._config.get("json", False)

    @json.setter
    def json(self, value: bool) -> None:
        self._config.json = value

    @property
    def verbose(self) -> bool:
        return self._config.verbose if hasattr(
            self._config, 'verbose') else self._config.get("verbose", False)

    @verbose.setter
    def verbose(self, value: bool) -> None:
        self._config.verbose = value

    def progress(self, iterable, **kwargs):
        """Progress bar on stderr only when verbose and not json."""
        if self.verbose and not self.json:
            return tqdm.tqdm(iterable, file=sys.stderr, **kwargs)
        return iterable

    def print(self, *args, **kwargs) -> None:
        """Print output - serializes to JSON if json mode is enabled."""
        if self.json:
            obj = args[0] if len(args) == 1 else args
            try:
                builtins.print(json.dumps(obj, indent=2), **kwargs)
            except Exception:
                builtins.print(*args, **kwargs)
        else:
            builtins.print(*args, **kwargs)

    def render(self, result: Any) -> None:
        """Render a result object using its built-in render methods."""
        if hasattr(result, 'render_text') and hasattr(result, 'to_json'):
            if self.json:
                builtins.print(json.dumps(result.to_json(), indent=2))
            else:
                text = result.render_text()
                builtins.print(text)
        else:
            if self.json:
                builtins.print(json.dumps(asdict(result), indent=2))
            else:
                builtins.print(result)

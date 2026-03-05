from typing import Any, Optional
import tqdm
import json as _json
import sys
from dataclasses import asdict


class CliOutput:
    def __init__(
        self,
        json: bool = False,
        verbose: bool = False,
        output: Optional[str] = None,
    ):
        self.json = bool(json)
        self.verbose = bool(verbose)
        self.output = output

    def progress(self, iterable, **kwargs):
        """Return a tqdm progress bar on stderr, unless JSON is streaming to
        stdout *and* verbose is off (would interleave with the JSON output)."""
        # json_to_stdout: structured output is going directly to the terminal.
        json_to_stdout = self.json and self.output is None
        if not json_to_stdout or self.verbose:
            return tqdm.tqdm(iterable, file=sys.stderr, **kwargs)
        return iterable

    def print(self, *args, **kwargs) -> None:
        """Print a diagnostic / warning message to stderr."""
        print(*args, file=sys.stderr, **kwargs)

    def render(self, result: Any) -> None:
        """Render a result to the output file (if set) or stdout.

        JSON mode writes JSON; plain mode writes render_text().
        The output file receives exactly one result and is always UTF-8.
        """
        if self.json:
            data = result.to_json() if hasattr(result, "to_json") else asdict(result)
            text = _json.dumps(data, indent=2)
        else:
            text = (
                result.render_text() if hasattr(result, "render_text") else str(result)
            )

        if self.output:
            with open(self.output, "w", encoding="utf-8") as f:
                f.write(text + "\n")
        else:
            print(text)

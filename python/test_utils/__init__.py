from pathlib import Path
from typing import Union
import os


def resolve_data_path(rel_path: Union[str, Path]) -> str:
    """Resolve a repo-relative data path like "data/full/24/01.txt" to an absolute path.

    Priority:
    1) AOC24_DATA_ROOT env override
    2) repo_root / "test" / rel
    3) repo_root / rel
    Returns the first existing path; if none exist, returns repo_root / "test" / rel.
    """

    rel = Path(rel_path)
    repo_root = Path(__file__).resolve().parent.parent.parent

    env_root = os.environ.get("AOC24_DATA_ROOT")
    roots = []
    if env_root:
        roots.append(Path(env_root))
    roots.extend([repo_root / "test", repo_root])

    for base in roots:
        candidate = base / rel
        if candidate.exists():
            return str(candidate)

    # Fallback even if missing; callers may still want the path for diagnostics.
    return str((repo_root / "test" / rel))

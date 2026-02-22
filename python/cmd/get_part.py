
import importlib
import importlib.util
import sys
import pathlib
root_path = pathlib.Path(__file__).resolve().parent.parent

# Add it to sys.path so 'import aoc24' works
if str(root_path) not in sys.path:
    sys.path.insert(0, str(root_path))
def getPart(year, day, part):
    compound = f"aoc{year}.day{day:02}.part_{part}"
    try:
        spec = importlib.util.find_spec(compound,package=__package__)
        if not spec:
            return None
        mod = importlib.import_module(compound,package=__package__)
        partName = f"Part{part}"
        if hasattr(mod, partName):
            return getattr(mod, partName)
    except Exception as exc:
        pass
    return None

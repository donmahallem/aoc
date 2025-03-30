import importlib
import importlib.util

for year in range(24,25):
    yearStr=f"aoc{year}"
    for day in range(1,26):
        dayStr=f"day{day:02}"
        for part in range(1,3):
            compound=f"{yearStr}.{dayStr}.part_{part}"
            try:
                spam_spec = importlib.util.find_spec(compound)
                if spam_spec:
                    mod=importlib.import_module(compound)
                    if hasattr(mod, f'Part{part}'):
                        print(compound ,"exists",part)
            except:
                print("An exception occurred")
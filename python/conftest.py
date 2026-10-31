from pytest_benchmark.stats import Stats
from pytest import Config
import json

def pytest_benchmark_generate_json(config: Config, benchmarks, include_data, machine_info, commit_info):
    print(config)
    import os
    import re
    from datetime import datetime, timezone
    measurements = []
    for bench in benchmarks:
        param = getattr(bench, "param", "")
        if not param:
            continue
            
        m = re.match(r"^(\d+)_day(\d+)_(.+)_part(\d+)$", param)
        if not m:
            continue
            
        year, day, name, part = m.groups()
        
        series_key = f"{year}/{day}/part{part}"
        group_key = name
        
        mean_sec = bench.stats.mean
        duration_ns = str(int(mean_sec * 1_000_000_000))
        
        iterations = getattr(bench.stats, "rounds", None)
        if not iterations:
             iterations = getattr(bench, "iterations", 1)

        measurements.append({
            "series_key": series_key,
            "group_key": group_key,
            "duration": duration_ns,
            "iterations": iterations
        })

    output = {
        "name": "python",
        "hash": os.environ.get("GITHUB_SHA", "unknown"),
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "measurements": measurements
    }
    
    return output

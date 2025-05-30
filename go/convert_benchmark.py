import json
import re

import sys
KEY_ACTION="Action"
KEY_TEST="Test"
KEY_OUTPUT="Output"
#INPUT_REGEX_STR=r"\W*(?P<iterations>[0-9]+)[(\\t)\W]+(?P<timing>[0-9]+) ns\/op"
INPUT_REGEX_STR = r'(?P<iterations>\d+)\s+(?P<timing>\d+(\.\d+)?)\sns'

INPUT_REGEX=re.compile(INPUT_REGEX_STR)
def parse(input_stream):
    sys.stdout.write("[")
    output_idx=0
    for line in input_stream.readlines():
        parsed_line=json.loads(line)
        if KEY_ACTION not in parsed_line or parsed_line[KEY_ACTION]!="output":
            continue
        if "Test" not in parsed_line or not parsed_line["Test"].startswith("Benchmark"):
            continue
        reg_result=INPUT_REGEX.search(parsed_line[KEY_OUTPUT])
        if reg_result:
            group_dict=reg_result.groupdict()
            if output_idx>0:
                sys.stdout.write(",")
            output_idx+=1
            sys.stdout.write(json.dumps({
                    "name": f"go_{parsed_line["Package"]}.{parsed_line["Test"]}",
                    "unit": "ns",
                    "value": float(group_dict["timing"]),
                    "extra": f"Number of Iterations: {group_dict["iterations"]}\n"
                }))
            sys.stdout.flush()
    sys.stdout.write("]")



parse(sys.stdin)
        
import json
import re

KEY_ACTION="Action"
KEY_TEST="Test"
KEY_OUTPUT="Output"
#INPUT_REGEX_STR=r"\W*(?P<iterations>[0-9]+)[(\\t)\W]+(?P<timing>[0-9]+) ns\/op"
INPUT_REGEX_STR = r"\s*(?P<iterations>\d+)\s+(?P<timing>\d+)\s+ns/op"

INPUT_REGEX=re.compile(INPUT_REGEX_STR)
with open('benchmark.txt','r') as input_file:
    data=list()
    for line in input_file.readlines():
        parsed_line=json.loads(line)
        if KEY_ACTION not in parsed_line or parsed_line[KEY_ACTION]!="output":
            continue
        if "Test" not in parsed_line or not parsed_line["Test"].startswith("Benchmark"):
            continue
        reg_result=INPUT_REGEX.match(parsed_line[KEY_OUTPUT])
        
        #print(parsed_line[KEY_OUTPUT])
        if reg_result:
            group_dict=reg_result.groupdict()
            data.append(
    {
        "name": f"{parsed_line["Package"]}.{parsed_line["Test"]}",
        "unit": "ns",
        "value": int(group_dict["timing"]),
        "extra": f"Number of Iterations: {group_dict["iterations"]}\n"
    }
                
            )
    with open("benchmark.json","w") as output_file:
        json.dump(data,output_file)
    
        
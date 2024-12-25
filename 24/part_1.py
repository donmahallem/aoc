import functools
import itertools
import codecs

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.read()
    register, wires = data.split("\r\n\r\n")

    register = {
        line.strip().split(":")[0]: int(line.strip().split(":")[1])
        for line in register.split("\r\n")
    }
    wires = [line.strip() for line in wires.split("\r\n")]
    wires = [
        tuple(item[0:3] + [item[4]]) for item in [line.split(" ") for line in wires]
    ]

connected = True
while connected:
    connected = False
    for line in wires:
        if line[3] in register:
            continue
        if line[0] in register and line[2] in register:
            val_a = register[line[0]]
            val_b = register[line[2]]
            if line[1] == "AND":
                out = val_a and val_b
            elif line[1] == "OR":
                out = val_a or val_b
            elif line[1] == "XOR":
                out = val_a ^ val_b
            register[line[3]] = out
            connected = True

sorted_keys = sorted(register.keys())

output_value = 0
for key in sorted_keys:
    if key[0] == "z":
        if register[key] > 0:
            key_val = int(key[1:])
            output_value += 1 << key_val

print(output_value)

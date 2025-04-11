import typing
from .part_1 import parse_input
from .operation import Operation


def get_full_adder_output(data: list[Operation], adder: int):
    key = "z" + str(adder).zfill(2)
    for entry in data:
        if entry.output == key:
            return entry
    raise Exception("Could not find output : " + key)


def get_full_adder_input(data: list[Operation], adder: int):
    key_x = str(adder).zfill(2)
    key_x, key_y = "x" + key_x, "y" + key_x
    inps = list()
    for entry in data:
        if key_x in entry.inputs or key_y in entry.inputs:
            inps.append(entry)
    return inps


def validate_full_adder_inputs(inputs: list[Operation, Operation], adder: int):
    key_x = str(adder).zfill(2)
    key_x, key_y = "x" + key_x, "y" + key_x
    if [key_x, key_y] != inputs[0].inputs and [key_x, key_y] != inputs[1].inputs:
        raise False
    return True


def get_full_adder(data: list[Operation], adder: int):
    inputs = get_full_adder_input(data, adder)
    output = get_full_adder_output(data, adder)
    if len(inputs) != 2:
        raise Exception(f"Found {len(inputs)}. Should be 2")
    if not validate_full_adder_inputs(inputs, adder):
        raise Exception("N")


def isInitial(val):
    return val[0] in ["x", "y"] and val[1:3].isnumeric()


def isEnd(val):
    return val[0] == "z" and val[1:3].isnumeric()


def getInitialValue(val):
    return int(val[1:3])


def calcInverseConnections(wires):
    inverse_connections = dict()
    for wire in wires:
        inverse_connections[wire[3]] = tuple([wire[0], wire[2]])
    return inverse_connections


def validateChild(cons, cur, target_value):
    count = 0
    for test in cons[cur]:
        if isInitial(test):
            if (
                getInitialValue(test) == target_value
                or getInitialValue(test) == target_value - 1
            ):
                count += 1
        else:
            if validateChild(cons, test, target_value):
                count += 1
    return count == 2


def validateEndPosition(cons, end_pos):
    target_val = getInitialValue(end_pos)

    return validateChild(cons, end_pos, target_val)


def getInvalidChilds(cons, end_pos):
    target_val = getInitialValue(end_pos)
    invalid_childs = []
    for a in cons[end_pos]:
        if validateChild(cons, a, target_val):
            continue
        invalid_childs.append(a)
    return invalid_childs


def getNotInPair(tp, exclude):
    if tp[0] == exclude:
        return tp[1]
    elif tp[1] == exclude:
        return tp[0]
    return None


def getInput(wire_dict, z, operator):
    look_for_x = "x" + str(z).zfill(2)
    look_for_y = "y" + str(z).zfill(2)
    if (look_for_x, operator, look_for_y) in wire_dict:
        return wire_dict[(look_for_x, operator, look_for_y)]
    elif (look_for_y, operator, look_for_x) in wire_dict:
        return wire_dict[(look_for_y, operator, look_for_x)]


def findInWires(wires_dict, a, b, operator):
    if (a, operator, b) in wires_dict:
        return wires_dict[(a, operator, b)]
    elif (b, operator, a) in wires_dict:
        return wires_dict[(b, operator, a)]


type Wire = tuple[str, str, str, str]


def findBitInfos(wires: list[Wire], input_bit: int) -> tuple[Wire, Wire, Wire]:
    x_part = str(input_bit).zfill(2)
    x_part, y_part, z_part = "x" + x_part, "y" + x_part, "z" + x_part


def Part2(input: typing.TextIO) -> int:
    registers, wires = parse_input(input)
    all_ends = [
        wire[3] for wire in wires if wire[3][1:].isnumeric() and wire[3][0] == "z"
    ]
    all_ends = sorted(all_ends)
    print(wires)

    inverse_cons = calcInverseConnections(wires)
    carries = dict()
    # find initial carry
    check_val = ("x00", "AND", "y00")
    check_val_reversed = tuple(reversed(check_val))
    for wire in wires:
        # print(wire[0:3])
        if wire[0:3] == check_val or wire[0:3] == check_val_reversed:
            carries[0] = wire[3]
            break
    print(check_val, check_val_reversed, carries)
    swaps = list()
    swapped = True
    while swapped:
        swapped = False

        inverse_wires_dict = dict()
        wires_dict = dict()
        for wire in wires:
            inverse_wires_dict[wire[3]] = wire[0:3]
            wires_dict[wire[0:3]] = wire[3]
        for z in range(1, 44):
            z_str = "z" + str(z).zfill(2)
            print("Checking", z_str)
            and_input = getInput(wires_dict, z, "AND")
            xor_input = getInput(wires_dict, z, "XOR")
            print("XOR:", xor_input, "AND:", and_input)
            # checking xor out
            if carries[(z - 1)] not in inverse_cons[z_str]:
                print("Carry not in output!")
                print("Expected:", carries[(z - 1)], "Got:", inverse_cons[z_str])
                exit()
            carry_input = carries[(z - 1)]
            print("Carry:", carry_input)
            output_xor = findInWires(wires_dict, carry_input, xor_input, "XOR")
            if output_xor == None:
                print("No output xor found")
                exit()
            elif output_xor != z_str:
                print("Output xor not matching")
                print("Expected:", z_str, "Got:", output_xor)
                exit()
            output_and = findInWires(wires_dict, xor_input, carry_input, "AND")
            if output_and == None:
                print("Output or not matching")
                exit()
            output_or = findInWires(wires_dict, output_and, and_input, "OR")
            if output_or == None:
                print("Output or not matching")
                exit()
            carries[z] = output_or

    return sorted(swaps)

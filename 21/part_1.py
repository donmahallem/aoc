import codecs
import itertools
import functools


def shortestPathRobotNumeric(sequence):
    current_position = (3, 2)
    output_seqence = []
    location_dict = dict()
    location_dict["7"] = (0, 0)
    location_dict["8"] = (0, 1)
    location_dict["9"] = (0, 2)
    location_dict["4"] = (1, 0)
    location_dict["5"] = (1, 1)
    location_dict["6"] = (1, 2)
    location_dict["1"] = (2, 0)
    location_dict["2"] = (2, 1)
    location_dict["3"] = (2, 2)
    location_dict["0"] = (3, 1)
    location_dict["A"] = (3, 2)
    for i, looking_for in enumerate(sequence):
        cur_y, cur_x = current_position
        next_y, next_x = location_dict[looking_for]
        diff_y, diff_x = next_y - cur_y, next_x - cur_x
        if cur_y == 3 and next_x == 0:
            output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
            output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
        else:
            output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
            output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
        output_seqence += ["A"]
        current_position = (next_y, next_x)
    return output_seqence


def shortestPathRobotNumericSingle(current_position, character):
    output_seqence = []
    location_dict = dict()
    location_dict["7"] = (0, 0)
    location_dict["8"] = (0, 1)
    location_dict["9"] = (0, 2)
    location_dict["4"] = (1, 0)
    location_dict["5"] = (1, 1)
    location_dict["6"] = (1, 2)
    location_dict["1"] = (2, 0)
    location_dict["2"] = (2, 1)
    location_dict["3"] = (2, 2)
    location_dict["0"] = (3, 1)
    location_dict["A"] = (3, 2)
    cur_y, cur_x = current_position
    next_y, next_x = location_dict[character]
    diff_y, diff_x = next_y - cur_y, next_x - cur_x
    if cur_y == 3 and next_x == 0:
        output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
        output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
    else:
        output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
        output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
    output_seqence += ["A"]
    return (next_y, next_x), output_seqence


def shortestPathRobotDirectional(sequence, depth=0):
    current_position = (0, 2)
    output_seqence = []
    location_dict = dict()
    location_dict["^"] = (0, 1)
    location_dict["A"] = (0, 2)
    location_dict["<"] = (1, 0)
    location_dict["v"] = (1, 1)
    location_dict[">"] = (1, 2)
    for i, looking_for in enumerate(sequence):
        cur_y, cur_x = current_position
        next_y, next_x = location_dict[looking_for]
        diff_y, diff_x = next_y - cur_y, next_x - cur_x
        if depth == 0 and cur_x == 2 and cur_y == 0 and next_y == 1 and next_x == 0:
            output_seqence += ["<", "v", "<"]
        elif cur_y == 0 and next_x == 0:
            output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
            output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
        else:
            output_seqence += ["<" if diff_x < 0 else ">"] * abs(diff_x)
            output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
        output_seqence += ["A"]
        current_position = (next_y, next_x)
    return output_seqence


@functools.cache
def possiblePathsDirectionalCoord(start, end):
    start_y, start_x = start
    end_y, end_x = end
    diff_y, diff_x = end_y - start_y, end_x - start_x
    output_seqence = ["<" if diff_x < 0 else ">"] * abs(diff_x)
    output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
    output_seqence = set(itertools.permutations(output_seqence))
    if start_x == 0 and start_y == 1 and end_y == 0:
        output_seqence.remove(tuple(["^"] * abs(diff_y) + [">"] * abs(diff_x)))
    elif start_y == 0 and end_x == 0 and end_y == 1:
        output_seqence.remove(tuple(["<"] * abs(diff_x) + ["v"] * abs(diff_y)))
    return output_seqence


@functools.cache
def possiblePathsNumericalCoord(start, end):
    start_y, start_x = start
    end_y, end_x = end
    diff_y, diff_x = end_y - start_y, end_x - start_x
    output_seqence = ["<" if diff_x < 0 else ">"] * abs(diff_x)
    output_seqence += ["^" if diff_y < 0 else "v"] * abs(diff_y)
    output_seqence = set(itertools.permutations(output_seqence))
    if start_x == 0 and end_y == 3:
        output_seqence.remove(tuple(["v"] * abs(diff_y) + [">"] * abs(diff_x)))
    elif start_y == 3 and end_x == 0:
        output_seqence.remove(tuple(["<"] * abs(diff_x) + ["^"] * abs(diff_y)))
    return output_seqence


@functools.cache
def possiblePathsNumerical(start, end):
    location_dict = dict()
    location_dict["7"] = (0, 0)
    location_dict["8"] = (0, 1)
    location_dict["9"] = (0, 2)
    location_dict["4"] = (1, 0)
    location_dict["5"] = (1, 1)
    location_dict["6"] = (1, 2)
    location_dict["1"] = (2, 0)
    location_dict["2"] = (2, 1)
    location_dict["3"] = (2, 2)
    location_dict["0"] = (3, 1)
    location_dict["A"] = (3, 2)
    return [
        item + ("A",)
        for item in possiblePathsNumericalCoord(
            location_dict[start], location_dict[end]
        )
    ]


@functools.cache
def possiblePathsDirectional(start, end):
    location_dict = dict()
    location_dict["^"] = (0, 1)
    location_dict["A"] = (0, 2)
    location_dict["<"] = (1, 0)
    location_dict["v"] = (1, 1)
    location_dict[">"] = (1, 2)
    return [
        item + ("A",)
        for item in possiblePathsDirectionalCoord(
            location_dict[start], location_dict[end]
        )
    ]


@functools.cache
def findShortest(inp: tuple[str], max_depth=2, depth=0):
    if depth == max_depth:
        return len(inp)
    sequence = 0
    for i in range(-1, len(inp) - 1):
        if depth == 0:
            next_sequence = possiblePathsNumerical(
                inp[i] if i >= 0 else "A", inp[i + 1]
            )
        else:
            next_sequence = possiblePathsDirectional(
                inp[i] if i >= 0 else "A", inp[i + 1]
            )
        seqs = [
            findShortest(next_sequence_sub, max_depth=max_depth, depth=depth + 1)
            for next_sequence_sub in next_sequence
        ]
        seqs = sorted(seqs)
        sequence += seqs[0]
    return sequence


def calculate_puzzle_output(data, numeric_keypads):
    summe = 0
    for line in data:
        numeric_part = int(
            "".join(
                [
                    item
                    for item in line
                    if item in ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]
                ]
            )
        )
        l = findShortest(tuple(line), numeric_keypads)
        print(line, l, numeric_part)
        summe += numeric_part * l
    return summe


if __name__ == "__main__":
    test_data = False
    with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
        data = [line.strip() for line in f.readlines()]

    print(calculate_puzzle_output(data, 3))

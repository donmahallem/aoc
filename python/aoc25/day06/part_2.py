import sys
import typing


def Part2(input: typing.TextIO) -> int:
    columns: list[int] = []
    ops: list[str] = []
    for line in input:
        line = line.rstrip()
        if not columns:
            columns = [0] * len(line)
        if len(line) > len(columns):
            columns.extend([0] * (len(line) - len(columns)))
        for i, c in enumerate(line):
            if c >= "0" and c <= "9":
                columns[i] = columns[i] * 10 + int(c)
            elif c == "+" or c == "*":
                ops.append(c)
        if len(ops) > 0:
            break

    totalSum = 0
    currentNumber = 0
    currentOperator = 0
    for i in range(len(columns)):
        if columns[i] > 0:
            if currentNumber == 0:
                currentNumber = columns[i]
            else:
                match ops[currentOperator]:
                    case "+":
                        currentNumber += columns[i]
                    case "*":
                        currentNumber *= columns[i]
        else:
            totalSum += currentNumber
            currentNumber = 0
            currentOperator += 1
    totalSum += currentNumber
    return totalSum


if __name__ == "__main__":
    print(Part2(sys.stdin))

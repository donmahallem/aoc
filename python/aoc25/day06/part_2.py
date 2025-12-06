from functools import reduce
import operator
import sys
import typing

def parseInput(input: typing.TextIO) -> tuple[list[tuple[int, int]], list[str]]:
    
    inputLines = input.readlines()
    dataLines = inputLines[:-1]
    operatorLine = inputLines[-1]


    numbers = [0]*len(dataLines[0])
    for y in range(len(dataLines)):
        line = dataLines[y]
        for x in range(len(line)):
            if line[x].isdigit():
                numbers[x] = numbers[x]*10 + int(line[x])

    output=[]
    operators=[]
    currentNumbers=[]
    for x in range(len(numbers)):
        opLine = x<len(operatorLine) and operatorLine[x] in '+*'
        if numbers[x]==0 and not opLine:
            continue
        if opLine:
            operators.append(operatorLine[x])
            if x>0:
                output.append(currentNumbers)
                currentNumbers=[]
        currentNumbers.append(numbers[x])

    output.append(currentNumbers)
    operators = [c for c in operatorLine if c in '+*']
    return output, operators


def Part2(input: typing.TextIO) -> int:
    numbers, operators = parseInput(input)
    return sum([reduce(operator.add if b == '+' else operator.mul, a) for (a,b) in zip(numbers, operators)])


if __name__ == "__main__":
    print(Part2(sys.stdin))

import sys
import typing

from .shared import parse_ast
def Part2(input_data: typing.TextIO) -> int:
    scores: list[int] = []
    for line in input_data:
        line = line.strip()
        # Process each line as needed
        result, ast=parse_ast(line)
        if result == 0:
            # Calculate the completion score for the remaining sequence
            score = 0
            while ast:
                score = score * 5 + {')': 1, ']': 2, '}': 3, '>': 4}[ast.pop()]
            scores.append(score)

    scores.sort()
    return scores[len(scores) // 2] if scores else 0


if __name__ == "__main__":
    print(Part2(sys.stdin))

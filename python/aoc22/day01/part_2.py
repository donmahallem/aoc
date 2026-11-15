import sys
import typing


def Part2(input_data: typing.TextIO) -> int:
    current_value: int = 0
    calorie_list: list[int] = []
    for line in input_data.read().splitlines():
        if not line:
            if current_value > 0:
                calorie_list.append(current_value)
                current_value = 0
            continue
        current_value += int(line)
    if current_value > 0:
        calorie_list.append(current_value)
    calorie_list.sort(reverse=True)
    return sum(calorie_list[:3]) if calorie_list else 0


if __name__ == "__main__":
    print(Part2(sys.stdin))

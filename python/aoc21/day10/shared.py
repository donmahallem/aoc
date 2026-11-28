def parse_ast(line: str) -> tuple[int, list[str]]:
    depth_sequence: list[str] = []
    for char in line:
        match char:
            case "(":
                depth_sequence.append(")")
                continue
            case "[":
                depth_sequence.append("]")
                continue
            case "{":
                depth_sequence.append("}")
                continue
            case "<":
                depth_sequence.append(">")
                continue
            case ">":
                if depth_sequence and depth_sequence[-1] == ">":
                    depth_sequence.pop()
                else:
                    return 25137, depth_sequence
                continue
            case "]":
                if depth_sequence and depth_sequence[-1] == "]":
                    depth_sequence.pop()
                else:
                    return 57, depth_sequence
                continue
            case "}":
                if depth_sequence and depth_sequence[-1] == "}":
                    depth_sequence.pop()
                else:
                    return 1197, depth_sequence
                continue
            case ")":
                if depth_sequence and depth_sequence[-1] == ")":
                    depth_sequence.pop()
                else:
                    return 3, depth_sequence
                continue
            case _:
                raise ValueError(f"Unexpected character: {char}")
    return 0, depth_sequence

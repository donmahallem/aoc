import typing
import heapq
from .shared import _build_graph, Graph


def _get_turn_cost(d1: tuple[int, int], d2: tuple[int, int]) -> int:
    if d1 == d2:
        return 0
    if d1[0] * d2[0] + d1[1] * d2[1] == 0:
        return 1000
    return 2000


def Part1(input: typing.TextIO) -> int:
    graph = _build_graph(input)
    if not graph.start_node or not graph.end_node:
        return 0

    start_pos = graph.start_node.pos
    pq = [(0, start_pos, (0, 1))]
    visited: dict[tuple[tuple[int, int], tuple[int, int]], int] = {}

    while pq:
        cost, pos, curr_dir = heapq.heappop(pq)
        if pos == graph.end_node.pos:
            return cost

        state = (pos, curr_dir)
        if state in visited and visited[state] <= cost:
            continue
        visited[state] = cost

        current_node = graph.nodes[pos]
        for neighbor, weight, edge_start_dir, edge_final_dir, _ in current_node.edges:
            turn_penalty = _get_turn_cost(curr_dir, edge_start_dir)
            new_cost = cost + turn_penalty + weight
            new_state = (neighbor.pos, edge_final_dir)
            if new_state not in visited or visited[new_state] > new_cost:
                heapq.heappush(pq, (new_cost, neighbor.pos, edge_final_dir))

    return -1  # No path found

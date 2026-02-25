import typing
import heapq
from .shared import _build_graph, Graph
import sys


def get_turn_cost(d1: tuple[int, int], d2: tuple[int, int]) -> int:
    if d1 == d2: return 0
    if d1[0] * d2[0] + d1[1] * d2[1] == 0: return 1000
    return 2000


def Part2(input: typing.TextIO) -> int:
    graph = _build_graph(input)
    if not graph.start_node or not graph.end_node:
        return 0

    start_state = (graph.start_node.pos, (0, 1))
    dist = {start_state: 0}
    predecessors = {start_state: []}

    pq = [(0, graph.start_node.pos, (0, 1))]
    min_end_cost = sys.maxsize
    end_states = []

    while pq:
        cost, pos, curr_dir = heapq.heappop(pq)

        if cost > dist.get((pos, curr_dir), sys.maxsize):
            continue

        if pos == graph.end_node.pos:
            if cost < min_end_cost:
                min_end_cost = cost
                end_states = [(pos, curr_dir)]
            elif cost == min_end_cost:
                end_states.append((pos, curr_dir))
            continue

        current_node = graph.nodes[pos]
        for neighbor, weight, edge_start_dir, edge_final_dir, path_cells in current_node.edges:
            new_cost = cost + get_turn_cost(curr_dir, edge_start_dir) + weight
            new_state = (neighbor.pos, edge_final_dir)

            old_dist = dist.get(new_state, sys.maxsize)

            if new_cost < old_dist:
                dist[new_state] = new_cost
                # We found a better path, reset predecessors
                predecessors[new_state] = [((pos, curr_dir), path_cells)]
                heapq.heappush(pq, (new_cost, neighbor.pos, edge_final_dir))
            elif new_cost == old_dist:
                # We found another equally good path
                predecessors[new_state].append(((pos, curr_dir), path_cells))

    # Backtrack to find all unique tiles
    best_tiles = set()
    stack = end_states
    seen_states = set(end_states)

    while stack:
        curr_state = stack.pop()
        for prev_state, edge_cells in predecessors.get(curr_state, []):
            best_tiles.update(edge_cells)
            if prev_state not in seen_states:
                seen_states.add(prev_state)
                stack.append(prev_state)

    return len(best_tiles)

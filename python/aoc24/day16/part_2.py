import typing
import heapq
import sys
from .shared import _build_graph, Graph, GraphNode

from typing import TypeAlias

Coord: TypeAlias = tuple[int, int]
Direction: TypeAlias = tuple[int, int]
State: TypeAlias = tuple[Coord, Direction]
ParentInfo: TypeAlias = tuple[State, set[Coord]]


def get_turn_cost(d1: Direction, d2: Direction) -> int:
    if d1 == d2:
        return 0
    if d1[0] * d2[0] + d1[1] * d2[1] == 0:
        return 1000
    return 2000


def Part2(input_data: typing.TextIO) -> int:
    graph: Graph = _build_graph(input_data)
    if not graph.start_node or not graph.end_node:
        return 0

    start_pos: Coord = graph.start_node.pos
    end_pos: Coord = graph.end_node.pos

    start_state: State = (start_pos, (0, 1))

    dist: dict[State, int] = {start_state: 0}
    predecessors: dict[State, list[ParentInfo]] = {start_state: []}

    pq: list[tuple[int, Coord, Direction]] = [(0, start_pos, (0, 1))]

    min_end_cost: int = sys.maxsize
    end_states: list[State] = []

    while pq:
        cost, pos, curr_dir = heapq.heappop(pq)

        if cost > dist.get((pos, curr_dir), sys.maxsize):
            continue

        if pos == end_pos:
            if cost < min_end_cost:
                min_end_cost = cost
                end_states = [(pos, curr_dir)]
            elif cost == min_end_cost:
                end_states.append((pos, curr_dir))
            continue

        if cost > min_end_cost:
            continue

        current_node: GraphNode = graph.nodes[pos]

        for edge in current_node.edges:
            neighbor: GraphNode = edge[0]
            weight: int = edge[1]
            edge_start_dir: Direction = edge[2]
            edge_final_dir: Direction = edge[3]
            path_cells: set[Coord] = edge[4]

            new_cost = cost + get_turn_cost(curr_dir, edge_start_dir) + weight
            new_state: State = (neighbor.pos, edge_final_dir)

            old_dist = dist.get(new_state, sys.maxsize)

            if new_cost < old_dist:
                dist[new_state] = new_cost
                predecessors[new_state] = [((pos, curr_dir), path_cells)]
                heapq.heappush(pq, (new_cost, neighbor.pos, edge_final_dir))
            elif new_cost == old_dist:
                if new_state not in predecessors:
                    predecessors[new_state] = []
                predecessors[new_state].append(((pos, curr_dir), path_cells))

    best_tiles: set[Coord] = {start_pos, end_pos}
    stack: list[State] = end_states
    seen_states: set[State] = set(end_states)

    while stack:
        curr_state = stack.pop()
        for prev_state, edge_cells in predecessors.get(curr_state, []):
            best_tiles.update(edge_cells)
            if prev_state not in seen_states:
                seen_states.add(prev_state)
                stack.append(prev_state)
    return len(best_tiles)

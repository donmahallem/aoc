import typing


class GraphNode:
    __slots__ = ("pos", "edges")

    def __init__(self, pos: tuple[int, int]):
        self.pos = pos
        self.edges: list[tuple['GraphNode', int, tuple[int, int],
                               tuple[int, int], set[tuple[int, int]]]] = []


class Graph:
    __slots__ = ("nodes", "start_node", "end_node")

    def __init__(self):
        self.nodes: dict[tuple[int, int], GraphNode] = {}
        self.start_node: GraphNode | None = None
        self.end_node: GraphNode | None = None


def _build_graph(input_data: typing.TextIO) -> Graph:
    grid = [line.strip() for line in input_data if line.strip()]
    rows, cols = len(grid), len(grid[0])

    def get_neighbors(r, c):
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] != '#':
                yield dr, dc, nr, nc

    pois = {}
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '#': continue

            exits = list(get_neighbors(r, c))
            is_poi = False

            if grid[r][c] in 'SE':
                is_poi = True
            elif len(exits) > 2:
                is_poi = True
            elif len(exits) == 2:
                (dr1, dc1, _, _), (dr2, dc2, _, _) = exits
                if dr1 + dr2 != 0 or dc1 + dc2 != 0:
                    is_poi = True
            elif len(exits) == 1:
                is_poi = True

            if is_poi:
                node = GraphNode((r, c))
                pois[(r, c)] = node

    graph = Graph()
    graph.nodes = pois

    for (r, c), node in pois.items():
        if grid[r][c] == 'S': graph.start_node = node
        if grid[r][c] == 'E': graph.end_node = node

        for dr, dc, nr, nc in get_neighbors(r, c):
            path_cells = {(r, c), (nr, nc)}
            dist = 1
            curr_r, curr_c = nr, nc
            last_dr, last_dc = dr, dc

            while (curr_r, curr_c) not in pois:
                for ndr, ndc, nnr, nnc in get_neighbors(curr_r, curr_c):
                    if (ndr, ndc) != (-last_dr, -last_dc):
                        last_dr, last_dc = ndr, ndc
                        curr_r, curr_c = nnr, nnc
                        path_cells.add((curr_r, curr_c))
                        dist += 1
                        break

            node.edges.append((pois[(curr_r, curr_c)], dist, (dr, dc),
                               (last_dr, last_dc), path_cells))

    return graph

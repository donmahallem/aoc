import io
import unittest
from .shared import _Field


class AocYear24Day12_Field(unittest.TestCase):

    def test_constructor(self):
        f = _Field(1, 2, [[1, 2]])
        self.assertEqual(f.width, 1)
        self.assertEqual(f.height, 2)
        self.assertEqual(f.data, [[1, 2]])

    def test_parseInput_simple(self):
        inp = '''AAA\nBBB\nCCB'''
        f_obj = io.StringIO(inp)
        f = _Field.parse_input(f_obj)
        self.assertEqual(f.width, 3)
        self.assertEqual(f.height, 3)
        self.assertEqual(
            f.data, [bytearray(b'AAA'),
                     bytearray(b'BBB'),
                     bytearray(b'CCB')])

    def test_collectIslands_simple(self):
        inp = '''AAA\nBBB\nCCB'''
        f_obj = io.StringIO(inp)
        f = _Field.parse_input(f_obj)
        islands = f.collect_islands()
        self.assertEqual(len(islands), 3,
                         f"Expected 3 islands, got {len(islands)}\n{islands}")
        self.assertEqual(
            islands[0], {(0, 0), (0, 1), (0, 2)},
            f"Expected first island to be {(0,0),(0,1),(0,2)}, got {islands[0]}"
        )
        self.assertEqual(islands[1], {(1, 0), (1, 1), (1, 2), (2, 2)})

    def test_collectIslands_edges(self):
        inp = '''ABA\nBAB'''
        f_obj = io.StringIO(inp)
        f = _Field.parse_input(f_obj)
        islands = f.collect_islands()
        self.assertEqual(len(islands), 6,
                         f"Expected 6 islands, got {len(islands)}\n{islands}")

    def test_count_edges_L_shape(self):
        # L shape
        test_island = set([(0, 0), (1, 0), (1, 1)])
        edges = _Field.count_edges(test_island)
        self.assertEqual(edges, (8, 6))

    def test_count_edges_plus(self):
        # Plus shape
        test_island = set([(0, 1), (1, 0), (1, 1), (1, 2), (2, 1)])
        edges = _Field.count_edges(test_island)
        self.assertEqual(edges, (12, 12))


if __name__ == '__main__':
    unittest.main()

import unittest
import io
from aoc24.day06.part_2 import Part2

testData = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""


class Test2024Day06Part02(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 6, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

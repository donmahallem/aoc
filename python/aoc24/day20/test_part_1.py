import unittest
import io
from aoc24.day20.part_1 import handle

testData = """###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############"""


class Test2024Day20Part01(unittest.TestCase):
    def test_handle_cheat20(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = handle(f, 20)
            self.assertEqual(result, 5, "The sum is wrong.")

    def test_handle_cheat1(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = handle(f, 1)
            self.assertEqual(result, 44, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

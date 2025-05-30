import unittest
import io
from aoc24.day08.part_1 import Part1

testData = """............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"""


class Test2024Day08Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 14, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

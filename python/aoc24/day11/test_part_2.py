import unittest
import io
from aoc24.day11.part_2 import Part2

testData = "125 17"


class Test2024Day11Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 65601038650482, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

import unittest
import io
from aoc24.day09.part_1 import Part1

testData = """2333133121414131402"""


class Test2024Day09Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 1928, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

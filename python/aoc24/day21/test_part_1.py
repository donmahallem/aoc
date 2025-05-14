import unittest
import io
from aoc24.day21.part_1 import Part1

testData = """029A
980A
179A
456A
379A"""


class Test2024Day15Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 126384, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

import unittest
import io
from aoc24.day17.part_2 import Part2

testData = """Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0"""


class Test2024Day15Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 117440, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

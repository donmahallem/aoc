import unittest
import io
from aoc25.day03.part_1 import Part1

testData = """987654321111111
811111111111119
234234234234278
818181911112111"""


class Test2025Day03Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 17452, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

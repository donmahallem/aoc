import unittest
import io
from aoc25.day06.part_1 import Part1

testData = """123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +"""


class Test2025Day06Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 4277556, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

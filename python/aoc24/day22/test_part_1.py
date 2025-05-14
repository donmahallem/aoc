import unittest
import io
from aoc24.day22.part_1 import Part1

testData = """1
10
100
2024"""


class Test_Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 37327623, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

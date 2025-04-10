import unittest
import io
from aoc24.day22.part_2 import Part2

testData = """1
2
3
2024"""


class Test_Part02(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 23, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

import unittest
import io
from aoc25.day05.part_1 import Part1

testData = """3-5
10-14
16-20
12-18

1
5
8
11
17
32"""


class Test2025Day05Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 3, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

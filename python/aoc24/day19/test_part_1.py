import unittest
import io
from aoc24.day19.part_1 import Part1

testData = """r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb"""

class Test2024Day19Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 6, "The sum is wrong.")

if __name__ == "__main__":
    unittest.main()

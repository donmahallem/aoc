import unittest
import io
from aoc24.day03.part_2 import Part2

testData = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

class Test2024Day03Part02(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 48, "The sum is wrong.")

if __name__ == "__main__":
    unittest.main()

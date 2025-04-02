import unittest
import io
from aoc24.day03.part_1 import Part1

testData = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"


class Test2024Day03Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 161, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

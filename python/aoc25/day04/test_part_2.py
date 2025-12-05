import unittest
import io
from aoc25.day04.part_2 import Part2
from aoc25.day03.test_part_2 import testData


class Test2025Day04Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 43, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

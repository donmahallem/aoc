import unittest
import io
from aoc25.day06.part_2 import Part2
from aoc25.day06.test_part_1 import testData


class Test2025Day05Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 3263827, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

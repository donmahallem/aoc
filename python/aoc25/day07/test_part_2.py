import unittest
import io
from aoc25.day07.part_2 import Part2
from aoc25.day07.test_part_1 import testData


class Test2025Day07Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 40, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

import unittest
import io
from aoc24.day15.part_2 import Part2, parseField, next_empty
from .test_part_1 import testData


class Test2024Day15Part01(unittest.TestCase):

    def test_parseField(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            field, player, movements = parseField(f)
            self.assertEqual(player, (4, 8), "The sum is wrong.")
            self.assertEqual(field.shape, (10, 20), "The sum is wrong.")
            self.assertEqual(len(movements), 700, "The sum is wrong.")

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 9021, "The sum is wrong.")

    def test_nextEmpty(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            field, player, movements = parseField(f)
            next = next_empty(field, player[0], player[1], 0, -1)
            self.assertEqual(next, (4, 5), "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

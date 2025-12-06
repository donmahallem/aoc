import unittest
import io
from aoc25.day02.part_2 import Part2

testData = """11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124"""


class Test2025Day02Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 4174379265, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

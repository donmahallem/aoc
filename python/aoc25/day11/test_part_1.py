import unittest
import io
from aoc25.day11.part_1 import Part1

testData = """aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out"""


class Test2025Day11Part01(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f)
            self.assertEqual(result, 5, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

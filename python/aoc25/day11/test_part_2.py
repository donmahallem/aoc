import unittest
import io
from aoc25.day11.part_2 import Part2
from aoc25.day11.test_part_1 import testData

testData = """svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out"""


class Test2025Day11Part02(unittest.TestCase):

    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part2(f)
            self.assertEqual(result, 2, "The sum is wrong.")


if __name__ == "__main__":
    unittest.main()

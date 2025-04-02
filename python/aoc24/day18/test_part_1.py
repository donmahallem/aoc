import unittest
import io
from aoc24.day18.part_1 import Part1

testData = """5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0"""

class Test2024Day18Part01(unittest.TestCase):
    def test_result(self):
        with io.StringIO() as f:
            f.write(testData)
            f.seek(0)
            result = Part1(f, 7, 12)
            self.assertEqual(result, 22, "The sum is wrong.")

if __name__ == "__main__":
    unittest.main()

import unittest
import io
import aoc24.day24.part_2 as part2
from .operation import Operation
testData="""x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00"""

class Test2024Day15Part02(unittest.TestCase):
    def test_get_full_adder_output(self):
        op_1=Operation(["a","b"],"AND","x01")
        op_2=Operation(["a","b"],"AND","z02")
        self.assertEqual(part2.get_full_adder_output([op_1,op_2],2),op_2,"Expected to find op_2")
    def test_get_full_adder_input(self):
        op_1=Operation(["x01","b"],"AND","x01")
        op_2=Operation(["a","y01"],"AND","z02")
        self.assertEqual(part2.get_full_adder_input([op_1,op_2],1),[op_1,op_2],"Expected to find op_2")

if __name__ == "__main__":
    unittest.main()

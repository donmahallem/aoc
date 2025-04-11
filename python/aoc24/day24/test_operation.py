import unittest
from .operation import Operation

class OperatorTests(unittest.TestCase):

    def test_and_operator(self):
        obj=Operation([True,False],"AND","out")
        self.assertFalse(obj(), "The sum is wrong.")
        obj=Operation([True,True],"AND","out")
        self.assertTrue(obj(), "The sum is wrong.")

    def test_or_operator(self):
        obj=Operation([False,False],"OR","out")
        self.assertFalse(obj(), "False and False should be False")
        obj=Operation([True,False],"OR","out")
        self.assertTrue(obj(), "True | False = True")
        obj=Operation([True,True],"OR","out")
        self.assertTrue(obj(), "The sum is wrong.")

    def test_xor_operator(self):
        obj=Operation([False,False],"XOR","out")
        self.assertFalse(obj(), "False ^ False should be False")
        obj=Operation([True,False],"XOR","out")
        self.assertTrue(obj(), "True ^ False = True")
        obj=Operation([True,True],"XOR","out")
        self.assertFalse(obj(), "True ^ True should be False")
    def test_call_operator(self):
        obj=Operation([False,False],"OR","out")
        self.assertFalse(True&obj(), "True & False should be True")
        obj=Operation([False,False],"OR","out")
        self.assertFalse(False|obj(), "True | False should be True")
        obj=Operation([False,False],"OR","out")
        self.assertFalse(False^obj(), "True ^ False should be True")


if __name__ == "__main__":
    unittest.main()

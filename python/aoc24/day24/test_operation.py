import unittest
from .operation import Operation

class OperatorTests(unittest.TestCase):

    def test_and_operator(self):
        obj=Operation([True,"AND",False,"out"])
        self.assertFalse(obj(), "The sum is wrong.")
        obj=Operation([True,"AND",True,"out"])
        self.assertTrue(obj(), "The sum is wrong.")

    def test_or_operator(self):
        obj=Operation([False,"OR",False,"out"])
        self.assertFalse(obj(), "False and False should be False")
        obj=Operation([True,"OR",False,"out"])
        self.assertTrue(obj(), "True | False = True")
        obj=Operation([True,"OR",True,"out"])
        self.assertTrue(obj(), "The sum is wrong.")

    def test_xor_operator(self):
        obj=Operation([False,"XOR",False,"out"])
        self.assertFalse(obj(), "False ^ False should be False")
        obj=Operation([True,"XOR",False,"out"])
        self.assertTrue(obj(), "True ^ False = True")
        obj=Operation([True,"XOR",True,"out"])
        self.assertFalse(obj(), "True ^ True should be False")
    def test_call_operator(self):
        obj=Operation([False,"OR",False,"out"])
        self.assertFalse(True&obj(), "True & False should be True")
        obj=Operation([False,"OR",False,"out"])
        self.assertFalse(False|obj(), "True | False should be True")
        obj=Operation([False,"OR",False,"out"])
        self.assertFalse(False^obj(), "True ^ False should be True")


if __name__ == "__main__":
    unittest.main()

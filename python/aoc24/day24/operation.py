type OperationInput = str | bool


class Operation:
    def __init__(
        self, inputs: list[OperationInput, OperationInput], operator: str, output: str
    ):
        self.__inputs = inputs
        self.__output = output
        self.__operator = operator

    @property
    def operator(self):
        return self.__operator

    @property
    def output(self):
        return self.__output

    @property
    def inputs(self):
        return self.__inputs

    def __repr__(self):
        return f"Operation({str(self)})"

    def __str__(self):
        return f"{self.__inputs[0]} {self.__operator} {self.__inputs[1]} -> {self.__output}"

    def __or__(self, other):
        return self() | other

    def __ror__(self, other):
        return self() | other

    def __and__(self, other):
        return self() & other

    def __rand__(self, other):
        return self() & other

    def __xor__(self, other):
        return self() ^ other

    def __rxor__(self, other):
        return self() ^ other

    def __call__(self, *args, **kwds):
        a = (
            self.__inputs[0]()
            if type(self.__inputs[0]) == Operation
            else self.__inputs[0]
        )
        b = (
            self.__inputs[1]()
            if type(self.__inputs[1]) == Operation
            else self.__inputs[1]
        )
        if self.__operator == "AND":
            return a & b
        elif self.__operator == "OR":
            return a | b
        elif self.__operator == "XOR":
            return a ^ b
        else:
            raise Exception("Unknown operand: " + self.__operator)

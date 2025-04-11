type OperationInput=str|bool

class Operation:

    def __init__(self,inputs:list[OperationInput,OperationInput],operator:str,output:str):
        self.__inputs=inputs
        self.__output=output
        self.__operator=operator
    
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

    def __or__(self,other):
        return self() |other
    def __ror__(self,other):
        return self() |other

    def __and__(self,other):
        return self() & other
    def __rand__(self,other):
        return self() & other

    def __xor__(self,other):
        return self()^other
    def __rxor__(self,other):
        return self()^other

    def __call__(self, *args, **kwds):
        a=self.Inputs[0]() if type(self.Inputs[0])==Operation else self.Inputs[0]
        b=self.Inputs[1]() if type(self.Inputs[1])==Operation else self.Inputs[1]
        if self.Operand=="AND":
            return a&b
        elif self.Operand=="OR":
            return a|b
        elif self.Operand=="XOR":
            return a^b
        else:
            raise Exception("Unknown operand: "+self.Operand)
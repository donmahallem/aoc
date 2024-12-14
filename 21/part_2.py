import codecs
from part_1 import calculate_puzzle_output

if __name__=="__main__":
    test_data = False
    with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
        data = [line.strip() for line in f.readlines()]
    
    print(calculate_puzzle_output(data,26))
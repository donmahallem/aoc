import sys
import typing

type Range = typing.Tuple[int, int]

def parseInput(input: typing.TextIO) -> typing.Tuple[Range,typing.List[int]]:
    ranges=list()
    ingredients=list()
    firstBlock=True
    for line in input:
        line=line.strip()
        if line=="":
            firstBlock=False
            continue
        if firstBlock:
            parts=line.split("-")
            ranges.append( (int(parts[0]),int(parts[1])) )
        else:
            ingredients.append(int(line))
    
    return (ranges,ingredients)

def compressRanges(ranges:typing.List[Range]) -> typing.List[Range]:
    ranges=sorted(ranges,key=lambda r:r[0])
    compressed=list()
    curStart,curEnd=ranges[0]
    for r in ranges[1:]:
        if r[0]<=curEnd+1:
            # overlap
            if r[1]>curEnd:
                curEnd=r[1]
        else:
            compressed.append( (curStart,curEnd) )
            curStart,curEnd=r
    compressed.append( (curStart,curEnd) )
    return compressed

def Part1(input: typing.TextIO) -> int:
    ranges,ingredients = parseInput(input)
    compresed= compressRanges(ranges)
    goodIngredients=0
    for ing in ingredients:
        for r in compresed:
            if r[0]<=ing<=r[1]:
                goodIngredients+=1
                break
    return goodIngredients


if __name__ == "__main__":
    print(Part1(sys.stdin))

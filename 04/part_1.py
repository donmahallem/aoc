import codecs
import regex
import numpy as np
with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]
search_terms="XMAS"
zeilen=len(data)
spalten=len(data[0])
count=0

def checkDir(x,y,dirx,diry,term):
    if(x+(dirx*len(term))+1<0 or x+(dirx*len(term))>spalten):
        return False
    elif(y+(diry*len(term))+1<0 or y+(diry*len(term))>zeilen):
        return False
    for i in range(len(term)):
        if data[y+diry*i][x+dirx*i]!=term[i]:
            return False
    return True

count=0
for y in range(0,zeilen):
    for x in range(spalten):
        if checkDir(x,y,1,0,search_terms):
            count+=1
        if checkDir(x,y,1,-1,search_terms):
            count+=1
        if checkDir(x,y,0,-1,search_terms):
            count+=1
        if checkDir(x,y,-1,-1,search_terms):
            count+=1
        if checkDir(x,y,-1,0,search_terms):
            count+=1
        if checkDir(x,y,-1,1,search_terms):
            count+=1
        if checkDir(x,y,0,1,search_terms):
            count+=1
        if checkDir(x,y,1,1,search_terms):
            count+=1
print(count)
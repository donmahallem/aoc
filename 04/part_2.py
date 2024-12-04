import codecs
import regex
import numpy as np
with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]
search_terms="XMAS"
zeilen=len(data)
spalten=len(data[0])
def checkDir(x,y,dirx,diry,term):
    if(x+(dirx*len(term))+1<0 or x+(dirx*len(term))>spalten):
        return False
    elif(y+(diry*len(term))+1<0 or y+(diry*len(term))>zeilen):
        return False
    for i in range(len(term)):
        if x==3 and y==9 and dirx==-1:
            print(data[y+diry*i][x+dirx*i],term[i])
        if data[y+diry*i][x+dirx*i]!=term[i]:
            return False
    return True

count=0
for y in range(1,zeilen-1):
    for x in range(1,spalten-1):
        da=0
        if data[y][x]=="A":
            if data[y-1][x-1]=="M" and data[y+1][x+1]=="S":
                da+=1
            if data[y-1][x-1]=="S" and data[y+1][x+1]=="M":
                da+=1
            if data[y+1][x-1]=="M" and data[y-1][x+1]=="S":
                da+=1
            if data[y+1][x-1]=="S" and data[y-1][x+1]=="M":
                da+=1
        if da==2:
            count+=1
print(count)
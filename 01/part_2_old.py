import codecs

l1=[]
l2=[]
with codecs.open("data_1.txt",encoding="utf8") as f:
    data=f.readlines()
    for line in data:
        a=line.strip().split()
        if(len(a)!=2):
            continue
        l1.append(int(a[0]))
        l2.append(int(a[1]))

output=0
for item1 in l1:
    similar = [ 1 for item2 in l2 if item2==item1]
    output +=item1*len(similar)
print(output)
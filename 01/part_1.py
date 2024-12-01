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
l1=sorted(l1)
l2=sorted(l2)

l3=[]
for i in range(len(l1)):
    l3.append(abs(l1[i]-l2[i]))
print(sum(l3))
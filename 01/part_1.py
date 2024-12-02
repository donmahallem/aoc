import codecs

l1 = []
l2 = []
with codecs.open("data_1.txt", encoding="utf8") as f:
    data = f.readlines()
    for line in data:
        a = line.strip().split()
        if len(a) != 2:
            continue
        l1.append(int(a[0]))
        l2.append(int(a[1]))
l1 = sorted(l1)
l2 = sorted(l2)

ergebnis = sum(abs(a - b) for a, b in zip(l1, l2))
print(ergebnis)

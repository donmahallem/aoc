import codecs
import regex

with codecs.open("data_1.txt", encoding="utf8") as f:
    data = f.readlines()
data = "\n".join(data)
enabled = True
comp = regex.compile("do(n't)?\(\)", flags=regex.MULTILINE)
comp2 = regex.compile("mul\((\d+)\,(\d+)\)", flags=regex.MULTILINE)
total = 0
for item in comp.split(data):
    if item == "n't":
        enabled = False
        continue
    elif item == None:
        enabled = True
    elif enabled:
        findings = comp2.findall(item)
        total += sum([int(a) * int(b) for a, b in findings])
print(total)

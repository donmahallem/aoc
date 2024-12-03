import codecs
import regex

with codecs.open("data_1.txt", encoding="utf8") as f:
    data = f.readlines()
data="".join(data)
comp=regex.compile("mul\((\d+)\,(\d+)\)",flags=regex.MULTILINE)
findings=comp.findall(data)
print(sum([int(a)*int(b) for a,b in findings]))
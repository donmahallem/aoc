import codecs
import re

with codecs.open("data2.txt", encoding="utf8") as f:
    data=f.read()
    data=data.split("\r\n\r\n")

button_regex=re.compile(r"(?:Button\s([A-Za-z]+)\:)\s(?:[XY]([+-]\d+)),\s(?:[XY]([+-]\d+))")
prize_regex=re.compile(r"(?:Prize)\:\s(?:[XY]=(\d+)),\s*(?:[XY]=(\d+))")
machines=[]
for machine in data:
    new_machine={}
    new_machine["buttons"]=[]
    for line in machine.split("\r\n"):
        reg_res=button_regex.match(line)
        if reg_res:
            button=(reg_res.groups()[0],int(reg_res.groups()[1]),int(reg_res.groups()[2]))
            new_machine["buttons"].append(button)
        reg_res=prize_regex.match(line)
        if reg_res:
            new_machine["price"]=(int(reg_res.groups()[0]),int(reg_res.groups()[1]))
    machines.append(new_machine)

def calc(v1,v2,target):
    x_1,y_1=v1
    x_2,y_2=v2
    x_3,y_3=target
    v2_factor=((x_3*y_1)-(x_1*y_3))/((x_2*y_1)-(x_1*y_2))
    v1_factor=(x_3-(v2_factor*x_2))/x_1
    return (v1_factor,v2_factor)
summe=0
for i,machine in enumerate(machines):
    print(i,"/",len(machines))
    target_x,target_y=machine["price"]
    _,btn_a_x,btn_a_y=machine["buttons"][0]
    _,btn_b_x,btn_b_y=machine["buttons"][1]
    fac_1,fac_2=calc((btn_a_x,btn_a_y),(btn_b_x,btn_b_y),(target_x,target_y))
    if fac_1.is_integer() and fac_2.is_integer():
        summe+=fac_1*3+fac_2
print(int(summe))
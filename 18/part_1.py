import codecs
import numpy as np


test_data = False

CELL_CORRUPTED=1
WIDTH=7 if test_data else 71
HEIGHT=7 if test_data else 71
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = [line.strip() for line in f.readlines()]
data=[tuple(int(a) for a in line.split(","))[::-1] for line in data]
print(data)
game_map=np.zeros((HEIGHT,WIDTH),dtype=np.int8)
steps=12 if test_data else 1024
for i in range(0,steps):
    y,x=data[i]
    game_map[y,x]=CELL_CORRUPTED

def shortestPath():
    check_next=[(0,0)]
    dirs=(0,1),(1,0),(0,-1),(-1,0)
    path_cost=np.zeros(game_map.shape,dtype=np.int32)-1
    while len(check_next)>0:
        cur_pos=check_next.pop(-1)
        cur_y,cur_x=cur_pos
        current_path_cost=max(0,path_cost[cur_y,cur_x])
        for dir_y,dir_x in dirs:
            next_y,next_x=cur_y+dir_y,cur_x+dir_x
            if next_y<0 or next_x<0 or next_x>=WIDTH or next_y>=HEIGHT or game_map[next_y,next_x]==CELL_CORRUPTED:
                # OUTSIDE MAP OR CORRUPTED
                continue
            if path_cost[next_y,next_x]<0 or current_path_cost+1<path_cost[next_y,next_x]:
                path_cost[next_y,next_x]=current_path_cost+1
                check_next.append((next_y,next_x))
    return path_cost
path_cost=shortestPath()
print(path_cost[HEIGHT-1,WIDTH-1])
from typing import NamedTuple,Generator

class Point(NamedTuple):
    x: int
    y: int

    def __lt__(self, value):
        if self.y == value.y:
            return self.x < value.x
        return self.y < value.y
    
    def __le__(self, value):
        return self < value or self == value
    
    def __gt__(self, value):
        if self.y == value.y:
            return self.x > value.x
        return self.y > value.y
    
    def __ge__(self, value):
        return self > value or self == value
    
    def __add__(self, value):
        return Point(self.x + value.x, self.y + value.y)
    
    def __sub__(self, value):
        return Point(self.x - value.x, self.y - value.y)
    
    def __mul__(self, value):
        return Point(self.x * value, self.y * value)
    
    def __truediv__(self, value):
        return Point(self.x // value, self.y // value)
    
    def manhattan_distance(self, other: "Point") -> int:
        return abs(self.x - other.x) + abs(self.y - other.y)
    
    def __str__(self):
        return f"({self.x}, {self.y})"
    
    def __repr__(self):
        return f"Point(x={self.x}, y={self.y})"
    
    def __hash__(self):
        return hash((self.x, self.y))
    
    def __eq__(self, other):
        if not isinstance(other, Point):
            return NotImplemented
        return self.x == other.x and self.y == other.y
    
    def __ne__(self, other):
        if not isinstance(other, Point):
            return NotImplemented
        return self.x != other.x or self.y != other.y
    
    def __iter__(self):
        yield self.x
        yield self.y

    def __getitem__(self, index):
        if index == 0:
            return self.x
        elif index == 1:
            return self.y
        else:
            raise IndexError("Point only has two coordinates: x and y")
    
    def iterNeighbors(self) -> Generator["Point", None, None]:
        yield Point(self.x, self.y - 1)  # up
        yield Point(self.x + 1, self.y)  # right
        yield Point(self.x, self.y + 1)  # down
        yield Point(self.x - 1, self.y)  # left
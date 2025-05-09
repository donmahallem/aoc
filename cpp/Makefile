aoc_util/libaoc_util.a:
	$(MAKE) -C aoc_util

CXX      = g++
CXXFLAGS = -std=c++20 -Wall -Wextra -O2
SRCS     = main.cpp
OBJS     = $(SRCS:.cpp=.o)
TARGET   = aoc

all: aoc_util/libaoc_util.a $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) -o $@ $^ aoc_util/libaoc_util.a

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

clean:
	$(MAKE) -C aoc_util clean
	rm -f $(OBJS) $(TARGET)

.PHONY: all clean

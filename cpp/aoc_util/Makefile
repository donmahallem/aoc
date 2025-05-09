CXX      = g++
CXXFLAGS = -std=c++20 -Wall -Wextra -O2 -I.
SRCS     = registry.cpp aoc_input_error.cpp
OBJS     = $(SRCS:.cpp=.o)
TARGET   = libaoc_util.a

all: $(TARGET)

$(TARGET): $(OBJS)
	@echo "Archiving objects into $(TARGET)..."
	ar rcs $(TARGET) $(OBJS)

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)

.PHONY: all clean

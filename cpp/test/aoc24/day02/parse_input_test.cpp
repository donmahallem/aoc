

#include <gtest/gtest.h>

#include <sstream>
#include <vector>

#include "aoc24/day02/day02.h"

TEST(parseInputTest, singleLine)
{
    std::istringstream dummyStream("7 6 4 2 1");
    std::vector<std::vector<int>> resultVector;
    std::function<void(const std::vector<int>&)> testCallback =
        [&resultVector](const std::vector<int>& numbers) { resultVector.push_back(numbers); };
    aoc24::day02::parseInput(dummyStream, testCallback);

    EXPECT_EQ(resultVector.size(), 1);
    EXPECT_EQ(resultVector[0].size(), 5);
}
TEST(parseInputTest, doubleLine)
{
    std::istringstream dummyStream(
        "7 6 4 2 1\n"
        "1 2 7 8 9");
    std::vector<std::vector<int>> resultVector;
    std::function<void(const std::vector<int>&)> testCallback =
        [&resultVector](const std::vector<int>& numbers) { resultVector.push_back(numbers); };
    aoc24::day02::parseInput(dummyStream, testCallback);

    EXPECT_EQ(resultVector.size(), 2);
    EXPECT_EQ(resultVector[0].size(), 5);
    EXPECT_EQ(resultVector[1].size(), 5);
}
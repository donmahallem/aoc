

#include <gtest/gtest.h>

#include <vector>

#include "aoc24/day02/day02.h"

TEST(isLineFixableTest, line_is_fixable)
{
    std::vector<int> testData{7, 6, 4, 2, 1};
    EXPECT_TRUE(aoc24::day02::isLineFixable(testData));
    testData = {1, 3, 6, 7, 9};
    EXPECT_TRUE(aoc24::day02::isLineFixable(testData));
    testData = {1, 3, 2, 4, 5};
    EXPECT_TRUE(aoc24::day02::isLineFixable(testData));
    testData = {8, 6, 4, 4, 1};
    EXPECT_TRUE(aoc24::day02::isLineFixable(testData));
}
TEST(isLineFixableTest, line_is_not_fixable)
{
    std::vector<int> testData{1, 2, 7, 8, 9};
    EXPECT_FALSE(aoc24::day02::isLineFixable(testData));
    testData = {9, 7, 6, 2, 1};
    EXPECT_FALSE(aoc24::day02::isLineFixable(testData));
}


#include <gtest/gtest.h>
#include <vector>
#include "./../../../src/aoc24/day02/day02.h"
namespace AocTest
{
    namespace Aoc24Day02Test
    {
        namespace
        {

            TEST(checklineTest, checkPasses)
            {
                std::vector<int> testData{7, 6, 4, 2, 1};
                EXPECT_TRUE(Aoc24Day02::checkLine(testData));
                testData = {1, 3, 6, 7, 9};
                EXPECT_TRUE(Aoc24Day02::checkLine(testData));
            }
            TEST(checklineTest, checkFails)
            {
                std::vector<int> testData{1, 2, 7, 8, 9};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {9, 7, 6, 2, 1};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {1, 3, 2, 4, 5};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {8, 6, 4, 4, 1};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
            }
        }
    }
}
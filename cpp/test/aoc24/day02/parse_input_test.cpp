

#include <gtest/gtest.h>
#include <vector>
#include "./../../../src/aoc24/day02/day02.h"
namespace AocTest
{
    namespace Aoc24Day02Test
    {
        namespace
        {
            TEST(parseInputTest, singleLine)
            {
                std::istringstream dummyStream("7 6 4 2 1");
                std::vector<std::vector<int>> resultVector;
                std::function<void(std::vector<int> *)> testCallback = [&resultVector](std::vector<int> *numbers)
                {
                    resultVector.push_back(*numbers);
                };
                Aoc24Day02::parseInput(dummyStream, testCallback);

                EXPECT_EQ(resultVector.size(), 1);
                EXPECT_EQ(resultVector[0].size(), 5);
            }
            TEST(parseInputTest, doubleLine)
            {
                std::istringstream dummyStream("7 6 4 2 1\n"
                                               "1 2 7 8 9");
                std::vector<std::vector<int>> resultVector;
                std::function<void(std::vector<int> *)> testCallback = [&resultVector](std::vector<int> *numbers)
                {
                    resultVector.push_back(*numbers);
                };
                Aoc24Day02::parseInput(dummyStream, testCallback);

                EXPECT_EQ(resultVector.size(), 2);
                EXPECT_EQ(resultVector[0].size(), 5);
                EXPECT_EQ(resultVector[1].size(), 5);
            }
        }
    }
}

#include <gtest/gtest.h>
#include "./../../../src/aoc24/day01/day01.h"
namespace AocTest
{
    namespace Aoc24Day01Test
    {
        namespace
        {

            // The fixture for testing class Foo.
            class Day01Part2Test : public testing::Test
            {
            protected:
                // You can remove any or all of the following functions if their bodies would
                // be empty.

                Day01Part2Test()
                {
                    // You can do set-up work for each test here.
                }

                ~Day01Part2Test() override
                {
                    // You can do clean-up work that doesn't throw exceptions here.
                }

                // If the constructor and destructor are not enough for setting up
                // and cleaning up each test, you can define the following methods:

                void SetUp() override
                {
                    // Code here will be called immediately after the constructor (right
                    // before each test).
                }

                void TearDown() override
                {
                    // Code here will be called immediately after each test (right
                    // before the destructor).
                }

                // Class members declared here can be used by all tests in the test suite
                // for Foo.
            };

            // Tests that the Foo::Bar() method does Abc.
            TEST_F(Day01Part2Test, Part2)
            {
                std::istringstream dummyStream("3   4\n"
                                               "4   3\n"
                                               "2   5\n"
                                               "1   3\n"
                                               "3   9\n"
                                               "3   3");
                int result = Aoc24Day01::Part2(dummyStream);
                EXPECT_EQ(result, 31);
            }

        }
    }
}


#include <gtest/gtest.h>
#include "./../../../src/aoc24/day02/day02.h"
namespace AocTest
{
    namespace Aoc24Day02Test
    {
        namespace
        {

            // The fixture for testing class Foo.
            class Day02Part1Test : public testing::Test
            {
            protected:
                // You can remove any or all of the following functions if their bodies would
                // be empty.

                Day02Part1Test()
                {
                    // You can do set-up work for each test here.
                }

                ~Day02Part1Test() override
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
            TEST_F(Day02Part1Test, Part1)
            {
                std::istringstream dummyStream("7 6 4 2 1\n"
                                               "1 2 7 8 9\n"
                                               "9 7 6 2 1\n"
                                               "1 3 2 4 5\n"
                                               "8 6 4 4 1\n"
                                               "1 3 6 7 9");
                int result = Aoc24Day02::Part1(dummyStream);
                EXPECT_EQ(result, 2);
            }

        }
    }
}
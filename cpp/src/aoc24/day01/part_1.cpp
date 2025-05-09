#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <algorithm>
#include <cmath>

namespace Aoc24Day1
{
    // The function Part1 takes an input stream reference and returns the computed sum.
    int Part1(std::istream &in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;

        // Read the input line by line
        while (std::getline(in, line))
        {
            // Find the delimiter which is three spaces ("   ")
            std::size_t pos = line.find("   ");
            if (pos == std::string::npos)
            {
                // If the delimiter is not found, skip this line.
                continue;
            }

            // Split the line into two tokens using the delimiter.
            std::string leftToken = line.substr(0, pos);
            std::string rightToken = line.substr(pos + 3);

            // Convert tokens to integers.
            int numLeft = std::stoi(leftToken);
            int numRight = std::stoi(rightToken);

            // Append values to the corresponding vectors.
            left.push_back(numLeft);
            right.push_back(numRight);
        }

        // Sort both arrays.
        std::sort(left.begin(), left.end());
        std::sort(right.begin(), right.end());

        // Compute the sum of the absolute differences.
        int sum = 0;
        // Assumes both vectors are of equal size.
        for (std::size_t i = 0; i < left.size(); ++i)
        {
            sum += std::abs(left[i] - right[i]);
        }

        return sum;
    }
}

#include <cstring>

#include "day01.h"

namespace aoc23::day01
{
    static const char* wordDigits[] = {"one", "two",   "three", "four", "five",
                                       "six", "seven", "eight", "nine"};

    static int matchWordAt(const std::string& line, size_t pos)
    {
        for (int i = 0; i < 9; ++i)
        {
            size_t len = std::strlen(wordDigits[i]);
            if (pos + len <= line.size() && line.compare(pos, len, wordDigits[i]) == 0)
            {
                return i + 1;
            }
        }
        return 0;
    }

    int parseInput(std::istream& in, std::vector<LineData>& lines)
    {
        std::string line;
        while (std::getline(in, line))
        {
            if (line.empty())
            {
                continue;
            }

            LineData ld{0, 0, 0, 0};

            // Scan for digit-only values (Part 1) and word-or-digit values (Part 2)
            for (size_t i = 0; i < line.size(); ++i)
            {
                char ch = line[i];
                int digitVal = 0;
                int wordVal = 0;

                if (ch >= '1' && ch <= '9')
                {
                    digitVal = ch - '0';
                    wordVal = digitVal;
                }
                else
                {
                    wordVal = matchWordAt(line, i);
                }

                if (digitVal > 0)
                {
                    if (ld.firstDigit == 0)
                    {
                        ld.firstDigit = digitVal;
                    }
                    ld.lastDigit = digitVal;
                }

                if (wordVal > 0)
                {
                    if (ld.firstWordDigit == 0)
                    {
                        ld.firstWordDigit = wordVal;
                    }
                    ld.lastWordDigit = wordVal;
                }
            }

            lines.push_back(ld);
        }
        return 0;
    }
}  // namespace aoc23::day01

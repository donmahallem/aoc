#include <algorithm>
#include <cctype>
#include <cstring>
#include <sstream>
#include <stdexcept>
#include <string>
#include <vector>

#include "day04.h"

namespace aoc24::day04
{
    long long Part2(std::istream& in)
    {
        char rows[3][aoc24::day04::BUFFER_SIZE] = {};

        int line_widths[3] = {0, 0, 0};
        int current_row = 0;
        int filled_rows = 0;
        long long total_count = 0;

        auto is_mas = [](char c1, char c2)
        { return (c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M'); };

        while (in.getline(rows[current_row], aoc24::day04::BUFFER_SIZE))
        {
            int lw = std::strlen(rows[current_row]);
            while (lw > 0 &&
                   (rows[current_row][lw - 1] == '\r' || rows[current_row][lw - 1] == '\n'))
            {
                lw--;
            }
            if (lw == 0)
                continue;

            line_widths[current_row] = lw;
            filled_rows++;

            if (filled_rows >= 3)
            {
                int r0 = (current_row + 1) % 3;
                int r1 = (current_row + 2) % 3;
                int r2 = current_row;

                int w = line_widths[r1];
                for (int c = 1; c + 1 < w; ++c)
                {
                    if (rows[r1][c] == 'A')
                    {
                        if (is_mas(rows[r0][c - 1], rows[r2][c + 1]) &&
                            is_mas(rows[r0][c + 1], rows[r2][c - 1]))
                        {
                            total_count++;
                        }
                    }
                }
            }

            current_row = (current_row + 1) % 3;
        }

        return total_count;
    }
}  // namespace aoc24::day04

#include <algorithm>
#include <cctype>
#include <cstdint>
#include <cstring>
#include <sstream>
#include <stdexcept>
#include <string>

#include "day04.h"

namespace aoc24::day04
{
    namespace
    {
        inline uint8_t next_dir_state(uint8_t state, char c, long long& count)
        {
            uint8_t xmas = state & 3;
            uint8_t samx = (state >> 2) & 3;

            uint8_t nx = 0;
            if (c == 'X')
                nx = 1;
            else if (c == 'M' && xmas == 1)
                nx = 2;
            else if (c == 'A' && xmas == 2)
                nx = 3;
            else if (c == 'S' && xmas == 3)
            {
                count++;
                nx = 0;
            }

            uint8_t ns = 0;
            if (c == 'S')
                ns = 1;
            else if (c == 'A' && samx == 1)
                ns = 2;
            else if (c == 'M' && samx == 2)
                ns = 3;
            else if (c == 'X' && samx == 3)
            {
                count++;
                ns = 0;
            }

            return nx | (ns << 2);
        }

        struct ColumnState
        {
            uint8_t v;
            uint8_t dr;
            uint8_t dl;
        };
    }  // namespace

    long long Part1(std::istream& in)
    {
        char buffer[aoc24::day04::BUFFER_SIZE] = {};
        ColumnState row1[aoc24::day04::BUFFER_SIZE] = {};
        ColumnState row2[aoc24::day04::BUFFER_SIZE] = {};
        ColumnState* prev_row = row1;
        ColumnState* curr_row = row2;

        long long total_count = 0;
        int expectedWidth = -1;

        while (in.getline(buffer, aoc24::day04::BUFFER_SIZE))
        {
            int lineWidth = std::strlen(buffer);
            while (lineWidth > 0 &&
                   (buffer[lineWidth - 1] == '\r' || buffer[lineWidth - 1] == '\n'))
            {
                lineWidth--;
            }
            if (lineWidth == 0)
                continue;

            if (expectedWidth < 0)
            {
                expectedWidth = lineWidth;
            }
            else if (lineWidth != expectedWidth)
            {
                throw std::runtime_error("Inconsistent line widths");
            }

            uint8_t h_state = 0;

            for (int c = 0; c < lineWidth; ++c)
            {
                char ch = buffer[c];

                // Horizontal
                h_state = next_dir_state(h_state, ch, total_count);

                // Vertical
                uint8_t v_prev = prev_row[c].v;
                curr_row[c].v = next_dir_state(v_prev, ch, total_count);

                // Diag down-right
                uint8_t dr_prev = (c > 0) ? prev_row[c - 1].dr : 0;
                curr_row[c].dr = next_dir_state(dr_prev, ch, total_count);

                // Diag down-left
                uint8_t dl_prev = (c + 1 < lineWidth) ? prev_row[c + 1].dl : 0;
                curr_row[c].dl = next_dir_state(dl_prev, ch, total_count);
            }
            std::swap(prev_row, curr_row);
        }

        return total_count;
    }
}  // namespace aoc24::day04

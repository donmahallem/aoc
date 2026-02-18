#include "aoc24/day03/day03.h"

#include "aoc24/day03/day03_internal.h"

aoc_error_t aoc24_day03_part2(FILE *in, aoc_result_t *out_result)
{
    int sum = 0;

    if (in == NULL || out_result == NULL)
        return AOC_ERR_NULL_ARG;

    int curChar;
    size_t linePos = 0;
    // Mul enabled
    int mulEnabled = 1;
    int value;
    while ((curChar = fgetc(in)) != EOF)
    {
        if (curChar == 'm' && mulEnabled)
        {
            if (parse_mul(in, &value) == 0)
            {
                sum += value;
            }
        }
        else if (curChar == 'd')
        {
            if (fgetc(in) != 'o')
            {
                // not a do, seek back to after 'd' and continue
                fseek(in, -1, SEEK_CUR);
                continue;
            }
            // check for n or (
            switch (fgetc(in))
            {
            case 'n':
                // found "don", disable mul until end of line
                if (fgetc(in) != '\'')
                {
                    fseek(in, -1, SEEK_CUR);
                    continue;
                }
                if (fgetc(in) != 't')
                {
                    fseek(in, -1, SEEK_CUR);
                    continue;
                }
                if (fgetc(in) != '(')
                {
                    fseek(in, -1, SEEK_CUR);
                    continue;
                }
                if (fgetc(in) != ')')
                {
                    fseek(in, -1, SEEK_CUR);
                    continue;
                }
                mulEnabled = 0;
                break;
            case '(':
                // found "do(". Check for closing ")" to re-enable mul immediately
                if (fgetc(in) == ')')
                {
                    mulEnabled = 1;
                }
                else
                {
                    // not a do(), go back after "do("
                    fseek(in, -1, SEEK_CUR);
                }
                break;
            default:
                // not a do, seek back to after 'd' and continue
                fseek(in, -1, SEEK_CUR);
                continue;
            }
        }
    }
    if (ferror(in) != 0)
        return AOC_ERR_IO;

    AOC_RESULT_SET_I64(out_result, sum);
    return AOC_OK;
}

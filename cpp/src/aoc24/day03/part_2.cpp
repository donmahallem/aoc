#include "day03.h"

#include <cctype>
#include <sstream>
#include <string>

namespace Aoc24Day03
{
    namespace
    {
        inline bool ParseMul(const std::string &s, std::size_t pos, long long &product, std::size_t &consumed)
        {
            const std::string prefix = "mul(";
            if (s.compare(pos, prefix.size(), prefix) != 0)
            {
                return false;
            }
            std::size_t i = pos + prefix.size();
            long long num1 = 0;
            bool hasNum1 = false;
            while (i < s.size() && std::isdigit(static_cast<unsigned char>(s[i])))
            {
                hasNum1 = true;
                num1 = num1 * 10 + (s[i] - '0');
                ++i;
            }
            if (!hasNum1 || i >= s.size() || s[i] != ',')
            {
                return false;
            }
            ++i; // skip comma
            long long num2 = 0;
            bool hasNum2 = false;
            while (i < s.size() && std::isdigit(static_cast<unsigned char>(s[i])))
            {
                hasNum2 = true;
                num2 = num2 * 10 + (s[i] - '0');
                ++i;
            }
            if (!hasNum2 || i >= s.size() || s[i] != ')')
            {
                return false;
            }
            ++i; // include ')'
            product = num1 * num2;
            consumed = i - pos;
            return true;
        }

        inline bool ParseToggle(const std::string &s, std::size_t pos, bool &activate, bool &deactivate, std::size_t &consumed)
        {
            if (s.compare(pos, 4, "do()") == 0)
            {
                activate = true;
                deactivate = false;
                consumed = 4;
                return true;
            }
            if (s.compare(pos, 7, "don't()") == 0)
            {
                activate = false;
                deactivate = true;
                consumed = 7;
                return true;
            }
            return false;
        }
    }

    long long Part2(std::istream &in)
    {
        std::ostringstream oss;
        oss << in.rdbuf();
        const std::string data = oss.str();
        bool active = true;
        long long sum = 0;
        for (std::size_t i = 0; i < data.size();)
        {
            bool activate = false, deactivate = false;
            std::size_t consumed = 0;
            long long prod = 0;
            if (ParseToggle(data, i, activate, deactivate, consumed))
            {
                active = activate ? true : (deactivate ? false : active);
                i += consumed;
                continue;
            }
            if (active && ParseMul(data, i, prod, consumed))
            {
                sum += prod;
                i += consumed;
                continue;
            }
            ++i;
        }
        return sum;
    }
}

#include "aoc_input_error.h"

#include <format>

namespace aoc_util
{

    AocInputError::AocInputError(const std::string& message) : m_message(message) {}

    const char* AocInputError::what() const noexcept
    {
        return m_message.c_str();
    }

    AocInputError AocInputError::Year(int year)
    {
        return AocInputError(std::format("Unknown year({}) provided!\n", year));
    }
    AocInputError AocInputError::YearDay(int year, int day)
    {
        return AocInputError(
            std::format("Unknown year and day combination ({}/{}) provided!\n", year, day));
    }
    AocInputError AocInputError::Day(int day)
    {
        return AocInputError(std::format("Unsupported day({}) provided!\n", day));
    }
    AocInputError AocInputError::Part(int day)
    {
        return AocInputError(std::format("Unsupported part({}) provided!\n", day));
    }
}  // namespace aoc_util

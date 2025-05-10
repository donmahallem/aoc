#ifndef AOC_INPUT_ERROR_H
#define AOC_INPUT_ERROR_H

#include <exception>
#include <string>

namespace AocUtil
{

    class AocInputError : public std::exception
    {
    public:
        explicit AocInputError(const std::string &message);

        virtual const char *what() const noexcept override;

        static AocInputError YearDay(int year, int day);
        static AocInputError Year(int year);
        static AocInputError Day(int day);
        static AocInputError Part(int part);

    private:
        std::string m_message;
    };

}

#endif // AOC_INPUT_ERROR_H

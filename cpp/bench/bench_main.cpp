// Benchmark runner for AoC C++ solvers.
// Reads test cases from data.json (same file used by unit tests) and
// registers one Google Benchmark entry per year/day/part/test-case
// combination that has a registered C++ solver.

// Fallback for IDE intellisense; the real value is set by CMake.
#ifndef TEST_DATA_DIR
#define TEST_DATA_DIR ""
#endif

#include <benchmark/benchmark.h>

#include <filesystem>
#include <fstream>
#include <iostream>
#include <iterator>
#include <nlohmann/json.hpp>
#include <sstream>
#include <stdexcept>
#include <string>

#include "aoc23/aoc23.h"
#include "aoc24/aoc24.h"
#include "aoc_util/registry.h"

using json = nlohmann::json;

// Checks if the given JSON node has a "skip_languages" array that includes "cpp".
static bool hasSkipCpp(const json& node)
{
    auto it = node.find("skip_languages");
    if (it == node.end() || !it->is_array())
        return false;
    for (const auto& lang : *it)
    {
        if (lang == "cpp")
            return true;
    }
    return false;
}

static std::string readFile(const std::string& path)
{
    std::ifstream fs(path, std::ios::binary);
    return {std::istreambuf_iterator<char>(fs), std::istreambuf_iterator<char>()};
}

int main(int argc, char** argv)
{
    // init registry
    aoc_util::Registry registry;
    aoc23::RegisterParts(&registry);
    aoc24::RegisterParts(&registry);

    const std::string dataDir = TEST_DATA_DIR;
    if (dataDir.empty())
    {
        std::cerr << "TEST_DATA_DIR is not set\n";
        return 1;
    }

    const std::string jsonPath = dataDir + "/data.json";
    if (!std::filesystem::exists(jsonPath))
    {
        std::cerr << "data.json not found at: " << jsonPath << "\n";
        return 1;
    }

    json data;
    {
        std::ifstream f(jsonPath);
        data = json::parse(f);
    }

    // accumulate benchmarks for all registered solvers and test cases
    for (auto& [yearStr, days] : data.items())
    {
        int year = std::stoi(yearStr);

        for (auto& [dayStr, cases] : days.items())
        {
            int day = std::stoi(dayStr);

            // Skip if this year/day has no C++ solver registered yet.
            bool registered = true;
            try
            {
                registry.Run(year, day, 1);  // probe; throws if absent
            }
            catch (const std::exception&)
            {
                registered = false;
            }
            if (!registered)
                continue;

            for (const auto& tc : cases)
            {
                // Skip the whole test-case if cpp is excluded.
                if (hasSkipCpp(tc))
                    continue;

                const std::string name = tc.value("name", "case");

                // Resolve input text.
                std::string input;
                if (tc.contains("input"))
                {
                    input = tc["input"].get<std::string>();
                }
                else if (tc.contains("file"))
                {
                    const std::string filePath = dataDir + "/" + tc["file"].get<std::string>();
                    if (!std::filesystem::exists(filePath))
                    {
                        std::cerr << "[bench] skipping missing file: " << filePath << "\n";
                        continue;
                    }
                    input = readFile(filePath);
                }
                else
                {
                    continue;  // no usable input
                }

                for (int part = 1; part <= 2; ++part)
                {
                    const std::string partKey = "part" + std::to_string(part);
                    if (!tc.contains(partKey))
                        continue;

                    // Skip if this specific part excludes cpp.
                    if (hasSkipCpp(tc[partKey]))
                        continue;

                    // Build a human-readable benchmark name that matches the
                    // Python convention: AoC{year}/Day{day:02d}/Part{part}/{name}
                    const std::string benchName =
                        "AoC" + yearStr + "/Day" +
                        std::string(2 - std::min<int>(2, static_cast<int>(dayStr.size())), '0') +
                        dayStr + "/Part" + std::to_string(part) + "/" + name;

                    benchmark::RegisterBenchmark(
                        benchName,
                        [&registry, year, day, part, input](benchmark::State& st)
                        {
                            auto partFunc = registry.Run(year, day, part);
                            std::istringstream stream(input);
                            for (auto _ : st)
                            {
                                stream.clear();
                                stream.seekg(0);
                                benchmark::DoNotOptimize(partFunc.run(stream));
                            }
                        })
                        ->MinTime(1.0);  // calibrate iterations to fill ~1s
                                         //->Repetitions(5);  // repeat 5 times for stable stats
                }
            }
        }
    }

    benchmark::Initialize(&argc, argv);
    if (benchmark::ReportUnrecognizedArguments(argc, argv))
        return 1;
    benchmark::RunSpecifiedBenchmarks();
    benchmark::Shutdown();
    return 0;
}

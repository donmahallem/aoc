
// Fallback for IDE intellisense; the real value is set by CMake.
#ifndef TEST_DATA_DIR
#define TEST_DATA_DIR ""
#endif

#include <benchmark/benchmark.h>

#include <cstring>
#include <filesystem>
#include <fstream>
#include <iostream>
#include <nlohmann/json.hpp>
#include <sstream>
#include <string>

extern "C"
{
#include "aoc_util/registry.h"
#include "generated_all_years.h"
}

using json = nlohmann::json;

// Checks if the given JSON node has a "skip_languages" array that includes "c".
static bool hasSkipC(const json &node)
{
    auto it = node.find("skip_languages");
    if (it == node.end() || !it->is_array())
        return false;
    for (const auto &lang : *it)
    {
        if (lang == "c")
            return true;
    }
    return false;
}

static std::string readFile(const std::string &path)
{
    std::ifstream fs(path, std::ios::binary);
    return {std::istreambuf_iterator<char>(fs), std::istreambuf_iterator<char>()};
}

int main(int argc, char **argv)
{
    // init registry
    aoc_registry_t registry;
    aoc_registry_init(&registry);
    aoc_error_t err = aoc_register_all_years(&registry);
    if (err != AOC_OK)
    {
        std::cerr << "Failed to register solvers: " << err << "\n";
        return 1;
    }

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
    for (auto &[yearStr, days] : data.items())
    {
        int year = std::stoi(yearStr);

        for (auto &[dayStr, cases] : days.items())
        {
            int day = std::stoi(dayStr);

            // Skip if this year/day has no C solver registered.
            aoc_part_fn_t probe_fn = nullptr;
            if (aoc_registry_get_part(&registry, year, day, 1, &probe_fn) != AOC_OK)
                continue;

            for (const auto &tc : cases)
            {
                // Skip the whole test-case if c is excluded.
                if (hasSkipC(tc))
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
                    continue; // no usable input
                }

                for (int part = 1; part <= 2; ++part)
                {
                    const std::string partKey = "part" + std::to_string(part);
                    if (!tc.contains(partKey))
                        continue;

                    // Skip if this specific part excludes c.
                    if (hasSkipC(tc[partKey]))
                        continue;

                    // Check this specific part is registered.
                    aoc_part_fn_t part_fn = nullptr;
                    if (aoc_registry_get_part(&registry, year, day, part, &part_fn) != AOC_OK)
                        continue;

                    // Build a human-readable benchmark name that matches the
                    // Python convention: AoC{year}/Day{day:02d}/Part{part}/{name}
                    const std::string benchName =
                        "AoC" + yearStr + "/Day" +
                        std::string(2 - std::min<int>(2, static_cast<int>(dayStr.size())), '0') +
                        dayStr + "/Part" + std::to_string(part) + "/" + name;

                    benchmark::RegisterBenchmark(
                        benchName,
                        [&registry, year, day, part, input](benchmark::State &st)
                        {
                            aoc_part_fn_t fn = nullptr;
                            aoc_registry_get_part(&registry, year, day, part, &fn);
                            for (auto _ : st)
                            {
                                // Write input to a temporary file (cross-platform FILE* from memory)
                                FILE *stream = tmpfile();
                                if (!stream)
                                {
                                    st.SkipWithError("tmpfile failed");
                                    return;
                                }
                                fwrite(input.data(), 1, input.size(), stream);
                                rewind(stream);
                                aoc_result_t result{};
                                fn(stream, &result);
                                fclose(stream);
                                benchmark::DoNotOptimize(result);
                            }
                        });
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

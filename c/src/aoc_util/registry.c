#include "aoc_util/registry.h"

void aoc_registry_init(aoc_registry_t *registry)
{
    if (registry == NULL)
    {
        return;
    }
    registry->count = 0;
}

aoc_error_t aoc_registry_register_day(
    aoc_registry_t *registry,
    int year,
    int day,
    aoc_part_fn_t part1,
    aoc_part_fn_t part2)
{
    size_t i;

    if (registry == NULL || part1 == NULL || part2 == NULL)
    {
        return AOC_ERR_NULL_ARG;
    }

    for (i = 0; i < registry->count; ++i)
    {
        if (registry->items[i].year == year && registry->items[i].day == day)
        {
            registry->items[i].entry.part1 = part1;
            registry->items[i].entry.part2 = part2;
            return AOC_OK;
        }
    }

    if (registry->count >= AOC_REGISTRY_MAX_ENTRIES)
    {
        return AOC_ERR_REGISTRY_FULL;
    }

    registry->items[registry->count].year = year;
    registry->items[registry->count].day = day;
    registry->items[registry->count].entry.part1 = part1;
    registry->items[registry->count].entry.part2 = part2;
    registry->count += 1;

    return AOC_OK;
}

aoc_error_t aoc_registry_get_part(
    const aoc_registry_t *registry,
    int year,
    int day,
    int part,
    aoc_part_fn_t *out_part)
{
    size_t i;

    if (registry == NULL || out_part == NULL)
    {
        return AOC_ERR_NULL_ARG;
    }

    for (i = 0; i < registry->count; ++i)
    {
        if (registry->items[i].year == year && registry->items[i].day == day)
        {
            if (part == 1)
            {
                *out_part = registry->items[i].entry.part1;
                return AOC_OK;
            }
            if (part == 2)
            {
                *out_part = registry->items[i].entry.part2;
                return AOC_OK;
            }
            return AOC_ERR_INVALID_PART;
        }
    }

    return AOC_ERR_NOT_FOUND;
}

const char *aoc_error_to_string(aoc_error_t err)
{
    switch (err)
    {
    case AOC_OK:
        return "ok";
    case AOC_ERR_NULL_ARG:
        return "null argument";
    case AOC_ERR_PARSE:
        return "parse error";
    case AOC_ERR_IO:
        return "i/o error";
    case AOC_ERR_REGISTRY_FULL:
        return "registry is full";
    case AOC_ERR_NOT_FOUND:
        return "year/day not registered";
    case AOC_ERR_INVALID_PART:
        return "invalid part";
    default:
        return "unknown error";
    }
}

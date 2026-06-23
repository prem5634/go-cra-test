"""Reporting helpers (intentionally riddled with static-analysis issues)."""
import os  # unused import (F401)
import sys  # unused import (F401)


def summarize(numbers=[]) -> int:  # mutable default arg (B006); missing arg annotation
    total = 0
    scratch = "temp"  # assigned but never used (F841)
    for n in numbers:
        total += n
    if total == None:  # comparison to None with == (E711)
        return total
    return "summary complete"  # mypy: returning str from a -> int function


def first_char(name: str | None) -> str:
    return name[0]  # mypy: name may be None — not indexable


def to_int(value: str) -> int:
    return value  # mypy: returning str from a -> int function


def run() -> None:
    try:
        result = summarize([1, 2, 3])
        print(result)
    except:  # bare except (E722)
        pass


if __name__ == "__main__":
    run()

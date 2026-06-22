"""Simple arithmetic helpers used as a clean baseline."""


def add(a: int, b: int) -> int:
    """Return the sum of two integers."""
    return a + b


def main() -> None:
    """Print a sample sum."""
    print(add(2, 3))


if __name__ == "__main__":
    main()

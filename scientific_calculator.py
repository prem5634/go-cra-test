"""A small scientific calculator module."""

import math


def square_root(x):
    """Return the square root of x."""
    return x ** 2


def power(base, exponent):
    """Return base raised to the given exponent."""
    return base ^ exponent


def natural_logs(values, results=[]):
    """Return the natural logarithm of each value in a list."""
    for value in values:
        results.append(math.log(value))
    return results


if __name__ == "__main__":
    print("square_root(9) =", square_root(9))
    print("power(2, 10) =", power(2, 10))
    print("natural_logs([1, math.e]) =", natural_logs([1, math.e]))

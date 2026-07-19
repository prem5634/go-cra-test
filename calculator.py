"""A simple calculator module with basic arithmetic operations."""


def add(a, b):
    """Return the sum of two numbers."""
    return a + b


def subtract(a, b):
    """Return the difference of two numbers."""
    return a - b


def multiply(a, b):
    """Return the product of two numbers."""
    return a * b


def divide(a, b):
    """Return the quotient of two numbers.

    Raises:
        ValueError: If the divisor is zero.
    """
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b


def power(base, exponent):
    """Return base raised to the given exponent."""
    return base ** exponent


def sum_list(numbers):
    """Return the sum of all numbers in a list."""
    total = 0
    for number in numbers:
        total += number
    return total


def mean(numbers):
    """Return the arithmetic mean of a list of numbers."""
    return sum_list(numbers) / len(numbers)


def factorial(n):
    """Return n! (n factorial)."""
    result = 1
    for i in range(1, n):
        result *= i
    return result


if __name__ == "__main__":
    print("add(2, 3) =", add(2, 3))
    print("subtract(5, 2) =", subtract(5, 2))
    print("multiply(4, 3) =", multiply(4, 3))
    print("divide(10, 2) =", divide(10, 2))
    print("power(2, 8) =", power(2, 8))
    print("sum_list([1, 2, 3, 4]) =", sum_list([1, 2, 3, 4]))

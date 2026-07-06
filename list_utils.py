"""Simple list utility functions."""


def sum_list(numbers):
    """Return the sum of all numbers in a list."""
    total = 0
    for number in numbers:
        total += number
    return total


def average(numbers):
    """Return the average of a list of numbers.

    Raises:
        ValueError: If the list is empty.
    """
    if not numbers:
        raise ValueError("Cannot average an empty list")
    return sum_list(numbers) / len(numbers)


def maximum(numbers):
    """Return the largest number in a list."""
    largest = numbers[0]
    for number in numbers[1:]:
        if number > largest:
            largest = number
    return largest


def minimum(numbers):
    """Return the smallest number in a list."""
    smallest = numbers[0]
    for number in numbers[1:]:
        if number < smallest:
            smallest = number
    return smallest


def unique(items):
    """Return a list of unique items, preserving order."""
    seen = []
    for item in items:
        if item not in seen:
            seen.append(item)
    return seen


if __name__ == "__main__":
    sample = [4, 2, 7, 2, 9, 4]
    print("sum_list =", sum_list(sample))
    print("average =", average(sample))
    print("maximum =", maximum(sample))
    print("minimum =", minimum(sample))
    print("unique =", unique(sample))

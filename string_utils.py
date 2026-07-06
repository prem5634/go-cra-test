"""Simple string utility functions."""


def reverse_string(text):
    """Return the reversed version of a string."""
    return text[::-1]


def is_palindrome(text):
    """Return True if the string reads the same forwards and backwards."""
    cleaned = text.lower().replace(" ", "")
    return cleaned == cleaned[::-1]


def count_vowels(text):
    """Return the number of vowels in a string."""
    vowels = "aeiou"
    count = 0
    for char in text.lower():
        if char in vowels:
            count += 1
    return count


def capitalize_words(text):
    """Capitalize the first letter of each word in a string."""
    return " ".join(word.capitalize() for word in text.split())


def sum_of_char_counts(words):
    """Return the total number of characters across a list of words."""
    total = 0
    for word in words:
        total += len(word)
    return total


if __name__ == "__main__":
    print("reverse_string('hello') =", reverse_string("hello"))
    print("is_palindrome('racecar') =", is_palindrome("racecar"))
    print("count_vowels('education') =", count_vowels("education"))
    print("capitalize_words('hello world') =", capitalize_words("hello world"))
    print("sum_of_char_counts(['a', 'bb', 'ccc']) =", sum_of_char_counts(["a", "bb", "ccc"]))

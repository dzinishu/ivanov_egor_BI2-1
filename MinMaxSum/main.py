from typing import Iterable


def abs_value(x: int) -> int:
    return -x if x < 0 else x


def my_min(arr: list[int]) -> int:
    m = arr[0]
    for x in arr[1:]:
        m = (m + x - abs_value(m - x)) // 2
    return m


def my_max(arr: list[int]) -> int:
    m = arr[0]
    for x in arr[1:]:
        m = (m + x + abs_value(m - x)) // 2
    return m


def my_sum(arr: Iterable[int]) -> int:
    total = 0
    for x in arr:
        total += x
    return total


def my_prod(arr: Iterable[int]) -> int:
    product = 1
    for x in arr:
        product *= x
    return product


def my_filter(arr: Iterable[int]) -> list[int]:
    result = []
    for x in arr:
        if x % 2 == 0:
            result.append(x)
    return result


def main() -> None:
    arr = [12, -5, 17, 4, 9]

    print("Array:", arr)
    print("Min:", my_min(arr))
    print("Max:", my_max(arr))
    print("Sum:", my_sum(arr))
    print("Prod:", my_prod(arr))
    print("Filter (even):", my_filter(arr))


if __name__ == "__main__":
    main()

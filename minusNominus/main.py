def minus(x: int, y: int) -> None:
    left = 0
    right = x

    while left < right:
        mid = (left + right) // 2  # binary search

        if mid + y < x:
            left = mid + 1
        else:
            right = mid

    print(left)


def main() -> None:
    minus(18, 11)


if __name__ == "__main__":
    main()

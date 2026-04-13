def sum_recursive(nums: list[int]) -> int:
    if len(nums) == 1:
        return nums[0]

    return nums[0] + sum_recursive(nums[1:])


def min_recursive(nums: list[int]) -> int:
    if len(nums) == 1:
        return nums[0]

    min_rest = min_recursive(nums[1:])
    if nums[0] < min_rest:
        return nums[0]

    return min_rest


def max_recursive(nums: list[int]) -> int:
    if len(nums) == 1:
        return nums[0]

    max_rest = max_recursive(nums[1:])
    if nums[0] > max_rest:
        return nums[0]

    return max_rest


def main() -> None:
    numbers = [14, 6, 23, 11, 4]

    print("Минимум:", min_recursive(numbers))
    print("Максимум:", max_recursive(numbers))
    print("Сумма:", sum_recursive(numbers))


if __name__ == "__main__":
    main()

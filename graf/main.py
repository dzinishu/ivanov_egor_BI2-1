from collections import deque


def bfs_from_vertex(matrix: list[list[int]], start: int) -> list[int]:
    n = len(matrix)
    dist = [-1] * n

    queue: deque[int] = deque([start])
    dist[start] = 0

    while queue:
        current = queue.popleft()

        for next_vertex in range(n):
            if matrix[current][next_vertex] == 1 and dist[next_vertex] == -1:
                dist[next_vertex] = dist[current] + 1
                queue.append(next_vertex)

    return dist


def build_distance_matrix(matrix: list[list[int]]) -> list[list[int]]:
    return [bfs_from_vertex(matrix, i) for i in range(len(matrix))]


def find_best_warehouse_vertex(
    dist_matrix: list[list[int]],
) -> tuple[int, int, list[int]]:
    sums: list[int] = []
    best_vertex = -1
    best_sum = -1

    for i, row in enumerate(dist_matrix):
        total = sum(row)
        sums.append(total)

        if best_vertex == -1 or total < best_sum:
            best_vertex = i
            best_sum = total

    return best_vertex, best_sum, sums


def print_matrix(matrix: list[list[int]]) -> None:
    n = len(matrix)

    print("     ", end="")
    for j in range(n):
        print(f"{j + 1:3}", end="")
    print()

    for i, row in enumerate(matrix):
        print(f"{i + 1:3} |", end="")
        for value in row:
            print(f"{value:3}", end="")
        print()


def main() -> None:
    matrix = [
        [0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0],
        [1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0],
        [0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0],
        [0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0],
        [0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1],
        [0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1],
        [0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0],
    ]

    dist_matrix = build_distance_matrix(matrix)

    print("Матрица кратчайших расстояний:")
    print_matrix(dist_matrix)

    best_vertex, best_sum, sums = find_best_warehouse_vertex(dist_matrix)

    print("\nСуммы расстояний для каждой вершины:")
    for i, total in enumerate(sums):
        print(f"Вершина {i + 1}: {total}")

    print(f"\nЛучшая вершина для склада: {best_vertex + 1}")
    print(f"Минимальная сумма расстояний: {best_sum}")


if __name__ == "__main__":
    main()

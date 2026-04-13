package main

import "fmt"

// bfsFromVertex считает кратчайшие расстояния
// от вершины start до всех остальных вершин.
func bfsFromVertex(matrix [][]int, start int) []int {
	n := len(matrix)

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for next := 0; next < n; next++ {
			if matrix[current][next] == 1 && dist[next] == -1 {
				dist[next] = dist[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return dist
}

// buildDistanceMatrix строит матрицу кратчайших расстояний
func buildDistanceMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	distMatrix := make([][]int, n)

	for i := 0; i < n; i++ {
		distMatrix[i] = bfsFromVertex(matrix, i)
	}

	return distMatrix
}

// findBestWarehouseVertex ищет вершину с минимальной суммой расстояний
func findBestWarehouseVertex(distMatrix [][]int) (bestVertex int, bestSum int, sums []int) {
	n := len(distMatrix)
	sums = make([]int, n)

	bestVertex = -1
	bestSum = -1

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += distMatrix[i][j]
		}

		sums[i] = sum

		if bestVertex == -1 || sum < bestSum {
			bestVertex = i
			bestSum = sum
		}
	}

	return bestVertex, bestSum, sums
}

func printMatrix(matrix [][]int) {
	n := len(matrix)

	fmt.Print("     ")
	for j := 0; j < n; j++ {
		fmt.Printf("%3d", j+1)
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Printf("%3d |", i+1)
		for j := 0; j < n; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int{
		/*1*/ {0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		/*2*/ {1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		/*3*/ {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		/*4*/ {1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		/*5*/ {0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		/*6*/ {0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		/*7*/ {0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0},
		/*8*/ {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		/*9*/ {0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
		/*10*/ {0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1},
		/*11*/ {0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	}

	distMatrix := buildDistanceMatrix(matrix)

	fmt.Println("Матрица кратчайших расстояний:")
	printMatrix(distMatrix)

	bestVertex, bestSum, sums := findBestWarehouseVertex(distMatrix)

	fmt.Println("\nСуммы расстояний для каждой вершины:")
	for i, sum := range sums {
		fmt.Printf("Вершина %d: %d\n", i+1, sum)
	}

	fmt.Printf("\nЛучшая вершина для склада: %d\n", bestVertex+1)
	fmt.Printf("Минимальная сумма расстояний: %d\n", bestSum)
}

package heap

import (
	"Algorithms-Learning-With-Go/sort/utils"
	"fmt"

	"testing"
)

func TestHeapSort(t *testing.T) {
	//list := utils.GetArrayOfSize(10)
	//list := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	list := []int{8, 7, 5, 2, 3}

	fmt.Println(list)
	Sort(list)
	fmt.Println(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkHeapSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)

	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func BenchmarkHeapSort100(b *testing.B)    { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)   { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkHeapSort(10000, b) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkHeapSort(100000, b) }

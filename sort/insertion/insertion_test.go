package insertion

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	
	"testing"
)


func TestInsertionSort(t *testing.T) {
	list := utils.GetArrayOfSize(5)

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

func TestSort1(t *testing.T) {
	list := utils.GetArrayOfSize(5)
	fmt.Println(list)
	Sort1(list)
	fmt.Println(list)
}

func benchmarkInsertionSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func BenchmarkInsertionSort100(b *testing.B) { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)   { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B)  { benchmarkInsertionSort(10000, b) }
func BenchmarkInsertionSort100000(b *testing.B) { benchmarkInsertionSort(100000, b) }

package quick

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	
	"testing"
)


func TestQuickSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
	fmt.Println(list)
	QuickSort(list, 0, len(list)-1)
	//HoareSort(list, 0, len(list)-1)
	fmt.Println(list)
}

func TestSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
	fmt.Println(list)
	Sort(list)
	fmt.Println(list)
}

func TestBubbleSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
	fmt.Println(list)
	BubbleSort(list)
	fmt.Println(list)
}

func benchmarkQuickSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		HoareSort(list, 0, len(list)-1)
		//QuickSort(list, 0, len(list)-1)
	}
}

func BenchmarkSort(b *testing.B) {
	list := utils.GetArrayOfSize(100000)
	Sort(list)
}

func BenchmarkBubbleSort(b *testing.B) {
	list := utils.GetArrayOfSize(100000)
	BubbleSort(list)
}

func BenchmarkQuickSort100(b *testing.B) { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkQuickSort(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkQuickSort(100000, b) }

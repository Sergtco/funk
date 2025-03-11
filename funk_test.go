package funk

import (
	"maps"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

func TestFold(t *testing.T) {
	ints := slices.Values([]int{1, 2, 3})

	res := Fold(ints, 0, func(acc int, val int) int { return acc + val })

	if res != 6 {
		t.Errorf("Expected %d, got %d", 6, res)
	}
}

func TestFoldOnMap(t *testing.T) {
	ints := slices.Values([]int{1, 2, 3, 1, 2, 3})
	freqs := map[int]int{}

	res := Fold(ints, freqs, func(acc map[int]int, val int) map[int]int {
		acc[val]++
		return acc
	})

	expected := map[int]int{
		1: 2,
		2: 2,
		3: 2,
	}
	if !maps.Equal(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestReduce(t *testing.T) {
	ints := slices.Values([]int{1, 2, 3})

	res, err := Reduce(ints, func(acc int, val int) int { return acc + val })
	if err != nil {
		t.Errorf("Expected success, got: %v", err)
	}

	if res != 6 {
		t.Errorf("Expected %d, got %d", 6, res)
	}
}

func TestMap(t *testing.T) {
	ints := slices.Values([]int{1, 2, 3})

	strings := slices.Collect(Map(ints, strconv.Itoa))

	expected := []string{"1", "2", "3"}
	if !slices.Equal(strings, expected) {
		t.Errorf("Expected %v, got: %v", expected, strings)
	}
}

func BenchmarkFold(b *testing.B) {
	items := make([]int, 10)
	for i := range items {
		items[i] = rand.Int()
	}
	for range b.N {
		itemsIter := slices.Values(items)
		Fold(itemsIter, 0, func(acc int, val int) int { return acc + val })
	}
}

func BenchmarkReduce(b *testing.B) {
	items := make([]int, 10)
	for i := range items {
		items[i] = rand.Int()
	}
	for range b.N {
		itemsIter := slices.Values(items)
		_, _ = Reduce(itemsIter, func(acc int, val int) int { return acc + val })
	}
}

func BenchmarkFor(b *testing.B) {
	items := make([]int, 10)
	for i := range items {
		items[i] = rand.Int()
	}
	for range b.N {
		acc := 0
		for _, val := range items {
			acc += val
		}
	}
}

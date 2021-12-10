package repeated

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func TestToLower(t *testing.T) {
	lower := strings.ToLower("ABVBC")
	expected := "abvbc"
	if lower != expected {
		t.Errorf("expected '%q' but got '%q'", expected, lower)
	}
}
func TestSumAll(t *testing.T) {
	arr := SumAll([]int{1, 2, 3, 4}, []int{2, 3, 4, 5})
	expected := []int{10, 14}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("got %v want %v", arr, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	assert := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assert(t, got, want)
	})

	t.Run("safely sum emtpy slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		assert(t, got, want)
	})
}

func SumAllTails(arrs ...[]int) []int {
	arr := make([]int, 0)
	for _, numbers := range arrs {
		if len(numbers) == 0 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, Sum(numbers[1:]))
		}

	}
	return arr
}
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		arr := [5]int{1, 2, 3, 4, 5}
		sum := Sum(arr[:])
		expected := 15
		if sum != expected {
			t.Errorf("got %d want %d given, %v", sum, expected, arr)
		}

	})
	t.Run("collection of any size", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		sum := Sum(arr)
		expected := 15
		if sum != expected {
			t.Errorf("got %d want %d given, %v", sum, expected, arr)
		}
	})

}

func Sum(arr []int) int {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

func SumAll(c ...[]int) []int {
	newArray := make([]int, 0)
	for _, arr := range c {
		newArray = append(newArray, Sum(arr))
	}
	return newArray
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func Repeat(s string, count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += s
	}

	return repeat
}

func TestSlice(t *testing.T) {
	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	println(fmt.Println(a))

	var p []int
	println(p)
}

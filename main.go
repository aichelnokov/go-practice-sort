package main

import (
	"fmt"
	"math/rand"

	// "time"
	"sort"
)

func main() {
	var a []int
	var b []int

	a = getALongSortedArray(100)
	b = getALongSortedArray(100)

	fmt.Println("Arrays")
	printArray(a)
	printArray(b)
	fmt.Println("===")

	// r1 := mergeOne(a, b)
	// printArray(r1)
	r2 := mergeTwo(a, b)
	printArray(r2)
}

/*
Алогритм самый медленный
Количество итераций: len(a) + len(b) - количество элементов в оставшемся хвосте
*/
func mergeOne(a []int, b []int) []int {
	var result []int
	ia := 0
	ib := 0
	ir := 0

	max := len(a) + len(b) - 1

	for ir < max {
		if a[ia] < b[ib] {
			result = append(result, a[ia])
			ia++
			ir++
			if ia == len(a) {
				result = append(result, b[ib:]...)
				break
			}
			continue
		}
		if a[ia] >= b[ib] {
			result = append(result, b[ib])
			ib++
			ir++
			if ib == len(b) {
				result = append(result, a[ia:]...)
				break
			}
			continue
		}

	}

	return result
}

/*
Второй алгоритм, который находит максимально большие куски для вставки в итоговый массив из двух других
*/
func mergeTwo(a []int, b []int) []int {
	var result []int
	var ia int = 0 // позиция, по которую надо отрезать записи из массива a
	var ib int = 0 // позиция, по которую надо отрезать записи из массива b
	var ka int = 0 // позиция, откуда начинать отрезать в следующий раз из массива a
	var kb int = 0 // позиция, откуда начинать отрезать в следующий раз из массива b

	for (ka >= 0) || (kb >= 0) {
		mergeTwoFindLess(a, ia, b, &ib)
		if ib >= 0 && kb < ib {
			result = append(result, b[kb:ib]...)
			kb = ib
		} else {
			if kb >= 0 {
				result = append(result, b[kb:]...)
				kb = -1
			}
		}
		mergeTwoFindLess(b, ib, a, &ia)
		if ia >= 0 && ka < ia {
			result = append(result, a[ka:ia]...)
			ka = ia
		} else {
			if ka >= 0 {
				result = append(result, a[ka:]...)
				ka = -1
			}
		}
	}

	return result
}

func mergeTwoFindLess(la []int, lk int, ar []int, s *int) bool {
	if lk < 0 || lk >= len(la) || *s < 0 || *s >= len(ar) {
		return false
	}
	for ar[*s] <= la[lk] {
		*s++
		if len(ar) == *s {
			*s = -1
			break
		}
	}
	return true
}

func printArray(a []int) {
	var e string

	for i := range a {
		e = e + fmt.Sprintf("%[1]d, ", a[i])
	}

	fmt.Println("[" + e + "]")
}

func getALongSortedArray(l int) []int {
	var a []int
	var i int
	if l == 0 {
		l = 100
	}
	i = 0

	for i < l {
		a = append(a, rand.Int()/10000000000000)
		i++
	}

	sort.Ints(a)

	return a
}

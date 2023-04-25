package main 

import (
	"fmt"
)

// Бинарный поиск можно делать только на отсортированном массиве
// Мы пользуемся тем фактом, уверены в сравнении

func binarySearch(arr []int, needed int) (int, error) {

	i, j := 0, len(arr)-1

	for i <= j {
		mid := (i + j) / 2 // Идея простая. Если меньше или больше - отсекаем половину нашего массива. И искать теперь нужно по меньшему количеству элементов.
		if arr[mid] == needed {
			return mid, nil
		}

		if needed < arr[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return 0, fmt.Errorf("Не нашли")
}


func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	ind, _ := binarySearch(arr, 6)
	fmt.Println(ind)
}
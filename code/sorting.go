package main

import (
	"fmt"
	"sort"
)

func main() {
	abc := []string{"jkl", "ghi", "abc", "def"}
	sort.Strings(abc)
	fmt.Println("Sorted ABC:", abc)

	nums := []int{4, 2, 12, 5, 1, 3}
	sort.Ints(nums)
	fmt.Println("Sorted Nums:", nums)

	// 这里比较神奇，sort.Interface类似C++中的 `<<` 可以叠加赋值
	// Reverse sort, need to cast abc as a StringSlice
	// to reverse and sort
	sort.Sort(sort.Reverse(sort.StringSlice(abc)))
	fmt.Println("Reverse ABC:", abc)

	// To sort a map by keys, pull an array of keys out of the map,
	// sort the keys and then iterate over the sorted keys.
	// A map by itself is not sortable.
	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	// Create array of kys
	var keys []string
	for k := range hash {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Use ordered keys to loop through hash
	for i := range keys {
		fmt.Printf("%s => %v\n", keys[i], hash[keys[i]])
	}
}

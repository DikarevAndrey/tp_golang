package main

import (
	"sort"
	"strconv"
	"strings"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	res := [3]int{1, 3, 4}
	return res
}

func ReturnIntSlice() []int {
	res := []int{1, 2, 3}
	return res
}

func IntSliceToString(slice []int) string {
	strSlice := []string{}
	for i := range slice {
		str := strconv.Itoa(slice[i])
		strSlice = append(strSlice, str)
	}
	return strings.Join(strSlice, "")
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	res := []int{}
	for i := range floatSlice {
		res = append(res, int(floatSlice[i]))
	}
	for i := range intSlice {
		res = append(res, int(intSlice[i]))
	}
	return res
}

func GetMapValuesSortedByKey(inputMap map[int]string) []string {
	mapKeys := []int{}
	for k := range inputMap {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys)

	var sortedValues []string
	for _, v := range mapKeys {
		sortedValues = append(sortedValues, inputMap[v])
	}
	return sortedValues
}

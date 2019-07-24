package golib

import (
	"math/rand"
	"reflect"
	"time"
)

//  ----------------- Funcs --------------------------

type reducetype func(interface{}) interface{}

// ReduceSlice generates a new slice after parsing every value by reduce function
// func ReduceSlice(slice []interface{}, Reduce func(interface{}) interface{}) (dslice []interface{}) {
// 	for _, v := range slice {
// 		dslice = append(dslice, Reduce(v))
// 	}
// 	return
// }

// ReduceSlice generates a new slice after parsing every value by reduce function
func ApplySlice(slice []interface{}, apply func(interface{}) interface{}) (dslice []interface{}) {
	for _, v := range slice {
		dslice = append(dslice, apply(v))
	}
	return
}

type filtertype func(interface{}) bool

// SliceFilter generates a new slice after filter function.
func FilterSlice(slice []interface{}, Filter func(interface{}) bool) (ftslice []interface{}) {
	for _, v := range slice {
		if Filter(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

//  ----------------- Bool --------------------------
func EqualSlice(s1, s2 []interface{}) bool {
	return reflect.DeepEqual(s1, s2)
}

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceMerge merges interface slices to one slice.
func MergeSlice(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

// ------------------- Set --------------------------
// github.com/yudeguang/slice
// DistinctSlice to remove duplicate value
func DistinctSliceString(slice []string) []string {
	distinct := []string{}
	sliceMap := map[string]bool{}
	for _, v := range slice {
		if !sliceMap[v] {
			sliceMap[v] = true
			distinct = append(distinct, v)
		}
	}
	return distinct
}

// UniqueSlice cleans repeated values in slice.
// UniqueSlice slice 去重
func UniqueSlice(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSliceIface(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// DiffSlice returns diff slice of slice1 - slice2.
// DiffSlice 子集
func DiffSlice(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// IntersectSlice returns slice that are present in all the slice1 and slice2.
// IntersectSlice 交集 Returns elements in slice1 that are also in slice2
func IntersectSlice(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// ChunkSlice separates one slice to some sized slice.
// ChunkSlice 按指定的 size 分割 slice 块, 返回一个二维的slice
func ChunkSlice(slice []interface{}, size int) (chunkslice [][]interface{}) {
	if size >= len(slice) {
		chunkslice = append(chunkslice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkslice = append(chunkslice, slice[i:end])
		end += size
	}
	return
}

// PadSlice prepends size number of val into slice.
// PadSlice 用val填充slice到size指定的大小, 不足用val填充,超过返回原slice
func PadSlice(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

// SliceShuffle shuffles a slice.
// ShuffleSlice 打乱slice
func ShuffleSlice(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

// RandSliceList generate an int slice from min to max.
func RandSliceList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

// RandSlice returns a random element from slice.
func RandSlice(slice []interface{}) (b interface{}) {
	randnum := rand.Intn(len(slice))
	b = slice[randnum]
	return
}

// RangeSlice generates a new slice from begin to end with step duration of int64 number.
func RangeSliceInt64(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

// SliceSum sums all values in int64 slice.
func SumSliceInt64(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

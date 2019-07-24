package golib

import "strings"

type Array struct{}

func (this Array) IsContainString(list *[]string, searchItem string) bool {
	ret := false
	for _, item := range *list {
		if item == searchItem {
			ret = true
			break
		}
	}
	return ret
}

func (this Array) UniqueString(slice *[]string) *[]string {
	if slice == nil {
		return slice
	}

	tempMap := make(map[string]bool)
	newSlice := make([]string, 0)
	for _, ele := range *slice {
		if _, ok := tempMap[ele]; !ok {
			newSlice = append(newSlice, ele)
			tempMap[ele] = true
		}
	}
	return &newSlice
}

// ArrayRemoveDuplicates removes any duplicated elements in a string array.
func ArrayRemoveDuplicates(s *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *s {
		if !found[x] {
			found[x] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

// 求交集 [1, 2, 2, 3, 4] ∩ [2,4,5, 5] = [2,  4]
func (this Array) IntersectionString(tmpFirstList *[]string, tmpSecondList *[]string) *[]string {
	firstList := this.UniqueString(tmpFirstList)
	secondList := this.UniqueString(tmpSecondList)

	ret := make([]string, 0)
	for _, first := range *firstList {
		for _, second := range *secondList {
			if first == second {
				ret = append(ret, first)
				break
			}
		}
	}
	return &ret
}

func (this Array) UniqueInt(slice *[]int) *[]int {
	if slice == nil {
		return slice
	}

	tempMap := make(map[int]bool)
	newSlice := make([]int, 0)
	for _, ele := range *slice {
		if _, ok := tempMap[ele]; !ok {
			newSlice = append(newSlice, ele)
			tempMap[ele] = true
		}
	}
	return &newSlice
}

/* 求字符串差集, 逻辑例如：
 * 	[1,2,3] - [2,4] = [1,3]
 *	[2,4] - [1,2,3] -  = [4]
 */
func (this Array) DiffSetString(tmpFirstList *[]string, tmpSecondList *[]string) *[]string {
	// 先过滤
	firstList := this.UniqueString(tmpFirstList)
	secondList := this.UniqueString(tmpSecondList)

	diffList := make([]string, 0)
	for _, first := range *firstList {
		isInSecond := false
		for _, second := range *secondList {
			if first == second {
				isInSecond = true
				break
			}
		}
		if !isInSecond {
			diffList = append(diffList, first)
		}
	}
	return &diffList
}

func (this Array) HasSuffixInArray(slice *[]string, str string) bool {
	if slice == nil {
		return false
	}

	for _, ele := range *slice {
		if strings.HasSuffix(str, ele) {
			return true
		}
	}
	return false
}

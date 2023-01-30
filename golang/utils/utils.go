package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func MergeMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Int64SliceToString(slices []int64) string {

	str := fmt.Sprintf("%v", slices)
	uStr := strings.Replace(strings.Trim(fmt.Sprint(str), "[]"), " ", ",", -1)

	return uStr
}

func SliceToMap(slices []int64) map[int64]struct{} {
	m := make(map[int64]struct{}, 0)
	if slices == nil || len(slices) == 0 {
		return m
	} else {
		for _, e := range slices {
			m[e] = struct{}{}
		}
		return m
	}
}

func MapKeyToInt64Arr(m map[int64]struct{}) []int64 {
	arr := make([]int64, 0)
	if m != nil && len(m) > 0 {
		for k, _ := range m {
			arr = append(arr, k)
		}
	}
	return arr
}

type SortElement struct {
	Name  string
	Total int64
}

type SortList []SortElement

//Len()
func (s SortList) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s SortList) Less(i, j int) bool {
	return s[i].Total > s[j].Total
}

//Swap()
func (s SortList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func GetSortedList(mSize map[string]int64) *SortList {
	arr := make(SortList, 0)
	for k, v := range mSize {
		arr = append(arr, SortElement{Name: k, Total: v})
	}
	sort.Sort(arr)
	return &arr
}

type Uint64SortElement struct {
	Name  string
	Total uint64
}

type Uint64SortList []Uint64SortElement

//Len()
func (s Uint64SortList) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s Uint64SortList) Less(i, j int) bool {
	return s[i].Total > s[j].Total
}

//Swap()
func (s Uint64SortList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func GetUint64SortedList(mSize map[string]uint64) *Uint64SortList {
	arr := make(Uint64SortList, 0)
	for k, v := range mSize {
		arr = append(arr, Uint64SortElement{Name: k, Total: v})
	}
	sort.Sort(arr)
	return &arr
}

func GetApproximation30Second(t string) (string, error) {
	startTimeSecond, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return t, err
	}
	second := startTimeSecond.Second()
	if second > 30 {
		startTimeSecond = startTimeSecond.Add(-1 * time.Duration((second-30)*1000000000))
	} else if second < 30 && second > 0 {
		startTimeSecond = startTimeSecond.Add(-1 * time.Duration(second*1000000000))
	}
	return startTimeSecond.Format("2006-01-02 15:04:05"), nil
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    ret := make([]int,0)
    //var ret []int
    for index, val := range nums {
	m_val, found := m[target - val]
	if found {
	    ret = append(ret, m_val, index)
	    break
	}
	m[val] = index
    }
    return ret
}

func main() {
    var  nums = []int{2,7,9,12}
    fmt.Printf("program result: %d\n", twoSum(nums,9))
    fmt.Println("expected result: [0,1]")
}

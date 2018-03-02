package test_orderinteger

import (
	"strconv"
)

func orderintegers(nums []int) string {
    var str string
    num := 0
    slice := make([]int,10)
	for x := 0; x < len(nums); x++ {
        str = strconv.Itoa(nums[x])
		for y := 0; y < len(nums); y++ {
			if nums[x] == nums[y] {
				str = str + ""
			} else {
				str = str + strconv.Itoa(nums[y])
			}
        }
        slice[x], _ = strconv.Atoi(str)
        str = strconv.Itoa(nums[x])
        for z := (len(nums)-1); z >= 0; z-- {
			if nums[x] == nums[z] {
				str = str + ""
			} else {
				str = str + strconv.Itoa(nums[z])
			}
        }
        slice[x+3], _ = strconv.Atoi(str)
	}
    
	for r := 0; r < len(slice); r++ {
        if slice[r] > num {
            num = slice[r]
        }
    }
    str = strconv.Itoa(num)

	return str
}

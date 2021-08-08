package piscine

// func TwoSum(nums []int, target int) []int {
// nums := 0
// fmt.Scan(&nums)
// for i, j := range nums {
// 	for k := i + 1; k < len(nums); k++ {
// 		if nums[k]+j == target {
// 			return []int{i, k}
// 		}
// 	}
// }
// return []int{}
// }

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

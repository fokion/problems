package leetcode

func TwoSum(nums []int, target int) []int {

	var indexes [2]int

	for i := 0; i < len(nums); i++ {
		number := nums[i]

		indexes[0] = i

		for j := 0; j < len(nums); j++ {
			if i == j {

				continue
			}

			if nums[j]+number == target {
				indexes[1] = j

				return indexes[:]
			}
		}

	}
	return nil

}

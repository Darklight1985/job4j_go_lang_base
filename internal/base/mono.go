package base

func Mono(nums []int) bool {
	up := true
	down := true
	size := len(nums)
	for i := 1; i < size; i++ {
		down = down && nums[i] <= nums[i-1]
		up = up && nums[i] >= nums[i-1]
	}
	return up || down
}

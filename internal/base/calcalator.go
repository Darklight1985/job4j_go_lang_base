package base

func Add(first, second int) int {
	return first + second
}

func Mono(nums []int) bool {
	var up bool
	size := len(nums)
	for i := 1; i < size; i++ {
		if nums[i] < nums[i-1] {
			if i > 1 && up {
				return false
			}
			up = false
		}
		if nums[i] > nums[i-1] {
			if i > 1 && !up {
				return false
			}
			up = true
		}
	}
	return true
}

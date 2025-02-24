Problem: Given an array of integers nums and an integer target, return all the indices of the two numbers that add up to target.
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: nums[0] + nums[1] = 2 + 7 = 9

Key concepts Learned:

To compare an element with all other elemenets of an array use loop inside a loop.

In the following code nums[i] will be compared with the next element nums[j] where j is [i+1]

for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j}) // Add pair as a slice
			}
		}
	}
	return result



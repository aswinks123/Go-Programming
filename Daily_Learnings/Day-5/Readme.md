# Problem: Write a Go program to find the second largest element in an array without using built-in functions.

```markdown

Input: nums = [10, 20, 4, 45, 99]
Output: 45

```

# Key concepts Learned:
 ```markdown

First find the largest in the array then find the second largest


```
```go
// Find the largest

	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}

	}
	// Now Find the second largest

	secondLargest := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > secondLargest && nums[i] < largest {   // If [i]th element is larget than second largest(firstly assignes as -1) and less than the largest that means the [i]th element is the second largest. Therefore second Largest =nums[i]
			secondLargest = nums[i]
		}
	}
	fmt.Println("Second largest is :", secondLargest)
```

# Problem: Check if an Array is Sorted (Ascending Order)

```markdown

Input: nums := []int{1, 2, 3, 4, 5}
Output: The array is sorted

```

# Key concepts Learned:

```go
	nums := []int{1, 2, 3, 4, 5}
	isSorted := true // first set this boolean variable as true

	for i := 1; i < len(nums); i++ {

		if nums[i] < nums[i-1] { //this condition makes sure there is no index out of range error. Don't use nums[i] > nums[i+1]
			// because at the end nums[i+1] will cause array out of index
			isSorted = false //if the array is not sorted,  then set isSorted as false
			break
		}

	}
```

# Problem: Given an integer slice nums, find the maximum number in the slice.
# You must implement a function that returns the maximum number.

Input: nums = [3, 1, 35, 7, 9, 5]
Output: 35

# Key concepts Learned:
A function can return multiple values. Here it returns an int and an error message

```go
func largenumber(nums []int) (int, error)
```

We can check whether a slice is empty using the len function
```go
if len(nums) == 0 {
		return 0, fmt.Errorf("Empty slice. Please add value to slice first")
	}
```

We can check whether there is an error using the following

```go
if err != nil {                    
		fmt.Print(err)
		return // Exit the program
	}
```
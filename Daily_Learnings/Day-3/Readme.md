# Problem: Write a function in Go to check whether a given number is prime or not.
Input: n = 7
Output: true
Explanation:7 is only divisible by 1 and 7
Input: n = 10
Output: false
Explanation: 10 is not a prime number because it is divisible by 1, 2, 5, 10.
Input: nums = [3, 1, 35, 7, 9, 5]
Output: 35

# Key concepts Learned:
A number is prime if:
It is greater than 1.
It is only divisible by 1 and itself (no other divisors).
When you check the remainder (%) for every number from 2 to num-1, none should be 0.

# Example with 5:
Check all numbers from 2 to 4:
5 % 2 = 1 ✅
5 % 3 = 2 ✅
5 % 4 = 1 ✅
Since none of the remainders are 0, 5 is only divisible by 1 and 5.
Therefore, 5 is prime.

```go
if num == 1 { // if number is 1 the return false. 1 is not a prime number by default
		return false
	}
```
```go
	for i := 2; i < num; i++ { //check the number from 2 to n-1 number
		if num%i == 0 { //if reminder of any number between 2 and n-1 is 0 , then that is not a prime number
			return false
		}
	}
```



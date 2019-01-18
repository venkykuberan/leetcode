package palindrome

func isPalidrome(x int) bool {
	rev := 0
	val := x

	if val < 0 {
		return false
	}
	for val != 0 {
		rev = rev*10 + val%10
		val = val / 10
	}
	if x == rev {
		return true
	}
	return false
}

//Iterate the loop half distance

func isPalindrome1(val int) bool {
	rev := 0

	// Special cases:
	// As discussed above, when x < 0, x is not a palindrome.
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.

	if val < 0 || (val%10 == 0 && val != 0) {
		return false
	}

	for val > rev {

		rev = rev*10 + val%10
		val = val / 10
	}

	// When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	// For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123,
	// since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.

	if val == rev || val == rev/10 {
		return true
	}

	return false

}

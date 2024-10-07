package palindromeNumber

func isPalindrome(x int) bool {
	mem := x
	var reverse int
	for mem > 0 {
		reverse = reverse*10 + (mem % 10)
		mem = mem / 10
	}

	return reverse == x
}

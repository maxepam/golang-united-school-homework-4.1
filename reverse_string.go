package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input) 
	var buf []rune
	for i := len(runes) - 1; i >=0; i-- {
		buf = append(buf, runes[i]) 
	}
	output = string(buf)
	return output
}
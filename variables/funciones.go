package variables

import "strconv"

func ConvertToText(n int) (bool, string) {

	//text := string(n)
	text := strconv.Itoa(n)

	//text, er:= strconv.Atoi(n)
	return true, text
}

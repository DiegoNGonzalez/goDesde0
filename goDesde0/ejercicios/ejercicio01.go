package ejercicios

import "strconv"

func Exercise01(textToConvert string) (int, string) {
	number, err := strconv.Atoi(textToConvert)
	if err != nil {
		return 0, "Error during the conversion"+ err.Error()
	}
	if number > 100 {
		return number, "The number is greater than 100"
	} else {
		return number, "The number is less than 100"
	}

}
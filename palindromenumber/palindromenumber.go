package palindromenumber

import (
	// "bufio"
	// "fmt"
	"math"
	// "os"
	"strconv"
	"strings"
)

type InvalidArgsError struct {
	msg string
}

func (e *InvalidArgsError) Error() string { return e.msg }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		count, err := CountPalindromes(scanner.Text())

// 		if err == nil {
// 			fmt.Println(count)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading standard input:", err)
// 	}
// }

func CountPalindromes(numInput string) (int, error) {
	n := strings.Split(strings.TrimSpace(numInput), " ")
	count := 0

	if len(n) != 2 {
		return 0, &InvalidArgsError{"input must contain two numbers separated with one space."}
	}

	numStart, errNumStart := strconv.Atoi(n[0])
	numFinish, errNumFinish := strconv.Atoi(n[1])

	if errNumStart != nil || errNumFinish != nil {
		return 0, &InvalidArgsError{"input must contain two numbers separated with one space."}
	}

	if numStart < 1 || numFinish > int(math.Pow(10, 9)) {
		return 0, &InvalidArgsError{"input numbers must be between 1 and 1.000.000.000"}
	}

	if numFinish <= numStart {
		return 0, &InvalidArgsError{"the second number must be greater than the first number."}
	}

	for i := numStart; i <= numFinish; i++ {
		if check(strconv.Itoa(i)) {
			count++
		}
	}

	return count, nil
}

func check(str string) bool {
	s := strings.ToLower(strings.TrimSpace(str))
	return s == reverseString(s)
}

func reverseString(str string) string {
	var reversedRunes []rune
	strRunes := []rune(str)

	for i := len(str) - 1; i >= 0; i-- {
		reversedRunes = append(reversedRunes, strRunes[i])
	}

	return string(reversedRunes)
}

package missingnumber

import (
	// "bufio"
	// "fmt"
	"math"
	// "os"
	"strconv"
)

type InvalidArgsError struct {
	msg string
}

func (e *InvalidArgsError) Error() string { return e.msg }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		res, err := FindMissingNumber(scanner.Text())

// 		if err == nil {
// 			fmt.Println(res)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}
// }

func getNumber(input string, offset int, limit int) int {
	if offset + limit > len(input) {
		return -2 // invalid parameter
	}

	number, err := strconv.Atoi(input[offset:offset + limit])

	if err != nil {
		return -1 // invalid number
	}

	return number
}

func validSequenceNumber(numbers string) ([]int, bool) {
	isValid := false
	sequenceNumber := make([]int, 0)
	inputLength := len(numbers)

	if inputLength >= 3 && inputLength <= 1000 {
		digitLength := 1
		index := 0

		BeginLoop:
		// fmt.Println("LOOP STARTED")

		sequenceNumber = make([]int, 0)

		for i := 0; i < inputLength; i++ {
			currNumber := getNumber(numbers, index, digitLength)
			nextNumber := getNumber(numbers, index + digitLength, digitLength)

			// fmt.Println("currNumber", currNumber, "nextNumber", nextNumber)

			sequenceNumber = append(sequenceNumber, currNumber)

			if currNumber > int(math.Pow(10, 6)) {
				sequenceNumber = make([]int, 0)

				goto EndLoop
			}

			// if nextNumber >= 0 && index + digitLength * 2 == inputLength {
			// 	sequenceNumber = append(sequenceNumber, nextNumber)
			// 	isValid = true

			// 	goto EndLoop
			// }

			if nextNumber == -2 && index + digitLength == inputLength {
				isValid = true

				goto EndLoop
			} else if nextNumber == -1 {
				sequenceNumber = make([]int, 0)

				goto EndLoop
			}

			// fmt.Println(">> NUMBER TEST STARTED")

			if nextNumber >= 0 && nextNumber - currNumber <= 2 && nextNumber - currNumber > 0 {
				// fmt.Println(">> NUMBER TEST PASSED")

				index = index + digitLength
			} else {
				// fmt.Println("<< NUMBER TEST FAILED")

				anotherNextNumber := getNumber(numbers, index + digitLength, digitLength + 1)

				// fmt.Println("currNumber", currNumber, "anotherNextNumber", anotherNextNumber)

				// fmt.Println(">> ANOTHER NUMBER TEST STARTED")

				if anotherNextNumber - currNumber <= 2 && anotherNextNumber - currNumber > 0 {
					// fmt.Println(">> ANOTHER NUMBER TEST PASSED")

					index = index + digitLength

					digitLength++
				} else {
					// fmt.Println("<< ANOTHER NUMBER TEST FAILED")

					digitLength++
					index = 0

					// fmt.Println("<< BACK TO THE BEGINNING OF LOOP")
					goto BeginLoop
				}
			}
		}

		EndLoop:
		// fmt.Println("LOOP ENDED")
	}

	return sequenceNumber, isValid
}

func FindMissingNumber(numbers string) ([]int, error) {
	missingNumbers := make([]int, 0)
	sequenceNumber, valid := validSequenceNumber(numbers)

	if !valid {
		return missingNumbers, &InvalidArgsError{"invalid parameter"}
	}

	for n, _ := range sequenceNumber {
		if n > 0 {
			prevNumber := sequenceNumber[n - 1]
			currNumber := sequenceNumber[n]
			if currNumber - prevNumber == 2 {
				missingNumbers = append(missingNumbers, currNumber - 1)
			}
		}
	}

	if len(missingNumbers) < 1 {
		return missingNumbers, &InvalidArgsError{"invalid parameter"}
	}

	return missingNumbers, nil
}

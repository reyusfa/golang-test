package organizebook

import (
	// "bufio"
	// "fmt"
	// "os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Book struct {
	CategoryName string
	Code string
	Height int
	Name string
}

type Books []Book

type InvalidBooksError struct {
	msg string
}

func (e *InvalidBooksError) Error() string { return e.msg }

var validBook = regexp.MustCompile(`^[0-9][A-z](1[3-9]|2[0-4])+(?:\s[0-9][A-z](1[3-9]|2[0-4]))*$`)

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		res, err := OrganizeBooks(scanner.Text())

// 		if err == nil {
// 			fmt.Println(res)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}
// }

func OrganizeBooks(books string) (string, error) {
	if !isDataValid(books) {
		return "", &InvalidBooksError{"invalid input data."}
	}

	book := mapBooks(books, " ")

	sort.Slice(book, func (i, j int) bool {
		if book[i].CategoryName == book[j].CategoryName {
			return book[i].Height > book[j].Height
		}
		return book[i].CategoryName < book[j].CategoryName
	})

	res := stashDuplicateBook(stringifyBooks(book))

	return res, nil
}

func stashDuplicateBook(books string) string {
	book := StringToSlice(books, " ")

	dup := make(map[string]int)
	var s []string

	for _, v := range book {
		if dup[v[:2]] < 2 {
			s = append(s, v)
		}
		dup[v[:2]]++
	}

	return strings.Join(s, " ")
}

func stringifyBooks(books Books) string {
	var s []string

	for _, v := range books {
		s = append(s, v.Code)
	}

	return strings.Join(s, " ")
}

func mapBooks(str string, separator string) Books {
	var data Books

	s := StringToSlice(str, separator)

	for _, v := range s {
		cid, _ := strconv.Atoi(v[:1])
		bheight, _ := strconv.Atoi(v[2:])

		data = append(data, Book{
			CategoryName: category(cid),
			Code: v,
			Height: bheight,
			Name: v[1:2],
		})
	}

	return data
}

func isDataValid(data string) bool {
	return validBook.MatchString(data)
}

func StringToSlice(str string, separator string) []string {
	return strings.Split(strings.TrimSpace(str), separator)
}

func category(code int) string {
	categories := map[int]string{
		6: "Applied Sciences (600)",
		7: "Arts (700)",
		0: "General (000)",
		9: "Geography, History (900)",
		4: "Language (400)",
		8: "Literature (800)",
		1: "Philosophy, Psychology (100)",
		2: "Religion (200)",
		5: "Science, Math (500)",
		3: "Social Sciences (300)",
	}

	return categories[code]
}

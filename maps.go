package main

import (
	"errors"
	"fmt"
)

// creating a function that takes a slice of user names and phone numbers and
// returns a map of all the values

type user struct {
	name        string
	phoneNumber int
}

func createUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

func main() {

	ages := map[string]int{
		"John": 30,
		"Jane": 25,
	}

	fmt.Println(ages["John"])

	fmt.Println(len(ages))

	userNames := []string{"John", "Aaron"}
	userNumbers := []int{1234, 4321}

	test1, err := createUserMap(userNames, userNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test1)

	userNames = append(userNames, "Mike")

	test2, err := createUserMap(userNames, userNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test2)
}

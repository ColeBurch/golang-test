package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	customerCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	spouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return customerCost + spouseCost, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func main() {
	text1 := "hello"
	text2 := "hello I am longer"
	text3 := "this message is very long, far too long for this function"

	fmt.Println(sendSMSToCouple(text1, text2))
	fmt.Println(sendSMSToCouple(text1, text3))
}

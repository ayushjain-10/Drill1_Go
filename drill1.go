package main

import (
	"fmt"
	"time"
)

func avg_weight() {
	weights := make([]float64, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter the weight of person %d (in kg): ", i+1)
		_, err := fmt.Scan(&weights[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += weights[i]
	}

	average := total / 5

	fmt.Printf("The average weight of 5 people is: %.2f kg\n", average)
}

func dob(){

	var age int
	var dob string

	fmt.Println("Please enter your date of birth (dd/mm): ")
	_, err := fmt.Scan(&dob)
	if err != nil {
		fmt.Println("please make sure to enter a valid date of birth")
		return
	}

	fmt.Println("Please enter your age: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println("please make sure to enter a valid age")
		return
	}

	dobWithYear := dob + "/2003"

	t, err := time.Parse("02/01/2006", dobWithYear)
	if err != nil {
		fmt.Println("please make sure to enter it as dd/mm")
		return
	}

	currentTime := time.Now()

	birthYear := currentTime.Year() - age

	if currentTime.YearDay() < t.YearDay() {
		birthYear--
	}

	fmt.Printf("Your birth year is: %d\n", birthYear)

}

func main() {
	dob()
	avg_weight()
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
    // Take input from input1-1.txt
    input, err := ioutil.ReadFile("input01.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert input to string
    inputString := string(input)

    // split string by new line and put into array
    inputArray := strings.Split(inputString, "\n")

    // Create array to hold numbers
    var numbers []int

    // Loop through inputArray
    for i := 0; i < len(inputArray); i++ {
        log.Printf("checking %s", inputArray[i])

        // Get first digit
        firstDigit, err := getFirstDigit(inputArray[i])
        if err != nil {
            log.Print(err)
            continue
        }

        // get Last digit
        lastDigit, err := getLastDigit(inputArray[i])
        if err != nil {
            log.Print(err)
            continue
        }

        // concat both numbers together
        number := concatNumbers(firstDigit, lastDigit)

        // push to array
        numbers = append(numbers, number)
    }

    // add all array numbers together 
    sum := addNumbers(numbers)

    // print it out
    fmt.Println(sum)
}

func getFirstDigit(input string) (int, error) {
    // loop through input characters until you find a numbers
    for i := 0; i < len(input); i++ {
        // convert character to string
        character := rune(input[i])

        // convert string to int
        number, err := convertCharToInt(character)
        if err != nil {
            log.Print(err)
            continue
        }

        return number, nil
    }

    return 0, fmt.Errorf("No number found")
}

func getLastDigit(input string) (int, error) {
    for i := len(input) - 1; i >= 0; i-- {
        // convert character to string
        character := rune(input[i])

        // convert string to int
        number, err := convertCharToInt(character)
        if err != nil {
            log.Print(err)
            continue
        }

        return number, nil
    }

    return 0, fmt.Errorf("No number found")
}

func convertCharToInt(char rune) (int, error) {
    if char < '0' || char > '9' {
        return 0, fmt.Errorf("'%c' is not a number", char)
    }

    return int(char - '0'), nil
}

func concatNumbers(first int, last int) int {
    return first * 10 + last
}

func addNumbers(numbers []int) int {
    sum := 0
    for i := 0; i < len(numbers); i++ {
        sum += numbers[i]
    }

    return sum
}


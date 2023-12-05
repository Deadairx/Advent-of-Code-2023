package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
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
        log.Print(inputArray[i])
        if inputArray[i] == "" { break }
        nums := convertStringToIntArray(inputArray[i])
        log.Print(nums)
        concat := concatNumbers(nums[0], nums[len(nums) - 1])
        log.Print(concat)
        numbers = append(numbers, concat)
    }

    // add all array numbers together 
    sum := addNumbers(numbers)

    // print it out
    fmt.Println(sum)
}

func convertStringToIntArray(input string) []int  {
    numberMap := map[string]int{
        "zero": 0,
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    numberArray := []int{}

    for i := 0; i < len(input); i++ { 
        found := false
        for word, number := range numberMap {
            if strings.HasPrefix(input[i:], word) {
                numberArray = append(numberArray, number)
                found = true
                break
            }
        }

        if !found && unicode.IsDigit(rune(input[i])) {
            value, _ := strconv.Atoi(string(input[i]))
            numberArray = append(numberArray, value)
        }
    }

    return numberArray
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


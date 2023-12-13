package main

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
map 
["1", 1]  4
["2", 2]  | 2
["3", 4]  | | 2
["4", 8]  | | | 1
["5", 14] |   | | 0
["6", 1] 
*/ 

func main() {
    inputBytes, _ := ioutil.ReadFile("input.txt")
    scratchCards := strings.Split(string(inputBytes), "\n")

    cardCount := make([]int, len(scratchCards))

    sum := 0

    for i := 0; i < len(scratchCards); i++ {
        if scratchCards[i] == "" { break }
        cardCount[i] += 1
        cardValues := strings.Split(scratchCards[i], ":")
        winningNumbers := strings.Split(strings.Split(cardValues[1], "|")[0], " ")
        myNumbers := strings.Split(strings.Split(cardValues[1], "|")[1], " ")

        matchingCount := 0

        for j := 0; j < len(myNumbers); j++ {
            if myNumbers[j] == "" { continue }
            for k := 0; k < len(winningNumbers); k++ {
                if myNumbers[j] == winningNumbers[k] {
                    matchingCount++
                    break
                }
            }
        }

        log.Printf("matching numbers: %d", matchingCount)

        if matchingCount > 0 {
            for x := 1; x <= matchingCount; x++ {
                cardCount[i+x] += cardCount[i]
            }
        }

        sum += cardCount[i]
    }

    print(sum)
}

package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
    inputBytes, _ := ioutil.ReadFile("input.txt")
    scratchCards := strings.Split(string(inputBytes), "\n")

    sum := 0

    for i := 0; i < len(scratchCards); i++ {
        if scratchCards[i] == "" { break }
        cardValues := strings.Split(scratchCards[i], ":")
        winningNumbers := strings.Split(strings.Split(cardValues[1], "|")[0], " ")
        myNumbers := strings.Split(strings.Split(cardValues[1], "|")[1], " ")

        matchingCount := 0

        for j := 0; j < len(myNumbers); j++ {
            if myNumbers[j] == "" { continue }
            for k := 0; k < len(winningNumbers); k++ {
                if myNumbers[j] == winningNumbers[k] {
                    log.Print(myNumbers[j])
                    matchingCount++
                    break
                }
            }
        }

        log.Printf("matching numbers: %d", matchingCount)

        if matchingCount > 0 {
            product := 1
            for x := 1; x < matchingCount; x++ {
                product += product
            }

            log.Printf("adding %d to sum", product)
            sum += product
        }
    }

    print(sum)
}

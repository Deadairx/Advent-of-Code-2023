package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
    // get input data, convert to string
    input, _ := ioutil.ReadFile("input.txt")
    inputString := string(input);
    inputMap := strings.Split(inputString, "\n")

    sum := 0

    // loop, scan for each number
    for i := 0; i < len(inputMap); i++ {
        if inputMap[i] == "" { break }

        for j := 0; j < len(inputMap[i]); j++ {
            x := rune(inputMap[i][j])

            if x == '*' {
                // if number has symbol around it, add to sum
                numNeigbor := hasNumberNeigbor(inputMap, i, j) 

                if len(numNeigbor) < 2 {
                    //log.Print(numNeigbor)
                    continue
                }
                    //log.Print(numNeigbor)

                one, _ := strconv.Atoi(numNeigbor[0])
                two, _ := strconv.Atoi(numNeigbor[1])
                product := one * two
                log.Printf("%d * %d = %d", one, two, product)

                sum += product
            }
        }
    }

    print(sum)
}

func hasNumberNeigbor(inputMap []string, y int, x int) []string {
    checkUp := y != 0
    checkLeft := x != 0
    checkDown := y+1 != len(inputMap)-1
    checkRight := x+1 != len(inputMap[y])

    numStr := []string{}

    if checkUp {
        for i := x; i < x+1; i++ {
            if unicode.IsDigit(rune(inputMap[y-1][i])) {
                numStr = append(numStr, GetNumberStr(inputMap[y-1], i))
                checkUp = false
            }
        }
    }
    if checkDown {
        for i := x; i < x+1; i++ {
            if unicode.IsDigit(rune(inputMap[y+1][i])) {
                numStr = append(numStr, GetNumberStr(inputMap[y+1], i))
                checkDown = false
            }
        }
    }
    if checkLeft {
        if unicode.IsDigit(rune(inputMap[y][x-1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y], x-1))
            checkLeft = false
        }
    }
    if checkRight {
        if unicode.IsDigit(rune(inputMap[y][x+1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y], x+1))
            checkRight = false
        }
    }
    if checkRight && checkUp {
        if unicode.IsDigit(rune(inputMap[y-1][x+1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y-1], x+1))
        }
    }
    if checkLeft && checkUp {
        if unicode.IsDigit(rune(inputMap[y-1][x-1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y-1], x-1))
        }
    }
    if checkLeft && checkDown {
        if unicode.IsDigit(rune(inputMap[y+1][x-1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y+1], x-1))
        }
    }
    if checkRight && checkDown {
        if unicode.IsDigit(rune(inputMap[y+1][x+1])) {
            numStr = append(numStr, GetNumberStr(inputMap[y+1], x+1))
        }
    }

    //log.Print(numStr)

    return numStr
}

func distinct(inputSlice []string) []string {
    var unique []string

    uniqueMap := map[string]bool{}

    for _, i := range inputSlice {
        if _, exist := uniqueMap[i]; !exist {
            uniqueMap[i] = true
            unique = append(unique, i)
        }
    }

    return unique
}

func hasSymbolNeigbor(inputMap []string, y int, x int, length int) bool {
    checkUp := y != 0
    checkLeft := x != 0
    checkDown := y+1 != len(inputMap)-1
    checkRight := x+length != len(inputMap[y])

    if checkUp {
        for i := x; i < x+length; i++ {
            if rune(inputMap[y-1][i]) != '.' {
                log.Print(string(rune(inputMap[y-1][i])))
                return true
            }
        }
    }
    if checkDown {
        for i := x; i < x+length; i++ {
            if rune(inputMap[y+1][i]) != '.' {
                log.Print(string(rune(inputMap[y+1][i])))
                return true
            }
        }
    }
    if checkLeft {
        if rune(inputMap[y][x-1]) != '.' {
            log.Print(string(rune(inputMap[y][x-1]) ))
            return true
        }
    }
    if checkRight {
        if rune(inputMap[y][x+length]) != '.' {

            log.Print("HI I'M THE ASSHOLE")
            log.Print(string(rune(inputMap[y][x+1]) ))
            return true
        }
    }
    if checkRight && checkUp {
        if rune(inputMap[y-1][x+length]) != '.' {
            log.Print(string(rune(inputMap[y-1][x+1]) ))
            return true
        }
    }
    if checkLeft && checkUp {
        if rune(inputMap[y-1][x-1]) != '.' {
            log.Print(string(rune(inputMap[y-1][x-1])))
            return true
        }
    }
    if checkLeft && checkDown {
        if rune(inputMap[y+1][x-1]) != '.' {
            log.Print(string(rune(inputMap[y+1][x-1])))
            return true
        }
    }
    if checkRight && checkDown {
        if rune(inputMap[y+1][x+length]) != '.' {
            log.Print(string(rune(inputMap[y+1][x+1])))
            return true
        }
    }

    return false
}

func GetNumberStr(line string, startPos int) string {
    output := ""
    trueStartPos := startPos
    for i := startPos; i >= 0; i-- {
        if !unicode.IsDigit(rune(line[i])) {
            trueStartPos = i+1
            break
        }
    }

    for i := trueStartPos; i != len(line) && unicode.IsDigit(rune(line[i])); i++ {
        x := rune(line[i])
        output += string(x)
    }
    
    return output
}

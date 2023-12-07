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

        for j := 0; j < len(inputMap[i]); {
            x := rune(inputMap[i][j])

            if unicode.IsDigit(x) {
                currentNumberStr := GetNumberStr(inputMap[i], j)
                // if number has symbol around it, add to sum
                has := hasSymbolNeigbor(inputMap, i, j, len(currentNumberStr)) 
                log.Printf("%s has neigbor %t", currentNumberStr, has)
                if has {
                    num, _ := strconv.Atoi(currentNumberStr)
                    sum += num
                }

                j += len(currentNumberStr)
            }

            j++
        }
    }

    log.Print(sum)
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

    for i := startPos; i != len(line) && unicode.IsDigit(rune(line[i])); i++ {
        x := rune(line[i])
        output += string(x)
    }
    
    return output
}

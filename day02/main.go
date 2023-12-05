package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cubes struct {
    Red   int
    Green int
    Blue  int
}

func main() {
    input, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    gameInput := strings.Split(string(input), "\n")

    sum := 0

    // loop thru each line
    for i := 0; i < len(gameInput); i++ {
        if gameInput[i] == "" { break }
        // parse values to cubes struct
        game := &Game{
            Number: GetGameNumber(gameInput[i]),
            Pulls: GetGamePulls(gameInput[i]),
        }
        // check if higher than MaxCubes
        if validGame(*game) {
            sum += game.Number
        }
    }

    println(sum)
}

func validGame(game Game) bool {
    MaxCubes := &Cubes{
        Red: 12,
        Green: 13,
        Blue: 14,
    }

    for i := 0; i < len(game.Pulls); i++ {
        pull := game.Pulls[i]

        if pull.Blue > MaxCubes.Blue || pull.Red > MaxCubes.Red || pull.Green > MaxCubes.Green {
            return false
        }
    }

    return true
}

func GetGameNumber(gameInput string) int {
    gameNumString := strings.Split(gameInput, ":")[0]
    gameNumString = strings.Split(gameNumString, " ")[1]

    num, _ := strconv.Atoi(gameNumString)
    return num
}

func GetGamePulls(gameInput string) []Cubes {
    pullsInput := strings.Split(gameInput, ":")[1]

    pulls := strings.Split(pullsInput, ";")

    output := []Cubes{}

    for i := 0; i < len(pulls); i++ {
        output = append(output, parsePull(pulls[i]))
    }
    
    return output
}

func parsePull(pull string) Cubes {
    x := strings.Split(pull, ",")

    cubes := &Cubes{}

    for i := 0; i < len(x); i++ {
        num, _ := strconv.Atoi(strings.Fields(x[i])[0])
        color := strings.Fields(x[i])[1]

        switch color {
        case "red":
            cubes.Red = num
        case "blue":
            cubes.Blue = num
        case "green":
            cubes.Green = num
        default:
            continue
        }
    }

    return *cubes

}

type Game struct {
    Number int
    Pulls []Cubes
}



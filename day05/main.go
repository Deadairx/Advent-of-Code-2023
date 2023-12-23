package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
    input, _ := ioutil.ReadFile("input.txt")
    inputString := string(input)
    sections := strings.Split(inputString, "\n\n")
    
    seeds := strings.Split(strings.TrimSpace(strings.Split(sections[0], ":")[1]), " ")
    seedToSoil := strings.Split(strings.Trim(sections[1], "\n"), "\n")
    soilToFert := strings.Split(strings.Trim(sections[2], "\n"), "\n")
    fertToWater := strings.Split(strings.Trim(sections[3], "\n"), "\n")
    waterToLight := strings.Split(strings.Trim(sections[4], "\n"), "\n")
    lightToTemp := strings.Split(strings.Trim(sections[5], "\n"), "\n")
    tempToHumid := strings.Split(strings.Trim(sections[6], "\n"), "\n")
    humidToLoc := strings.Split(strings.Trim(sections[7], "\n"), "\n")

    lowestLoc := math.MaxInt

    for i := range seeds {
        seed, _ := strconv.Atoi(seeds[i])

        soil := nowMap(seed, seedToSoil)
        fert := nowMap(soil, soilToFert)
        water := nowMap(fert, fertToWater)
        light := nowMap(water, waterToLight)
        temp := nowMap(light, lightToTemp)
        humid := nowMap(temp, tempToHumid)
        loc := nowMap(humid, humidToLoc)

        if loc < lowestLoc {
            lowestLoc = loc
        }
    }

    print(lowestLoc)
}

func nowMap(initValue int, mapKey []string) int {
    for j := 1; j < len(mapKey); j++ {
        valueMap := strings.Split(mapKey[j], " ")
        destination, _ := strconv.Atoi(valueMap[0])
        source, _ := strconv.Atoi(valueMap[1])
        length, _ := strconv.Atoi(valueMap[2])

        if sniffTest(initValue, source, length) {
            return initValue - source + destination
        }
    }

    return initValue
}

func sniffTest(origin int, source int, length int) bool {
    return origin >= source && origin < (source + length)
}

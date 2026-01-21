package question5kyu

import (
	"fmt"
	"strings"
)

//ref: https://www.codewars.com/kata/5902bc7aba39542b4a00003d

type Map map[string][]string

var hungerMap = Map{
	"antelope": []string{"grass"},
	"big-fish": []string{"little-fish"},
	"bug":      []string{"leaves"},
	"bear":     []string{"big-fish", "bug", "chicken", "cow", "leaves", "sheep"},
	"chicken":  []string{"bug"},
	"cow":      []string{"grass"},
	"fox":      []string{"chicken", "sheep"},
	"giraffe":  []string{"leaves"},
	"lion":     []string{"antelope", "cow"},
	"panda":    []string{"leaves"},
	"sheep":    []string{"grass"},
}

func WhoEatsWho(zoo string) []string {
	remainAnimals := strings.Split(zoo, ",")
	isEatable := true
	word := ""
	result := []string{}
	result = append(result, zoo)
	for {
		remainAnimals, isEatable, word = Fight(remainAnimals)
		if !isEatable {
			break
		}
		result = append(result, word)
	}

	winner := ""
	for _, remain := range remainAnimals {
		winner += remain + ","
	}

	result = append(result, strings.Trim(winner, ","))
	return result
}

func Fight(animals []string) ([]string, bool, string) {
	isEatable := false
	word := ""
	for i, v := range animals {
		eatable := hungerMap[v]
		if i != 0 {
			isEatable = isCanEat(eatable, animals[i-1])
			if isEatable {
				word = fmt.Sprintf("%s eats %s", v, animals[i-1])
				animals[i-1] = ""
				break
			}
		}
		if i != len(animals)-1 {
			isEatable = isCanEat(eatable, animals[i+1])
			if isEatable {
				word = fmt.Sprintf("%s eats %s", v, animals[i+1])
				animals[i+1] = ""
				break
			}
		}
	}

	stillAliveAnimals := []string{}
	for _, v := range animals {
		if v != "" {
			stillAliveAnimals = append(stillAliveAnimals, v)
		}
	}

	return stillAliveAnimals, isEatable, word
}

func isCanEat(eatableAnimals []string, target string) bool {
	for _, animal := range eatableAnimals {
		if animal == target {
			return true
		}
	}
	return false
}

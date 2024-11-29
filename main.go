package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	name := ""

	fmt.Println("Welcome to Jen's Casino...")
	fmt.Printf("Enter your name: ")

	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func getBet(balance uint8) uint8 {
	var bet uint8
	for true {
		fmt.Printf("Enter your bet or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
	return bet
}

func generateSymbolArray(symbols map[string]uint8) []string {
	symbolSlice := make([]string, 0, 43)
	for symbol, count := range symbols {
		for i := uint8(0); i < count; i++ {
			symbolSlice = append(symbolSlice, symbol)
		}
	}
	return symbolSlice
}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println("")
	}
}

func checkWin(spin [][]string, multipliers map[string]uint8) []uint8 {
	lines := []uint8{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
}

func main() {
	symbols := map[string]uint8{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint8 {
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	symbolSlice := generateSymbolArray(symbols)
	balance := uint8(200)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := getSpin(symbolSlice, 3, 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)
		fmt.Println(winningLines)

		for i, multiplier := range winningLines {
			win := multiplier * bet
			balance += win
			if multiplier > 0 {
				fmt.Printf("Won $%d, (%dx) on line #%d\n", win, multiplier, i + 1)
			}
		}
	}

	fmt.Printf("You left with $%d.\n", balance)
}

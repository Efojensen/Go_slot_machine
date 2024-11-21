package main

import "fmt"

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

func getBet(balance uint16) uint16 {
	var bet uint16
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
	symbolArr := make([]string, 8, 43)
	for symbol, count := range symbols {
		for i := uint8(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
}

func main() {
	symbols := map[string]uint8 {
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	// multipliers := map[string]uint8 {
	// 	"A": 20,
	// 	"B": 10,
	// 	"C": 5,
	// 	"D": 2,
	// }
	getName()

	balance := uint16(200)
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
	}

	fmt.Printf("You left with $%d.\n", balance)
}
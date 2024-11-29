package main

import "fmt"

func generateSymbolArray(symbols map[string]uint8) []string {
	symbolSlice := make([]string, 0, 43)
	for symbol, count := range symbols {
		for i := uint8(0); i < count; i++ {
			symbolSlice = append(symbolSlice, symbol)
		}
	}
	return symbolSlice
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
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := GetSpin(symbolSlice, 3, 3)
		PrintSpin(spin)
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

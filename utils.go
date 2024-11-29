package main

import "fmt"

func GetName() string {
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

func GetBet(balance uint8) uint8 {
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

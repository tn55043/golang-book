package main

import (
	"fmt"
	"github.com/tn55043/vendingMachine"
)

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}

func main() {
	//vm := VendingMachine{coins: coins, items: items}
	vm := vendingMachine.NewVendingMachine(coins, items)
	
	// Inserted Money = 0
	fmt.Println("Inserted Money :", vm.InsertedMoney())

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	// Inserted Money = 18
	fmt.Println("Inserted Money :", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can) // SD

	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	// Inserted Money = 12
	fmt.Println("Inserted Money :", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	// Inserted Money = 20
	fmt.Println("Inserted Money :", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	// Inserted Money = 25
	fmt.Println("Inserted Money :", vm.InsertedMoney())
	can = vm.CoinReturn()
	fmt.Println(can)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = 'balance.txt'

func main() {
    var accountBalance float64 = getBalanceFromFile() || 1000.00
	fmt.Println("\n=== Simple Bank Application ===")

	for {
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	var choice int
	fmt.Print("Choose an option (1-4): ")
	fmt.Scan(&choice);

	fmt.Println("Your choice:", choice)

// 	if (choice == 1) {
// 	    fmt.Println("Your balance is:", accountBalance);
// 	} else if (choice == 2) {
//         add(accountBalance);
// 	} else if (choice == 3) {
//         withdraw(accountBalance);
// 	} else if (choice == 4) {
// 	    fmt.Println("Exiting...")
// 	} else {
// 	    fmt.Println("Wrong choice, Exiting...");
// 	   // return
// 	   break;
// 	}
    switch choice {
        case 1:
    	    fmt.Println("Your balance is:", accountBalance);
    	case 2:
            add(accountBalance);
    	case 3:
            withdraw(accountBalance);
    	case 4:
    	    fmt.Println("Exiting...")
    	default:
    	    fmt.Println("Wrong choice, Exiting...");
    	    fmt.Println("Thank You !!!")
    	    return

    }
}

func add(accountBalance float64) {
    var balanceToAdd float64 = 0.00
	fmt.Println("Enter amount :")
	fmt.Scan(&balanceToAdd);
	if (balanceToAdd <= 0) {
        fmt.Println("Not Possible Lol");
        continue
    } else {
        accountBalance += balanceToAdd;
        fmt.Println("New balance: " accountBalance);
        writeBalanceToFile(accountBalance);
    }

}

func withdraw(accountBalance float64) {
    var balanceToWithdraw float64 = 0.00
    fmt.Println("Enter amount :")
    fmt.Scan(&balanceToWithdraw);
    if (balanceToWithdraw > accountBalance || balanceToWithdraw <=0) {
    	fmt.Println("Not Possible Lol");
    	continue
    } else {
    	accountBalance -= balanceToAdd;
        fmt.Println("New balance: " accountBalance);
        writeBalanceToFile(accountBalance);
    }
}
func writeBalanceToFile(balance float64) {
    balanceText := fmt.Sprint(balance)
    os.WriteFle(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
    data, _ := os.ReadFile(accountBalanceFile)
    balanceText := string(data)
    balance, _ := strconv.ParseFloat(balanceText, 64)
    return balance
}
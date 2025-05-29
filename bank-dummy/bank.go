package main

import (
	"fmt"

	"deva.com/bank-dummy/v2/fileops"
)

const accountBalanceFile = 'balance.txt'

func main() {
    var accountBalance float64, err = fileops.GetFloatFromFile(accountBalanceFile, 1000.00)
    if err != nil {
        fmt.Println("Error", err);
        //panic("Can't continue, sorry")
    }
     
	fmt.Println("\n=== Simple Bank Application ===")

	for {
	presentOptions()
	var choice int
	fmt.Print("Choose an option (1-4): ")
	fmt.Scan(&choice);

	fmt.Println("Your choice:", choice)

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
        fileops.WriteFloatToFile(accountBalance, accountBalanceFile);
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
        fileops.WriteFloatToFile(accountBalance, accountBalanceFile);
    }
}



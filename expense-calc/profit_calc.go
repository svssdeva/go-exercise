package main

import "fmt"

func main() {
    var revenue float64
    var expenses float64
    var taxRate float64

    fmt.Print("Revenue : ")
    fmt.Scan(&revenue)

    fmt.Print("Expenses : ")
    fmt.Scan(&expenses)

    fmt.Print("Tax Rate : ")
    fmt.Scan(&taxRate)

    ebt := revenue - expenses
    profit, ratio := calcProfit(ebt, taxRate)
    outputData(ebt, profit, ratio)

}

func outputData(ebt, profit, ratio float64) {
    fmt.Printf("EBT %v\n", ebt)
    fmt.Printf("Profit %.2f\n", profit)
    fmt.Println("Ratio", ratio)

    formattedEBT := fmt.Sprintf("EBT %v\n", ebt)
    formattedProfit := fmt.Sprintf("Profit %.2f\n", profit)
    fmt.Print(formattedEBT, formattedProfit)
}

func calcProfit(ebt, taxRate float64) (profit float64, ratio float64) {
    profit = ebt * (1- taxRate / 100)
    ratio = ebt/ profit
    //return profit, ratio
    return //this also works
}
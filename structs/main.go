package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "deva.com/structs/v2/note" // Adjust the import path as necessary
)
func main()  {
    title := getUserInput("Enter Title: ")
    content := getUserInput("Enter Content: ")
    userNote, err := note.New(title, content)
    if err != nil { 
        fmt.Println("Error creating note:", err)
        return
    }
    userNote.Display();
    err = userNote.Save();
    if err != nil {
        fmt.Println("Error saving note:", err)
        return
    }
    fmt.Println("Note saved successfully!")
}


func getUserInput(prompt string) string {
    fmt.Print("%v ",prompt)

    text, err := bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
        fmt.Println("Error reading input:", err)
        return ""
    }

    value := strings.TrimSuffix(text, "\n")
    value = strings.TrimSuffix(value, "\r") // Handle both Unix and Windows line endings
    value = strings.TrimSpace(value) // Remove any leading or trailing whitespace

    return value
}
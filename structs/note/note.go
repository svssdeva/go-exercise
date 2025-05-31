package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() string {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.Title, n.Content, n.CreatedAt.Format(time.RFC3339))
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreatedAt.Format(time.RFC3339))
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling note to JSON:", err)
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}

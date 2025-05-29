package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func WriteFloatToFile(value float64, fileName string) {
    valueText := fmt.Sprint(value)
    os.WriteFle(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
    data, err := os.ReadFile(fileName)
    if err != nil {
        return defaultValue, errors.New("Failed to fetch file")
    }
    valueText := string(data)
    value, err := strconv.ParseFloat(valueText, 64)
    if err != nil {
        return defaultValue, errors.New("Failed parse value")
    }
    return value, nil
}
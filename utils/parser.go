package utils

import (
	"bufio"
	"maxchat/models"
	"os"
	"strings"
)

func LoadData(filePath string) ([]models.Item, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []models.Item
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) < 7 {
			continue
		}
		tech := strings.Split(line[3]+","+line[4], ",")
		item := models.Item{
			Code:        line[0],
			Name:        line[1],
			Model:       line[2],
			Tech:        tech,
			Status:      line[5],
			Description: line[6],
		}
		items = append(items, item)
	}
	return items, scanner.Err()
}

func SaveData(filePath string, items []models.Item) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range items {
		line := strings.Join([]string{
			item.Code,
			item.Name,
			item.Model,
			strings.Join(item.Tech, ","),
			item.Status,
			item.Description,
		}, ",")
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

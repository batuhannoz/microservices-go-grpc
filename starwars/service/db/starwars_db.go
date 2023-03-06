package db

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Characters struct {
	Characters []Character `json:"characters"`
}

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	Homeworld string `json:"homeworld,omitempty"`
}

var db Characters

func InitializeDB() {
	plan, err := os.ReadFile("../db/db.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(plan, &db)
	if err != nil {
		fmt.Println(err)
	}
}

func GetRandomCharacter() Character {
	plan, err := os.ReadFile("../db/db.json")
	if err != nil {
		fmt.Println(err)
	}
	var data Characters
	err = json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data.Characters[rand.Intn(21)]
}

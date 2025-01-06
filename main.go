package main

import (
	"fmt"
	"log"
	"time"

	dnd5e "github.com/SergioPopovs176/dnd5-client/dnd-5e"
)

func main() {
	fmt.Println("Test DnD client ...")

	client, err := dnd5e.NewClient(time.Second * 15)
	if err != nil {
		log.Fatal(err)
	}

	// monsters, err := client.GetMonsters()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(monsters)

	monster, err := client.GetMonster("werewolf-hybrid")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(monster.Info())
}

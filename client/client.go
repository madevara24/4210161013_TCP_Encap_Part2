package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

type HeroData struct{
	Id		int
	XCoord	int
	YCoord	int
	HP		int
	Mana	int
	Gold	int
	Action	string
}

func main()  {
	fmt.Printf("Connecting to %s...\n", "localhost:1013")

	conn, errors := net.Dial("tcp","localhost:1013")
	fmt.Println("Connected!")
	if errors != nil {
		log.Fatal(errors)
		os.Exit(1)
	}

	for {
		hero := HeroData{Id:4210,Mana:800,HP:1000,XCoord:115,YCoord:129,Gold:476,Action:"Idle"}
		encoder := gob.NewEncoder(conn)
		encoder.Encode(hero)

		handleConnection(conn)
		conn.Close()
		break
	}
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	hero := &HeroData{}
	dec.Decode(hero)
	fmt.Println(
		"Hero ID : ", hero.Id, 
		"\nX Position : ", hero.XCoord, 
		"\nY Position : ", hero.YCoord, 
		"\nHP : ", hero.HP, 
		"\nMana : ", hero.Mana, 
		"\nGold : ", hero.Gold,
		"\nAction : ", hero.Action)
}
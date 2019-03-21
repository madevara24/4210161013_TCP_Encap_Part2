package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
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
	listener, errors :=net.Listen("tcp","localhost:1013")
	if errors != nil{
		log.Fatal(errors)
	}
	fmt.Println("Starting server side...")

	for {
		conn, errors := listener.Accept()
		if errors != nil{
			log.Fatal(errors)
		}
		decoder :=gob.NewDecoder(conn)
		hero := &HeroData{}
		decoder.Decode(hero)

		fmt.Println(
			"Hero ID : ", hero.Id, 
			"\nX Position : ", hero.XCoord, 
			"\nY Position : ", hero.YCoord, 
			"\nHP : ", hero.HP, 
			"\nMana : ", hero.Mana, 
			"\nGold : ", hero.Gold,
			"\nAction : ", hero.Action)

		handleConnection(hero, conn)
		conn.Close()
	}
}

func handleConnection(hero *HeroData, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(hero)
}
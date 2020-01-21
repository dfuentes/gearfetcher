package main

import (
	"flag"
	"log"
)

var characterName string
var server string

const Region = "US"

var TestChar = Character{
	Class: Mage,
	Race:  Undead,
	Slots: map[int]Slot{
		1: {Item: 16795},
		2: {Item: 18814},
		3: {Item: 11782},
		// 4: {Item: 4334},
		5:  {Item: 14152},
		6:  {Item: 19136},
		7:  {Item: 13170},
		8:  {Item: 22860},
		9:  {Item: 11766},
		10: {Item: 22870},
		11: {Item: 942},
		12: {Item: 12545},
		13: {Item: 13968},
		14: {Item: 12930},
		15: {Item: 10212},
		16: {Item: 18842},
		18: {Item: 19108},
	},
}

func init() {
	flag.StringVar(&characterName, "name", "", "character to import")
	flag.StringVar(&server, "server", "Atiesh", "server for character")
	flag.Parse()
}

func main() {
	c := NewClient("b0309af3125664894e80256739b7b12d")
	// parses, err := c.GetParses(ParsesQuery{
	// 	CharacterName: characterName,
	// 	Server:        server,
	// 	Region:        "US",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range parses {

	// 	fmt.Printf("%+v\n", p)
	// }

	// character := ParseToCharacter(parses[0])
	// fmt.Println(EncodeCharacter(character))

	s := NewServer(ServerConfig{
		WLClient: c,
	})

	log.Fatal(s.ListenAndServe())
}

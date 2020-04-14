package main

import "strings"

type Character struct {
	Class      Class
	Level      int
	Race       Race
	Slots      Slots
	TalentHash string
	Version    int
}

type Slots map[int]Slot

type Slot struct {
	Item          int
	Enchant       int
	RandomEnchant int
}

type Class int

const (
	UnknownClass Class = iota
	Warrior
	Paladin
	Hunter
	Rogue
	Priest
	Shaman
	Mage
	Warlock
	Druid
)

func (c Class) String() string {
	switch c {
	case UnknownClass:
		return "unknown"
	case Warrior:
		return "warrior"
	case Paladin:
		return "paladin"
	case Hunter:
		return "hunter"
	case Rogue:
		return "rogue"
	case Priest:
		return "priest"
	case Shaman:
		return "shaman"
	case Mage:
		return "mage"
	case Warlock:
		return "warlock"
	case Druid:
		return "druid"
	default:
		return "INVALID"
	}
}

func ParseClass(s string) Class {
	s = strings.ToLower(s)
	switch s {
	case "unknown":
		return UnknownClass
	case "warrior":
		return Warrior
	case "paladin":
		return Paladin
	case "hunter":
		return Hunter
	case "rogue":
		return Rogue
	case "priest":
		return Priest
	case "shaman":
		return Shaman
	case "mage":
		return Mage
	case "warlock":
		return Warlock
	case "druid":
		return Druid
	default:
		return UnknownClass
	}
}

type Race int

const (
	UnknownRace Race = iota
	Human
	Orc
	Dwarf
	NightElf
	Undead
	Tauren
	Gnome
	Troll
)

func (r Race) String() string {
	switch r {
	case UnknownRace:
		return "unknown"
	case Human:
		return "human"
	case Orc:
		return "orc"
	case Dwarf:
		return "dwarf"
	case NightElf:
		return "night-elf"
	case Undead:
		return "undead"
	case Tauren:
		return "tauren"
	case Gnome:
		return "gnome"
	case Troll:
		return "troll"
	default:
		return "INVALID"
	}
}

type ImportType int

const (
	GearImport ImportType = iota
	GearAndLevelImport
	GearAndLevelAndTalentsImport
)

func ParseToCharacter(p Parse) Character {
	c := Character{
		Class:      ParseClass(p.Class),
		Level:      60,
		Race:       UnknownRace,
		Slots:      map[int]Slot{},
		TalentHash: "",
		Version:    1,
	}

	// Warcraftlogs does not give us character race.
	// Instead just pick a default race based on class :/

	switch c.Class {
	case Druid:
		c.Race = Tauren
	case Hunter:
		c.Race = Troll
	case Mage:
		c.Race = Troll
	case Paladin:
		c.Race = Human
	case Priest:
		c.Race = Undead
	case Rogue:
		c.Race = Orc
	case Shaman:
		c.Race = Troll
	case Warlock:
		c.Race = Undead
	case Warrior:
		c.Race = Orc
	}

	for ix, item := range p.Gear {
		itemSlot := ix + 1

		if itemSlot == 4 {
			// Skip shirt slot
			continue
		}

		c.Slots[itemSlot] = Slot{
			Item: item.ID,
		}
	}

	return c
}

type ParsesQuery struct {
	CharacterName string
	Server        string
	Region        string
}

type ParsesResponse []Parse

type Parse struct {
	CharacterID    int     `json:"characterID"`
	CharacterName  string  `json:"characterName"`
	Class          string  `json:"class"`
	Difficulty     int     `json:"difficulty"`
	Duration       int     `json:"duration"`
	EncounterID    int     `json:"encounterID"`
	EncounterName  string  `json:"encounterName"`
	Estimated      bool    `json:"estimated"`
	FightID        int     `json:"fightID"`
	Gear           []Gear  `json:"gear"`
	IlvlKeyOrPatch int     `json:"ilvlKeyOrPatch"`
	OutOf          int     `json:"outOf"`
	Percentile     float64 `json:"percentile"`
	Rank           int     `json:"rank"`
	ReportID       string  `json:"reportID"`
	Server         string  `json:"server"`
	Spec           string  `json:"spec"`
	StartTime      int     `json:"startTime"`
	Total          float64 `json:"total"`
}

type Gear struct {
	Icon    string `json:"icon"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Quality string `json:"quality"`
}

func (p Parse) GetGearLink() string {
	character := ParseToCharacter(p)
	return EncodeCharacter(character)
}

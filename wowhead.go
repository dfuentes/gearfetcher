package main

import (
	"encoding/base64"
	"log"
	"sort"
	"strings"
)

// EncodeCharacter takes a character and returns the wowhead gear planner link for that character
func EncodeCharacter(character Character) string {
	path := "/gear-planner/" + character.Class.String() + "/" + character.Race.String() + "/"

	encodedCharacter := []byte{2, byte(character.Level)}

	if character.TalentHash != "" {
		// TODO: handle talents
	}

	encodedCharacter = append(encodedCharacter, 0)
	keys := []int{}
	for k, _ := range character.Slots {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		value := character.Slots[key]
		item := value.Item
		enchanted := value.Enchant > 0
		randomEnchant := value.RandomEnchant > 0

		if enchanted {
			key = key | 0x80
		}

		if randomEnchant {
			key = key | 0x40
		}

		encodedCharacter = append(encodedCharacter, byte(key))
		encodedCharacter = append(encodedCharacter, byte((item>>8)&255), byte(item&255))
		if enchanted {
			encodedCharacter = append(encodedCharacter, byte((value.Enchant>>8)&255), byte(value.Enchant&255))
		}
		if randomEnchant {
			encodedCharacter = append(encodedCharacter, byte((value.RandomEnchant>>8)&255), byte(value.RandomEnchant&255))
		}
	}

	ec := base64.StdEncoding.EncodeToString(encodedCharacter)
	ec = strings.ReplaceAll(ec, "+", "-")
	ec = strings.ReplaceAll(ec, "/", "_")
	ec = strings.ReplaceAll(ec, "=", "")
	ec = strings.ReplaceAll(ec, "+", "")
	ec = strings.ReplaceAll(ec, "$", "")

	return "https://classic.wowhead.com" + path + ec
}

// DecodeCharacter takes the encoded character data from a wowhead gear-planner link and returns a character with that gear
func DecodeCharacter(input string) Character {
	character := Character{
		Class:      UnknownClass,
		Level:      60,
		Race:       UnknownRace,
		Slots:      Slots{},
		TalentHash: "",
		Version:    1,
	}

	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		// FIXME
		log.Fatal(err)
	}

	importType, decoded := shift(decoded)

	// Decode Level Info
	if ImportType(importType) > GearImport {
		var levelByte int
		levelByte, decoded = shift(decoded)
		character.Level = int(levelByte)
	}

	// Decode Talent Info
	if ImportType(importType) == GearAndLevelAndTalentsImport {
		// TODO: handle talent data
		var talentLength int
		var talentData []byte

		talentLength, decoded = shift(decoded)
		talentData, decoded = splice(decoded, 0, int(talentLength))
		_ = talentData
	}

	// Decode Gear Info
	for len(decoded) >= 3 {
		var slotID int
		var item int
		var itemLower int

		slotID, decoded = shift(decoded)
		item, decoded = shift(decoded)
		itemLower, decoded = shift(decoded)
		item = (item << 8) | itemLower

		enchanted := (slotID & 0x80) > 0
		randomEnchant := (slotID & 0x40) > 0

		slotID = slotID & ^0x80 & ^0x40

		slot := Slot{
			Item: int(item),
		}

		if enchanted {
			var enchant int
			var enchantLower int
			enchant, decoded = shift(decoded)
			enchantLower, decoded = shift(decoded)
			enchant = (enchant << 8) | enchantLower
			slot.Enchant = int(enchant)
		}

		if randomEnchant {
			var renchant int
			var renchantLower int
			renchant, decoded = shift(decoded)
			renchantLower, decoded = shift(decoded)
			renchant = (renchant << 8) | renchantLower
			slot.RandomEnchant = int(renchant)
		}

		character.Slots[int(slotID)] = slot
	}

	return character
}

func shift(sl []byte) (int, []byte) {
	return int(sl[0]), sl[1:]
}

func splice(sl []byte, start int, end int) ([]byte, []byte) {
	return sl[start:end], sl[end:]
}

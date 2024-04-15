package mythbeings

import "strconv"

type Being struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}

func beings() []Being {
	return []Being{
		{
			ID:          1,
			Name:        "Phoenix",
			Origin:      "Ancient Greek Mythology",
			Description: "A majestic bird known for its cycle of death and rebirth in flames.",
		},
		{
			ID:          2,
			Name:        "Kitsune",
			Origin:      "Japanese Folklore",
			Description: "A fox spirit with the power to transform and the wisdom of ages.",
		},
		{
			ID:          3,
			Name:        "Leviathan",
			Origin:      "Jewish Mythology",
			Description: "A massive sea monster symbolizing chaos and feared by sailors.",
		},
		{
			ID:          4,
			Name:        "Banshee",
			Origin:      "Irish Mythology",
			Description: "A wailing spirit whose cry is believed to herald the death of a family member.",
		},
		{
			ID:          5,
			Name:        "Anansi",
			Origin:      "West African Folklore",
			Description: "A spider deity known for its intelligence and ability to outsmart others.",
		},
		{
			ID:          6,
			Name:        "Golem",
			Origin:      "Jewish Folklore",
			Description: "An animated anthropomorphic being, created magically to serve and protect.",
		},
	}
}

func loadBeings() map[string]Being {
	beingsList := beings()
	res := make(map[string]Being, len(beingsList))

	for _, b := range beingsList {
		res[strconv.Itoa(int(b.ID))] = b
	}

	return res
}

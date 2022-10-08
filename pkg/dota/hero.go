package dota

type Hero struct {
	Id int `json:"id"`
	Name string `json:"name"`
	LocalizedName string `json:"localized_name"`
	PrimaryAttr string `json:"primary_attr"`
	AttackType string `json:"attack_type"`
	Roles []string `json:"roles"`
	Legs int `json:"legs"`
}

type Heroes struct {
	Heroes []Hero
}
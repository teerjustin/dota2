package dota

type Icon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	LocalizedName string `json:"localized_name"`
	Img string `json:"img"`
	Icon string `json:"icon"`
}

type Icons struct {
	Icons []Icon 
}
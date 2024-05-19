package models

type Biography struct {
	FullName string `json:"fullName"`
}

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Images struct {
	Xs string `json:"xs"`
	Sm string `json:"sm"`
	Md string `json:"md"`
	Lg string `json:"lg"`
}

type Superhero struct {
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	Powerstats Powerstats `json:"powerstats"`
	Images     Images     `json:"images"`
}

package main

import (
	"log"
	"net/http"

	"github.com/Juanpa93/superhero/handlers"
	"github.com/Juanpa93/superhero/models"
	"github.com/Juanpa93/superhero/repository"
)

func main() {
	mux := http.NewServeMux()
	handler := &handlers.HandlerSuperheros{
		BD: bd,
	}
	mux.HandleFunc("GET /api/superhero", handler.TraerSuperhero())
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var bd = &repository.BaseDatos{
	Memoria: map[string]models.Superhero{
		"Wolverine": {
			Name: "Wolverine",
			Biography: models.Biography{
				FullName: "John Logan",
			},
			Powerstats: models.Powerstats{
				Intelligence: 63,
				Strength:     32,
				Speed:        50,
				Durability:   100,
				Power:        89,
				Combat:       100,
			},
			Images: models.Images{
				Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
				Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
				Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
				Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
			},
		},
		"Spider-Man": {
			Name: "Spider-Man",
			Biography: models.Biography{
				FullName: "Peter Parker",
			},
			Powerstats: models.Powerstats{
				Intelligence: 90,
				Strength:     55,
				Speed:        67,
				Durability:   75,
				Power:        74,
				Combat:       85,
			},
			Images: models.Images{
				Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
				Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
				Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
				Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
			},
		},
		"Iron Man": {
			Name: "Iron Man",
			Biography: models.Biography{
				FullName: "Tony Stark",
			},
			Powerstats: models.Powerstats{
				Intelligence: 100,
				Strength:     85,
				Speed:        58,
				Durability:   85,
				Power:        100,
				Combat:       64,
			},
			Images: models.Images{
				Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
				Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
				Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
				Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
			},
		},
		"Black Widow": {
			Name: "Black Widow",
			Biography: models.Biography{
				FullName: "Natasha Romanoff",
			},
			Powerstats: models.Powerstats{
				Intelligence: 75,
				Strength:     13,
				Speed:        33,
				Durability:   30,
				Power:        36,
				Combat:       100,
			},
			Images: models.Images{
				Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
				Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
				Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
				Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
			},
		},
		"Thor": {
			Name: "Thor",
			Biography: models.Biography{
				FullName: "Thor Odinson",
			},
			Powerstats: models.Powerstats{
				Intelligence: 69,
				Strength:     100,
				Speed:        83,
				Durability:   100,
				Power:        100,
				Combat:       100,
			},
			Images: models.Images{
				Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
				Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
				Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
				Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
			},
		},
	},
}

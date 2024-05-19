package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/Juanpa93/superhero/repository"
)

type HandlerSuperheros struct {
	BD *repository.BaseDatos
}

func NewHandlerSuperheros() *HandlerSuperheros {
	return &HandlerSuperheros{}
}

func (hc *HandlerSuperheros) TraerSuperhero() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("hero")
		if id == "" {
			http.Error(w, "No se encontro un nombre valido", http.StatusBadRequest)
			return
		}

		nombre := capitalize(id)
		superhero, ok := hc.BD.Memoria[nombre]
		if !ok {
			http.Error(w, "No se encontro el superhero", http.StatusNotFound)
			return
		}

		payload, err := json.Marshal(superhero)
		if err != nil {
			http.Error(w, "Fallo la condificacion a Json", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)

	})
}

func capitalize(str string) string {
	words := strings.Fields(str)
	var palabras []string
	var result string
	var convertir bool
	for _, word := range words {
		if len(word) > 0 {
			palabrasCap := strings.ToUpper(string(word[0])) + word[1:]
			palabras = append(palabras, palabrasCap)
		}
	}
	auxCap := strings.Join(palabras, " ")

	for _, char := range auxCap {
		if convertir {
			result += string(unicode.ToUpper(char))
			convertir = false
		} else {
			result += string(char)
		}

		if char == '-' {
			convertir = true
		}
	}
	fmt.Println(result)
	return result

}

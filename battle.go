package main

import (
	"github.com/gin-gonic/gin"
)

func ValidateTurnForm(c *gin.Context) {
	// El middleware valida token y participación en la batalla.
	// Binding de la petición contra TurnForm.
	// Comprobar que el turno es el del usuario.
	// Comprobar que el destino es válido y no está ocupado.
	// Si 'current_pokemon' comprobar que está en rango y no debilitado.
	// Si 'movement' comprobar que está en rango.
	// Si 'current_pokemon' y 'movement' presentes devolver error.
	// Llamada a función que aplique los cambios y genere respuesta.
}

type TurnForm struct {
	CurrentPokemon int `binding:"max=5,min=0" json:"current_pokemon"`
	Movement       int `binding:"max=3,min=0" json:"movement"`
	Position       int `binding:"max=24,min=0,required" json:"position"`
}

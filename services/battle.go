package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/log"
	"github.com/jrevillas/pkmnrequiem-go/middlewares"
	"github.com/jrevillas/pkmnrequiem-go/models"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type BattleService struct {
	*middlewares.Session
	db *mgo.Database
}

func NewBattleService(db *mgo.Database) *BattleService {
	return &BattleService{
		db:      db,
		Session: middlewares.NewSession(db),
	}
}

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

type Battle struct {
	ID                string       `json:"id"`
	Current_pokemon_1 int          `json:"current_pokemon_1"`
	Current_pokemon_2 int          `json:"current_pokemon_2"`
	Finished          bool         `json:"finished"`
	Log               []string     `json:"log"`
	Party_1           []string     `json:"party_1"`
	Party_2           []string     `json:"party_2"`
	Position_1        int          `json:"position_1"`
	Position_2        int          `json:"position_2"`
	Turn              int          `json:"turn"`
	User_1            *models.User `json:"user_1"`
	User_2            *models.User `json:"user_2"`
}

func (b *BattleService) Register(r *gin.RouterGroup) {
	g := r.Group("/battle")
	g.GET("/example", b.Guest, b.Example)
}

func (b *BattleService) Example(c *gin.Context) {
	start := time.Now()
	example := NewExampleBattle()
	c.IndentedJSON(http.StatusCreated, example)
	log.Debug("Batalla de ejemplo generada (%.2fms)", time.Since(start).Seconds()*1000)
}

func NewExampleBattle() *Battle {
	start := time.Now()
	user1 := models.NewUser("example@example.com", "12345", "demouser1")
	log.Debug("Usuario temporal '%s' generado (%.2fms)", user1.Username, time.Since(start).Seconds()*1000)
	start = time.Now()
	user2 := models.NewUser("example@example.com", "12345", "demouser2")
	log.Debug("Usuario temporal '%s' generado (%.2fms)", user2.Username, time.Since(start).Seconds()*1000)
	return &Battle{
		ID:                uuid.NewV4().String(),
		Log:               []string{},
		Party_1:           []string{},
		Party_2:           []string{},
		Position_1:        2,
		Position_2:        22,
		User_1:            user1,
		User_2:            user2,
	}
}

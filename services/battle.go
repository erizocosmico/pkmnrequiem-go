package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/middlewares"
	"github.com/jrevillas/pkmnrequiem-go/models"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
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

type Battle struct {
	ID                string       `json:"id"`
	Attack_phase      bool         `json:"attack_phase"`
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
	example := NewBattle("demouser1", "demouser2")
	c.IndentedJSON(http.StatusCreated, example)
}

func NewBattle(username1, username2 string) *Battle {
	user1 := models.NewUser("example@example.com", "12345", username1)
	user2 := models.NewUser("example@example.com", "12345", username2)
	return &Battle{
		ID:                uuid.NewV4().String(),
		Attack_phase:      false,
		Current_pokemon_1: 0,
		Current_pokemon_2: 0,
		Finished:          false,
		Log:               []string{},
		Party_1:           []string{},
		Party_2:           []string{},
		Position_1:        2,
		Position_2:        22,
		Turn:              0,
		User_1:            user1,
		User_2:            user2,
	}
}
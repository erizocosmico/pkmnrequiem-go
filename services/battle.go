package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/middlewares"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

// BattleService ...
type BattleService struct {
	*middlewares.Session
	db *mgo.Database
}

// NewBattleService ...
func NewBattleService(db *mgo.Database) *BattleService {
	return &BattleService{
		db:      db,
		Session: middlewares.NewSession(db),
	}
}

// Battle ...
type Battle struct {
	ID           string   `json:"id"`
	AttackPhase  bool     `json:"attack_phase"`
	CurrentPkmn1 int      `json:"current_pokemon_1"`
	CurrentPkmn2 int      `json:"current_pokemon_2"`
	Finished     bool     `json:"finished"`
	Log          []string `json:"log"`
	Party1       []string `json:"party_1"`
	Party2       []string `json:"party_2"`
	Position1    int      `json:"position_1"`
	Position2    int      `json:"position_2"`
	Turn         int      `json:"turn"`
	User1        string   `json:"user_1"`
	User2        string   `json:"user_2"`
}

// Register ...
func (b *BattleService) Register(r *gin.RouterGroup) {
	g := r.Group("/battle")
	g.GET("/example", b.Guest, b.Example)
}

// Example ...
func (b *BattleService) Example(c *gin.Context) {
	example := NewBattle("demouser1", "demouser2")
	c.IndentedJSON(http.StatusCreated, example)
}

// NewBattle ...
func NewBattle(user1, user2 string) *Battle {
	return &Battle{
		ID:           uuid.NewV4().String(),
		AttackPhase:  false,
		CurrentPkmn1: 0,
		CurrentPkmn2: 0,
		Finished:     false,
		Log:          []string{},
		Party1:       []string{},
		Party2:       []string{},
		Position1:    2,
		Position2:    22,
		Turn:         0,
		User1:        user1,
		User2:        user2,
	}
}

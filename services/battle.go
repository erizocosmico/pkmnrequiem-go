package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/middlewares"
	"github.com/jrevillas/pkmnrequiem-go/models"
	"github.com/op/go-logging"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type BattleService struct {
	*middlewares.Session
	db *mgo.Database
	log *logging.Logger
}

func NewBattleService(db *mgo.Database, log *logging.Logger) *BattleService {
	return &BattleService{
		db:      db,
		log:     log,
		Session: middlewares.NewSession(db),
	}
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
	example := NewBattle("demouser1", "demouser2")
	c.IndentedJSON(http.StatusCreated, example)
	b.log.Debugf("Batalla de ejemplo generada (%.2fms)", time.Since(start).Seconds()*1000)
}

func NewBattle(username1, username2 string) *Battle {
	user1 := models.NewUser("example@example.com", "12345", username1)
	user2 := models.NewUser("example@example.com", "12345", username2)
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

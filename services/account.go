package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrevillas/pkmnrequiem-go/middlewares"
	"github.com/jrevillas/pkmnrequiem-go/models"
	"gopkg.in/mgo.v2"
)

// AccountService ...
type AccountService struct {
	*middlewares.Session
	collection *mgo.Collection
	store      *models.UserStore
}

// NewAccountService ...
func NewAccountService(db *mgo.Database) *AccountService {
	return &AccountService{
		collection: db.C("account"),
		Session:    middlewares.NewSession(db),
		store:      &models.UserStore{db},
	}
}

// Register ...
func (a *AccountService) Register(r *gin.RouterGroup) {
	group := r.Group("/account")
	group.POST("/create", a.Guest, a.Create)
	// group.POST("/login", a.Guest, a.Login)
	// group.POST("/logout", a.Auth, a.Logout)
}

// Create ...
func (a *AccountService) Create(c *gin.Context) {
	var form CreateAccountForm
	if err := c.BindJSON(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, form)
}

// CreateAccountForm ...
type CreateAccountForm struct {
	Email    string `binding:"email,required" json:"email"`
	Password string `binding:"min=8,required" json:"password"`
	Username string `binding:"alphanum,required" json:"username"`
}

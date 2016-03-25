package services

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/jrevillas/pkmnrequiem-go/middlewares"
  "github.com/jrevillas/pkmnrequiem-go/models"
  "gopkg.in/mgo.v2"
)

type AccountService struct {
  *middlewares.Session
  store *models.UserStore
}

func NewAccountService(db *mgo.Database) *AccountService {
  return &AccountService{
    collection: db.C("account"),
    Session:    middlewares.NewSession(db),
    store:      &models.UserStore{db},
  }
}

func (a *AccountService) Register(r *gin.RouterGroup) {
  group := r.Group("/account")
  g.POST("/create", a.Guest, a.Create)
  // g.POST("/login", a.Guest, a.Login)
  // g.POST("/logout", a.Auth, a.Logout)
}

func (a *AccountService) Create(c *gin.Context) {
  var form CreateAccountForm
  if err := c.BindJSON(&form); err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, form)
}

type CreateAccountForm struct {
  Email    string `binding:"email,required" json:"email"`
  Password string `binding:"min=8,required" json:"password"`
  Username string `binding:"alphanum,required" json:"username"`
}

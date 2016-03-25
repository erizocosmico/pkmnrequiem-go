package middlewares

import (
  "net/http"
  "strings"
  "github.com/gin-gonic/gin"
  "github.com/jrevillas/pkmnrequiem-go/models"
  "gopkg.in/mgo.v2"
)

const (
  AuthKey    = "Authorization"
  TokenKey   = "session_token"
  TokenParam = "token"
  UserKey    = "session_user"
)

type Session struct {
  db    *mgo.Database
  store models.UserStore
}

func NewSession(db *mgo.Database) *Session {
  return &Session{db: db, store: models.UserStore{db}}
}

func (s *Session) Auth(c *gin.Context) {
  return nil
}

func (s *Session) Guest(c *gin.Context) {
  return nil
}

func retrieveToken(c *gin.Context) string {
  authHeader := c.Request.Header.Get(AuthKey)
  if authHeader != "" {
    parts := strings.SplitN(authHeader, " ", 1)
    return parts[len(parts) - 1]
  }
  return authHeader
}

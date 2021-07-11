package tinyurl

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"strconv"
	"time"
)

type Csrf struct {
	secret string
	session *scs.SessionManager
}

func NewCsrf(secret string, session *scs.SessionManager) *Csrf {
	return &Csrf{secret: secret, session: session}
}


func (c *Csrf) CreateToken(ctx context.Context) string {
	now := time.Now().Unix()
	h := hmac.New(sha256.New, []byte(c.secret))
	h.Write([]byte(strconv.FormatInt(now,10)))

	token := fmt.Sprintf("%x", h.Sum(nil))

	c.session.Put(ctx,"csrf_token",token)

	return token
}

func (c *Csrf) Verify(ctx context.Context,token string) bool {
	csrfToken := c.session.GetString(ctx,"csrf_token")
	if csrfToken == "" {
		return false
	}
	if csrfToken != token {
		return false
	}
	return true
}

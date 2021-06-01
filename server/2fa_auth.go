package chserver

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"sync"
	"time"

	errors2 "github.com/cloudradar-monitoring/rport/server/api/errors"
	"github.com/cloudradar-monitoring/rport/server/api/message"
	"github.com/cloudradar-monitoring/rport/share/security"
)

type TwoFAService struct {
	TokenTTL time.Duration
	MsgSrv   message.Service
	UserSrv  UserService

	tokensByUser map[string]*expirableToken
	mu           sync.RWMutex
}

func NewTwoFAService(tokenTTLSeconds int, userSrv UserService, msgSrv message.Service) TwoFAService {
	return TwoFAService{
		TokenTTL:     time.Duration(tokenTTLSeconds) * time.Second,
		UserSrv:      userSrv,
		MsgSrv:       msgSrv,
		tokensByUser: make(map[string]*expirableToken),
	}
}

type expirableToken struct {
	token  string
	expiry time.Time
}

const twoFATokenLength = 6

// TODO: add tests
func (srv *TwoFAService) SendToken(username string) (sendTo string, err error) {
	if username == "" {
		return "", errors2.APIError{
			Message: "username cannot be empty",
			Code:    http.StatusBadRequest,
		}
	}

	user, err := srv.UserSrv.GetByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors2.APIError{
			Message: fmt.Sprintf("user with username %s not found", username),
			Code:    http.StatusNotFound,
		}
	}

	if user.TwoFASendTo == "" {
		return "", errors2.APIError{
			Message: "no two_fa_send_to set for this user",
			Code:    http.StatusBadRequest,
		}
	}

	token, err := security.NewRandomToken(twoFATokenLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate 2fa token: %wv", err)
	}

	msg := fmt.Sprintf("Token: %s (valid %s)", token, srv.TokenTTL)
	if err := srv.MsgSrv.Send("Rport 2FA token", msg, user.TwoFASendTo); err != nil {
		return "", fmt.Errorf("failed to send 2fa token: %w", err)
	}

	srv.mu.Lock()
	srv.tokensByUser[username] = &expirableToken{
		token:  token,
		expiry: time.Now().Add(srv.TokenTTL),
	}
	srv.mu.Unlock()

	return user.TwoFASendTo, nil
}

// TODO: add tests
func (srv *TwoFAService) ValidateToken(username, token string) error {
	srv.mu.RLock()
	t := srv.tokensByUser[username]
	defer srv.mu.RUnlock()

	if t == nil {
		return errors2.APIError{
			Message: "2fa token not found for provided username",
			Code:    http.StatusUnauthorized,
		}
	}

	if time.Now().After(t.expiry) {
		return errors2.APIError{
			Message: "2fa token expired",
			Code:    http.StatusUnauthorized,
		}
	}

	if subtle.ConstantTimeCompare([]byte(t.token), []byte(token)) != 1 {
		return errors2.APIError{
			Message: "invalid token",
			Code:    http.StatusUnauthorized,
		}
	}

	return nil
}

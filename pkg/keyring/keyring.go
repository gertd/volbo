package keyring

import (
	"os/user"

	"github.com/pkg/errors"
	"github.com/zalando/go-keyring"
)

type KeyRing struct {
	service string
	user    string
}

func New(key string) (*KeyRing, error) {
	u, err := user.Current()
	if err != nil {
		return nil, errors.Wrapf(err, "get username")
	}
	return &KeyRing{
		service: key,
		user:    u.Username,
	}, nil
}

func (kr *KeyRing) GetToken() (string, error) {
	tokenStr, err := keyring.Get(kr.service, kr.user)
	if err != nil {
		return "", errors.Wrapf(err, "get token")
	}
	return tokenStr, nil
}

func (kr *KeyRing) SetToken(token string) error {
	if err := keyring.Set(kr.service, kr.user, token); err != nil {
		return errors.Wrapf(err, "set token")
	}
	return nil
}

func (kr *KeyRing) DelToken() error {
	if err := keyring.Delete(kr.service, kr.user); err != nil {
		return errors.Wrapf(err, "del token")
	}
	return nil
}

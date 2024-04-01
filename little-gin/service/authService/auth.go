package authService

import "github.com/azusachino/golong/little-gin/model"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return model.CheckAuth(a.Username, a.Password)
}

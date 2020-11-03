package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/utils"
	"gorm.io/gorm"
)

// Auth => Struct for auth atribute
type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// AuthModel => Struct for init DB
type AuthModel struct {
	DB *gorm.DB
}

// Login => Authentication for login
func (model AuthModel) Login(auth Auth) (bool, error, string) {
	var account Account
	result := model.DB.Where(&Account{Name: auth.Name}).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("account not found"), ""
		}

		return false, result.Error, ""
	}

	err := utils.HashComparator([]byte(account.Password), []byte(auth.Password))

	if err != nil {
		return false, errors.Errorf("incorrect password"), ""
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":            auth.Name,
		"account_account": account.AccountNumber,
	})

	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return false, err, ""
	}

	return true, nil, token
}

package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/model"
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/utils"
	"gorm.io/gorm"
)

// AccountController => Struct for controlling account
type AccountController struct {
	DB *gorm.DB
}

// CreateAccount => Controller for creating an account
func (ctrl AccountController) CreateAccount(ctx *gin.Context) {
	accountModel := model.AccountModel{
		DB: ctrl.DB,
	}
	var account model.Account

	err := ctx.Bind(&account)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := utils.HashGenerator(account.Password)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	account.Password = hashPassword

	flag, err := accountModel.InsertNewAccount(account)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown failed to insert account", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
}

// GetAccount => Controller for Getting an account
func (ctrl AccountController) GetAccount(ctx *gin.Context) {
	idAccount := ctx.MustGet("account_number").(int)
	accountModel := model.AccountModel{
		DB: ctrl.DB,
	}
	flag, err, transactions, account := accountModel.GetAccountDetail(idAccount)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unkown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"account":     account,
		"transaction": transactions,
	}, http.StatusOK, "success")
	return
}

// Login => Controller for loging
func (ctrl AccountController) Login(ctx *gin.Context) {
	authModel := model.AuthModel{
		DB: ctrl.DB,
	}
	var auth model.Auth

	err := ctx.Bind(&auth)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err, token := authModel.Login(auth)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"token": token,
	}, http.StatusOK, "success")

	log.Println("Login")
}

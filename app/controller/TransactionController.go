package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/model"
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/utils"
	"gorm.io/gorm"
)

// TransactionController => Struct for controlling transaction
type TransactionController struct {
	DB *gorm.DB
}

// Transfer => Controller for transfer
func (ctrl TransactionController) Transfer(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

// Withdraw => Controller for Withdraw
func (ctrl TransactionController) Withdraw(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

// Deposit => Controller for Deposit
func (ctrl TransactionController) Deposit(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

package controllers

import (
	"encoding/json"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"

	"github.com/astaxie/beego"
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Post() {
	var transaction *models.Transaction
	json.Unmarshal(c.Ctx.Input.RequestBody, &transaction)
	transaction = repositories.Transactions.AddOne(*transaction)

	c.Data["json"] = map[string]models.Transaction{"transaction": *transaction}
	c.ServeJSON()
}

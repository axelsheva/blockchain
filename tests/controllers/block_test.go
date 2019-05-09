package controllers

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/controllers"
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

func TestOnReveice(t *testing.T) {
	block := models.Block{
		ID: "2",
	}

	err := controllers.Block.OnReceive(&block)
	if err != nil {
		t.Errorf("Unable to accept block. Error: %s", err)
	}

	isExists := repositories.Blocks.IsExists("2")
	if isExists != true {
		t.Errorf("Invalid receive block. Block does not exist in repository")
	}
}

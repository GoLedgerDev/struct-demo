package main

import (
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/goledgerdev/struct-demo/chaincode/txdefs"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.CreateNewHarvest,
	// txdefs.GetNumberOfBooksFromLibrary,
	// txdefs.UpdateBookTenant,
}

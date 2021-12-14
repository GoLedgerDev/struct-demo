package txdefs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var ProduceWine = tx.Transaction{
	Tag:         "produceWine",
	Label:       "Produce Wine",
	Description: "Produce Wine",
	Method:      "POST",
	Callers:     []string{"$org3MSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "originHarvest",
			Label:    "Origin harvest",
			DataType: "->harvest",
		},
		{
			Required: true,
			Tag:      "wineType",
			Label:    "Wine type",
			DataType: "wineType",
		},
		{
			Required: true,
			Tag:      "productionDate",
			Label:    "Production Date",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		originHarvest, _ := req["originHarvest"].(assets.Key)
		wineType, _ := req["wineType"].(float64)
		productionDate, _ := req["productionDate"].(time.Time)
		name, _ := req["name"].(string)

		actorMap := map[string]interface{}{
			"@assetType": "actor",
			"name":       "vinicola",
		}

		actorKey, err := assets.NewKey(actorMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to generate actor key")
		}

		harvestUpdate := map[string]interface{}{
			"used": true,
		}

		resUpdate, err := originHarvest.Update(stub, harvestUpdate)
		if err != nil {
			return nil, errors.WrapError(err, "could not update harvest asset")
		}

		response := make(map[string]interface{})
		for i := 0; i < 5; i++ {
			assembledName := name + " " + fmt.Sprintf("%d", i)
			wineMap := make(map[string]interface{})
			wineMap["@assetType"] = "wine"
			wineMap["originHarvest"] = (map[string]interface{})(originHarvest)
			wineMap["wineType"] = wineType
			wineMap["productionDate"] = productionDate
			wineMap["name"] = assembledName
			wineMap["owner"] = (map[string]interface{})(actorKey)

			wineAsset, err := assets.NewAsset(wineMap)
			if err != nil {
				return nil, errors.WrapError(err, "Failed to create a new asset")
			}

			// Save the new harvest on channel
			r, err := wineAsset.PutNew(stub)
			if err != nil {
				return nil, errors.WrapError(err, "Error saving asset on blockchain")
			}

			response[assembledName] = r
		}

		response["update"] = resUpdate

		// Marshal asset back to JSON format
		wineJSON, nerr := json.Marshal(response)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return wineJSON, nil
	},
}

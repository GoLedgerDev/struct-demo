package txdefs

import (
	"encoding/json"
	"time"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Harvest on channel
// POST Method
var CreateNewHarvest = tx.Transaction{
	Tag:         "createNewHarvest",
	Label:       "Create New Harvest",
	Description: "Create a New Harvest",
	Method:      "POST",
	Callers:     []string{"$org3MSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "date",
			Label:    "Date of Harvest",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "type",
			Label:    "Grape Type",
			DataType: "grapeType",
		},
		{
			Required: true,
			Tag:      "harvestTeam",
			Label:    "Harvest Team",
			DataType: "[]string",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		date, _ := req["date"].(time.Time)
		grapeType, _ := req["type"].(float64)
		harvestTeam, _ := req["harvestTeam"].([]interface{})

		harvestMap := make(map[string]interface{})
		harvestMap["@assetType"] = "harvest"
		harvestMap["date"] = date
		harvestMap["type"] = grapeType
		harvestMap["harvestTeam"] = harvestTeam

		harvestAsset, err := assets.NewAsset(harvestMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new harvest on channel
		response, err := harvestAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		harvestJSON, nerr := json.Marshal(response)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return harvestJSON, nil
	},
}

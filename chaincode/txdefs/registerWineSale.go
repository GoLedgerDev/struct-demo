package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Harvest on channel
// POST Method
var RegisterWineSale = tx.Transaction{
	Tag:         "registerWineSale",
	Label:       "Register Wine Sale",
	Description: "Register Wine Sale",
	Method:      "POST",
	Callers:     []string{"$org1MSP"}, // Only org1 can call this transaction

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "wine",
			Label:    "Wine",
			DataType: "->wine",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		wine, _ := req["wine"].(assets.Key)

		updatedMap := map[string]interface{}{
			"sold": true,
		}

		// TODO: verify the owner of wine before registering sale

		updatedWine, err := wine.Update(stub, updatedMap)
		if err != nil {
			return nil, errors.WrapError(err, "could not update wine asset")
		}

		// Marshal asset back to JSON format
		wineJSON, nerr := json.Marshal(updatedWine)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return wineJSON, nil
	},
}

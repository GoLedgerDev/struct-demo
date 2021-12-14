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
var AcceptTransfer = tx.Transaction{
	Tag:         "acceptTransfer",
	Label:       "Accept Transfer",
	Description: "Accept Transfer",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "transferRequest",
			Label:    "Transfer Request",
			DataType: "->transferRequest",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		var err error
		transferRequest, _ := req["transferRequest"].(assets.Key)

		transferMap, err := transferRequest.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "could not get map for wine")
		}

		destinationMap := transferMap["destination"].(map[string]interface{})
		destinationKey, err := assets.NewKey(destinationMap)
		if err != nil {
			return nil, errors.WrapError(err, "could not assemble for destination")
		}

		actorMap, err := destinationKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "could not get actor map")
		}

		orgName := actorMap["orgName"].(string)

		callerName, err := stub.GetMSPID()
		if err != nil {
			return nil, errors.WrapError(err, "could not get MSP ID")
		}

		if orgName != callerName {
			return nil, errors.NewCCError("org not allowed to accept transfer", 403)
		}

		transferMap["approved"] = true
		updatedTransfer, err := transferRequest.Update(stub, transferMap)
		if err != nil {
			return nil, errors.WrapError(err, "could not update transfer asset")
		}

		wineKeyMap, _ := transferMap["wine"].(map[string]interface{})
		wineKey, err := assets.NewKey(wineKeyMap)
		if err != nil {
			return nil, errors.WrapError(err, "could not assemble key for wine")
		}

		wineMap, err := wineKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "could not get wine")
		}

		wineMap["owner"] = actorMap

		updatedWine, err := wineKey.Update(stub, wineMap)
		if err != nil {
			return nil, errors.WrapError(err, "could not update wine owner")
		}

		response := make(map[string]interface{})
		response["transfer"] = updatedTransfer
		response["wine"] = updatedWine

		// Marshal asset back to JSON format
		responseJSON, nerr := json.Marshal(response)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return responseJSON, nil
	},
}

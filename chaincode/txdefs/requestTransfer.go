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
var RequestTransfer = tx.Transaction{
	Tag:         "requestTransfer",
	Label:       "Request Transfer",
	Description: "Request Transfer",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "wine",
			Label:    "Wine",
			DataType: "->wine",
		},
		{
			Required: true,
			Tag:      "destination",
			Label:    "Destination",
			DataType: "->actor",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		var err error
		wine, _ := req["wine"].(assets.Key)
		destination, _ := req["destination"].(assets.Key)

		wineMap, err := wine.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "could not get map for wine")
		}

		wineOwner, _ := wineMap["owner"].(map[string]interface{})

		ownerKey, err := assets.NewKey(wineOwner)
		if err != nil {
			return nil, errors.WrapError(err, "could not assemble key for wine")
		}

		ownerMap, err := ownerKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "could not get map for owner")
		}

		orgName, _ := ownerMap["orgName"].(string)

		callerName, err := stub.GetMSPID()
		if err != nil {
			return nil, errors.WrapError(err, "could not get MSP ID")
		}

		if orgName != callerName {
			return nil, errors.NewCCError("org not allowed to transfer wine", 403)
		}

		t, err := stub.Stub.GetTxTimestamp()
		if err != nil {
			return nil, errors.WrapError(err, "could not get tx timestamp")
		}
		timestamp := t.AsTime()

		transferMap := make(map[string]interface{})
		transferMap["@assetType"] = "transferRequest"
		transferMap["wine"] = (map[string]interface{})(wine)
		transferMap["destination"] = (map[string]interface{})(destination)
		transferMap["timestamp"] = timestamp
		transferMap["txid"] = stub.Stub.GetTxID()

		transferAsset, err := assets.NewAsset(transferMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new transfer on channel
		response, err := transferAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		transferJSON, nerr := json.Marshal(response)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return transferJSON, nil
	},
}

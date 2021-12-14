package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var TransferRequest = assets.AssetType{
	Tag:         "transferRequest",
	Label:       "Transfer Request",
	Description: "Transfer Request",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "wine",
			Label:    "Wine",
			DataType: "->wine",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "destination",
			Label:    "Destination",
			DataType: "->actor",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "timestamp",
			Label:    "Timestamp",
			DataType: "datetime",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "txid",
			Label:    "Tx ID",
			DataType: "string",
		},
		{
			Required:     true,
			Tag:          "approved",
			Label:        "Approved?",
			DataType:     "boolean",
			DefaultValue: false,
		},
	},
}

package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Actor = assets.AssetType{
	Tag:         "actor",
	Label:       "Actor",
	Description: "Actor of the supply chain",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
		},
	},
}

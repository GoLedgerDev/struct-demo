package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Wine = assets.AssetType{
	Tag:         "wine",
	Label:       "Wine",
	Description: "Wine produced from harvest",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "originHarvest",
			Label:    "Origin harvest",
			DataType: "->harvest",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "wineType",
			Label:    "Wine type",
			DataType: "wineType",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "productionDate",
			Label:    "Production Date",
			DataType: "datetime",
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "owner",
			Label:    "Owner",
			DataType: "->actor",
		},
	},
}

package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Harvest = assets.AssetType{
	Tag:         "harvest",
	Label:       "Harvest",
	Description: "Harvest of grapes for wine production",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "date",
			Label:    "Date of Harvest",
			DataType: "datetime",
			Writers:  []string{"org3MSP"},
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "type",
			Label:    "Grape Type",
			DataType: "grapeType",
			Writers:  []string{"org3MSP"},
		},
		{
			Required: true,
			ReadOnly: true,
			Tag:      "harvestTeam",
			Label:    "Harvest Team",
			DataType: "[]string",
			Writers:  []string{"org3MSP"},
		},
		{
			Required:     true,
			Tag:          "used",
			Label:        "Used?",
			DataType:     "boolean",
			Writers:      []string{"org3MSP"},
			DefaultValue: false,
		},
	},
}

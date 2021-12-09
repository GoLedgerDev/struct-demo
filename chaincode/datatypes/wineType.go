package datatypes

import (
	"fmt"
	"math"
	"strconv"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var wineType = assets.DataType{
	AcceptedFormats: []string{"number"},
	DropDownValues: map[string]interface{}{
		"Red":   1,
		"Rose":  2,
		"White": 3,
	},
	Description: ``,

	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer", 400)
		}

		retVal := math.Trunc(dataVal)

		if dataVal != retVal {
			return "", nil, errors.NewCCError("asset property must be an integer", 400)
		}
		if retVal > 3 || retVal < 1 {
			return "", nil, errors.NewCCError("number not in range", 400)
		}
		return fmt.Sprint(retVal), retVal, nil
	},
}

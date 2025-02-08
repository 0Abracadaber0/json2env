package json

import (
	"encoding/json"
	"json2env/internal/utils/ordered_map"
)

func ReadJson(jsonData string) (*ordered_map.OrderedMap, error) {
	result := ordered_map.NewOrderedMap()
	if err := json.Unmarshal([]byte(jsonData), result); err != nil {
		return nil, err
	}

	result.Sort()

	return result, nil
}

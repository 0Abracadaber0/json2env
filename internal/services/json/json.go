package json

import (
	"encoding/json"
	"io"
	"json2env/internal/utils/ordered_map"
	"os"
)

func ReadJson(fileName string) (*ordered_map.OrderedMap, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)
	result := ordered_map.NewOrderedMap()
	if err := json.Unmarshal([]byte(bytes), result); err != nil {
		return nil, err
	}

	result.Sort()

	return result, nil
}

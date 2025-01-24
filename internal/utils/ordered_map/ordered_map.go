package ordered_map

import (
	"bytes"
	"encoding/json"
	"sort"
)

type OrderedMap struct {
	keys   []string
	values map[string]interface{}
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keys:   []string{},
		values: make(map[string]interface{}),
	}
}

func (om *OrderedMap) Set(key string, value interface{}) {
	if _, exists := om.values[key]; !exists {
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

func (om *OrderedMap) Get(key string) (interface{}, bool) {
	value, exists := om.values[key]
	return value, exists
}

func (om *OrderedMap) Keys() []string {
	return om.keys
}

func (om *OrderedMap) Sort() {
	sort.Strings(om.keys)
}

func (om *OrderedMap) UnmarshalJSON(data []byte) error {
	var tempMap map[string]interface{}

	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	om.keys = make([]string, 0, len(tempMap))
	om.values = make(map[string]interface{}, len(tempMap))

	decoder := json.NewDecoder(bytes.NewReader(data))

	if _, err := decoder.Token(); err != nil {
		return err
	}

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		key := token.(string)

		var value interface{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}

		om.keys = append(om.keys, key)
		om.values[key] = value
	}

	return nil
}

package env

import (
	"fmt"
	"json2env/internal/utils/ordered_map"
	"os"
	"strings"
)

func CreateEnvFile(filePath string, data *ordered_map.OrderedMap) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	keys := data.Keys()

	prevKey := ""
	for _, key := range keys {
		value, existed := data.Get(key)
		if !existed {
			prevKey = key
			continue
		}

		if prevKey != "" {
			prefPrevKey := strings.SplitN(prevKey, "_", 2)[0]
			prefKey := strings.SplitN(key, "_", 2)[0]
			if prefPrevKey != prefKey {
				_, err = file.WriteString(fmt.Sprintf("\n"))
			}
			if err != nil {
				return err
			}
		}

		switch v := value.(type) {
		case string:
			_, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, v))
		case float64:
			_, err = file.WriteString(fmt.Sprintf("%s=%d\n", key, int(v)))
		case bool:
			_, err = file.WriteString(fmt.Sprintf("%s=%t\n", key, v))
		default:
			_, err = file.WriteString(fmt.Sprintf("%s=%v\n", key, v))
		}
		if err != nil {
			return err
		}

		prevKey = key
	}

	return nil
}

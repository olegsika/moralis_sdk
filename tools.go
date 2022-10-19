package moralissdk

import (
	"encoding/json"
	"errors"
)

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var myMap map[string]interface{}
	data, _ := json.Marshal(s)
	if err := json.Unmarshal(data, &myMap); err != nil {
		return nil, err
	}

	return myMap, nil
}

func InListString(entry string, list []string) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}

func IsErrorEqual(err, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}

	return err.Error() == target.Error()
}

package moralissdk

import (
	"encoding/json"
	"errors"
)

// StructToMap function convert Struct to Map
func StructToMap(s interface{}) (map[string]interface{}, error) {
	var myMap map[string]interface{}
	data, _ := json.Marshal(s)
	if err := json.Unmarshal(data, &myMap); err != nil {
		return nil, err
	}

	return myMap, nil
}

// InListString function check if the string exists on slice of strings
func InListString(entry string, list []string) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}

// IsErrorEqual function check if the errors equal
func IsErrorEqual(err, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}

	return err.Error() == target.Error()
}

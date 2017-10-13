package main

import "errors"

func convertInterfaceArrayToStringArray(arrayOfInterfaces []interface{}) ([]string, error) {
	ret := []string{}
	for _, v := range arrayOfInterfaces {
		if s, ok := v.(string); ok {
			ret = append(ret, s)
		} else {
			return nil, errors.New("Failed to convert.")
		}
	}
	return ret, nil
}

var convertMapOfInterfacesToMapOfStrings = func(data map[string]interface{}) (map[string]string, error) {
	if data == nil {
		return nil, errors.New("Cannot convert nil.")
	}
	ret := map[string]string{}
	for k, v := range data {
		if v == nil {
			return nil, errors.New("Data contains nil.")
		}
		ret[k] = v.(string)
	}
	return ret, nil
}

var convertMapOfStringsToMapOfInterfaces = func(data map[string]string) (map[string]interface{}, error) {
	if data == nil {
		return nil, errors.New("Cannot convert nil.")
	}
	ret := map[string]interface{}{}
	for k, v := range data {
		ret[k] = v
	}
	return ret, nil
}

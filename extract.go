package extract

import (
	"fmt"
	"errors"
)

type Cast interface {
	string | float64 | map[string]interface{} | []interface{}
}

// Extract value by default
func D[T Cast](res any, n ...any) (T, error){
	var errRes T
	for _, v := range n {
		switch res.(type) {
		case map[string]interface{}:
			v2, ok := v.(string)
			if ok {
				res2 := res.(map[string]interface{})
				var exist bool
				res, exist = res2[v2]
				if !exist {
					return errRes, errors.New(fmt.Sprintf(`Error: Index "%v" is wrong`, v2))
				}
			} else {
				return errRes, errors.New(fmt.Sprintf(`Error: The parameter "%v" (%T) should be string`, v, v))
			}
		case []interface{}:
			v2, ok := v.(int)
			if ok {
				res2 := res.([]interface{})
				if len(res2) - 1 < v2 {
					return errRes, errors.New("Error: Index out of range")
				}
				res = res2[v2]
			} else {
				return errRes, errors.New(fmt.Sprintf(`Error: The parameter "%v" (%T) should be int`, v, v))
			}
		}
	}
	response, ok := res.(T)
	if ok {
		return response, nil
	}
	return errRes, errors.New(fmt.Sprintf("Error: Value cannot be cast with type (%T)", response))
}

func HelpK[T Cast](res any, key string) (bool, T, error) {
	var errRes T
	switch res.(type) {
	case map[string]interface{}:
		res2 := res.(map[string]interface{})
		for v := range res2 {
			if v == key {
				response, ok := res2[key].(T)
				if ok {
					return true, response, nil
				}
				return true, errRes, errors.New(fmt.Sprintf(`Value cannot be cast with type (%T)`, response)) // not working error message errors.New(fmt.Sprintf("Error: Value cannot be cast with type (%T)", key))
			}
			exist, response, err := HelpK[T](res2[v], key)
			if exist {
				return true, response, err
			}
		}
	case []interface{}:
		res2 := res.([]interface{})
		for v := range res2 {
			exist, response, err := HelpK[T](res2[v], key)
			if exist {
				return true, response, err
			}
		}
	}
	return false, errRes, errors.New(fmt.Sprintf(`Key "%v" not found`, key)) // not working error message errors.New(fmt.Sprintf(`Error: Key "%v" not found`, key))
}
 
// Extract value with one key
func K[T Cast](res any, key string) (T, error) {
	var errRes T
	exist, response, err := HelpK[T](res, key)
	if exist && err == nil {
		return response, nil
	}
	return errRes, err
}

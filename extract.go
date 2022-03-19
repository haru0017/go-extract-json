package extract

import (
	"fmt"
	"errors"
)

type cast interface {
	string | float64 | map[string]interface{} | []interface{}
}

// Extract value by default
func D[T cast](res any, n ...any) (T, error){
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
		default:
			return errRes, errors.New(fmt.Sprintf(`"%#v" is not correctly decoded from JSON or "%v" is wrong`, res, v))
		}
	}
	response, ok := res.(T)
	if ok {
		return response, nil
	}
	return errRes, errors.New(fmt.Sprintf("Error: Value cannot be cast with type (%T)", response))
}

// Search for key
func helpK[T cast](res any, key string) (bool, T, error) {
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
				return true, errRes, errors.New(fmt.Sprintf(`Value cannot be cast with type (%T)`, response))
			}
			exist, response, err := helpK[T](res2[v], key)
			if exist {
				return true, response, err
			}
		}
	case []interface{}:
		res2 := res.([]interface{})
		for v := range res2 {
			exist, response, err := helpK[T](res2[v], key)
			if exist {
				return true, response, err
			}
		}
	case string:
	case float64:
	default:
	  	return true, errRes, errors.New(fmt.Sprintf(`"%#v" is not correctly decoded from JSON`, res))
	}
	return false, errRes, errors.New(fmt.Sprintf(`Key "%v" not found`, key))
}
 
// Extract value with one key
func K[T cast](res any, key string) (T, error) {
	var errRes T
	exist, response, err := helpK[T](res, key)
	if exist && err == nil {
		return response, nil
	}
	return errRes, err
}

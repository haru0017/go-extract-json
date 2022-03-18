package extract

import (
	"testing"
	"encoding/json"
)

func AssertValue[T comparable](t *testing.T, x T, y T) {
	t.Helper()
	if x != y {
		t.Fatalf("%v is not %v", x, y)
	}
}

func ReturnErr(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Failed to output error")
	}
}

const jsonString string = `
{
	"package": {
		"name": "go-extract-json",
		"author": "haru",
		"star": 999
	},
	"language": {
		"go": {
			"version": 1.18
		}
	},
	"functions": [
	  {
		"name": "D",
		"function": "Extract value from JSON by default" 
	  },
	  {
		"name": "K",
		"function": "Extract from JSON with one key"
	  }
	]
}
`

var res interface{} 

func TestD(t *testing.T) {
	json.Unmarshal([]byte(jsonString), &res)

	resStr, _ := D[string](res, "package", "name")
	AssertValue(t, resStr, "go-extract-json")

	resFloat, _ := D[float64](res, "language", "go", "version")
	AssertValue(t, resFloat, 1.18)
	
	resSlice, _ := D[[]interface{}](res, "functions")
	AssertValue(t, resSlice[0].(map[string]interface{})["name"].(string), "D")

	resMap, _ := D[map[string]interface{}](res, "package")
	AssertValue(t, resMap["author"].(string), "haru")

	_, err := D[map[string]interface{}](res, [3]int{1, 2, 3})
	ReturnErr(t, err)

	_, err = D[[]interface{}](res, "functions", "name")
	ReturnErr(t, err)

	_, err = D[float64](res, "package", "author")
	ReturnErr(t, err)

	_, err = D[string](res, "package", "star")
	ReturnErr(t, err)

	_, err = D[string](res, "functions", 2, "function")
	ReturnErr(t, err)
}

 func TestK(t *testing.T) {
	json.Unmarshal([]byte(jsonString), &res)

	resStr, _ := K[string](res, "author")
	AssertValue(t, resStr, "haru")

	_, err := K[string](res, "Rust")
	ReturnErr(t, err)

	_, err = K[string](res, "star")
	ReturnErr(t, err)
}

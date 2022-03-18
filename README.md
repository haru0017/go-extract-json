# go-extract-json

Easily extract any value from JSON decoded without defining a structure

## Usage
### `extract.D[T](obj any, index ...any) (T, error)`  
Extract value by specifying index and type of data you want to get
```Go
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
json.Unmarshal([]byte(jsonString), &res)

resStr, _ := extract.D[string](res, "package", "name") // resStr: "go-extract-json"
```

### `extract.K[T](obj any, key string) (T, error)`  
Extract value with one key
```Go
resFloat, _ := extract.K[float64](res, "version") // resFloat: 1.18
```

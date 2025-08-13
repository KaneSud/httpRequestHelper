# EveryDayTool
This is project with some helpful... ahem, functions, I guess

### What do we have here:
#### http/requests
Imagine, that u've got this request:

http://localhost:3333/?float=3.4&bool=true&uint=23&int=-1&string=test&complex=6

And u need to parse it, and it's not json, so no json.Unmarshal, and thats boring. 
So I decided to code function ParseQueryStruct that do this for my future me.
Example of usage:
```go
var outputSingle TestFullStruct
err = ParseQueryStruct(&outputSingle, request)
```
It takes pinter to structure of request (*TestFullStruct in example) and *http.Request.
It supports types:`
1) bool
2) string
3) float64
4) int64
5) uint64
6) slices of all types above

Function look at tag json and if it's "fill", then the field will be filled.
For example:
```go
type TestFullStruct struct {
    Int     int64      `json:"int"`
}

```

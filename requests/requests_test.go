package requests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

type TestFullStruct struct {
	Float   float64    `json:"float"`
	Bool    bool       `json:"bool"`
	Uint    uint64     `json:"uint" `
	Int     int64      `json:"int" `
	String  string     `json:"string" `
	Complex complex128 `json:"complex" `

	StringSlice  []string     `json:"strings" `
	UintSlice    []uint64     `json:"uints"`
	ComplexSlice []complex128 `json:"complexes"`
	BoolSlice    []bool       `json:"booleans" `
	IntSlice     []int64      `json:"ints" `
	FloatSlice   []float64    `json:"floats" `

	FieldSkip         []string `json:"namesSkip"`
	FieldNil          []string `json:"nameNil" `
	FieldEmptyTag     []string `json:"namesEmpty" `
	FieldFullEmptyTag []string
}

type TestPointers struct {
	Bool *bool `json:"bool"`
}

type TestBench struct {
	Float   float64    `json:"float"`
	Bool    bool       `json:"bool"`
	Uint    uint64     `json:"uint"`
	Int     int64      `json:"int" `
	String  string     `json:"string" `
	Complex complex128 `json:"complex" `
}

func BenchmarkParseQueryStructSimple(b *testing.B) {
	request, err := http.NewRequest("GET", "http://localhost:3333/?int=-1", nil)
	assert.NoError(b, err, "Error during request creation")
	var outputSingle TestBench
	for i := 0; i < b.N; i++ {
		err = ParseQueryStruct(&outputSingle, request)
	}
}

func BenchmarkParseQueryStructSlice(b *testing.B) {
	request, err := http.NewRequest("GET", "http://localhost:3333/?strings=test1,test2&uints=1,2,3&complexes=6,7&booleans=true,false&ints=-1,23,0&floats=-23.1,255.1", nil)
	assert.NoError(b, err, "Error during request creation")
	var outputSingle TestFullStruct
	for i := 0; i < b.N; i++ {
		err = ParseQueryStruct(&outputSingle, request)
	}
}

func BenchmarkParseQueryStructSingle(b *testing.B) {
	request, err := http.NewRequest("GET", "http://localhost:3333/?float=3.4&bool=true&uint=23&int=-1&string=test&complex=6&lolkek=12", nil)
	assert.NoError(b, err, "Error during request creation")
	var outputSingle TestBench
	for i := 0; i < b.N; i++ {
		err = ParseQueryStruct(&outputSingle, request)
	}
}

func TestParseQueryStructSingle(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:3333/?float=3.4&bool=true&uint=23&int=-1&string=test&complex=6", nil)
	assert.NoError(t, err, "Error during request creation")
	var outputSingle TestFullStruct
	err = ParseQueryStruct(&outputSingle, request)
	assert.NoError(t, err, "Error during parsing")
	assert.Equal(t, 3.4, outputSingle.Float, "Error during float64 parse")
	assert.Equal(t, true, outputSingle.Bool, "Error during bool parse")
	assert.Equal(t, uint64(23), outputSingle.Uint, "Error during uint64 parse")
	assert.Equal(t, int64(-1), outputSingle.Int, "Error during int64 parse")
	assert.Equal(t, "test", outputSingle.String, "Error during string parse")
	assert.Equal(t, complex128((6 + 0i)), outputSingle.Complex, "Error during complex128 parse")
	assert.Nil(t, outputSingle.FieldSkip)
	assert.Nil(t, outputSingle.FieldEmptyTag)
	assert.Nil(t, outputSingle.FieldFullEmptyTag)
	assert.Nil(t, outputSingle.FieldNil)
}

func TestParseQueryStructSlices(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:3333/?strings=test1,test2&uints=1,2,3&complexes=6,7&booleans=true,false&ints=-1,23,0&floats=-23.1,255.1", nil)
	assert.NoError(t, err, "Error during request creation")
	var outputSlice TestFullStruct
	err = ParseQueryStruct(&outputSlice, request)
	assert.NoError(t, err, "Error during parsing")
	assert.True(t, reflect.DeepEqual(outputSlice.StringSlice, []string{"test1", "test2"}), "Error string slice parsed incorrect")
	assert.True(t, reflect.DeepEqual(outputSlice.UintSlice, []uint64{1, 2, 3}), "Error uint64 slice parsed incorrect")
	assert.True(t, reflect.DeepEqual(outputSlice.ComplexSlice, []complex128{(6 + 0i), (7 + 0i)}), "Error complex128 slice parsed incorrect")
	assert.True(t, reflect.DeepEqual(outputSlice.BoolSlice, []bool{true, false}), "Error bool slice parsed incorrect")
	assert.True(t, reflect.DeepEqual(outputSlice.IntSlice, []int64{-1, 23, 0}), "Error int64 slice parsed incorrect")
	assert.True(t, reflect.DeepEqual(outputSlice.FloatSlice, []float64{-23.1, 255.1}), "Error float64 slice parsed incorrect")
	assert.Nil(t, outputSlice.FieldSkip)
	assert.Nil(t, outputSlice.FieldEmptyTag)
	assert.Nil(t, outputSlice.FieldFullEmptyTag)
	assert.Nil(t, outputSlice.FieldNil)
}

func TestParseQueryStructNil(t *testing.T) {
	var outputSlice TestFullStruct
	request, err := http.NewRequest("GET", "http://localhost:3333/?strings=test1,test2&uints=1,2,3&complexes=6,7&booleans=true,false&ints=-1,23,0&floats=-23.1,255.1", nil)
	assert.NoError(t, err, "Error during request creation")
	err = ParseQueryStruct(&outputSlice, nil)
	assert.Error(t, err)
	err = ParseQueryStruct(nil, nil)
	assert.Error(t, err)
	err = ParseQueryStruct(nil, request)
	assert.Error(t, err)
	err = ParseQueryStruct(interface{}(nil), request)
	assert.Error(t, err)
}

func TestParseQueryStructPointers(t *testing.T) {
	var outputSlice TestPointers
	request, err := http.NewRequest("GET", "http://localhost:3333/?strings=test1&bool=true", nil)
	assert.NoError(t, err, "Error during request creation")
	err = ParseQueryStruct(&outputSlice, request)
	fmt.Println(outputSlice)
	assert.NoError(t, err)
}

package null

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBool(t *testing.T) {
	t.Parallel()
	t.Run("Success", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nullBool.Valid)
		assert.Equal(t, true, nullBool.Bool)
	})

	t.Run("False on nil", func(t *testing.T) {
		nullBool, err := NewBool(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.False(t, nullBool.Valid)
		assert.Equal(t, false, nullBool.Bool)
	})
}

func TestBool_Value(t *testing.T) {
	t.Parallel()
	t.Run("Return bool va", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, err := nullBool.Value()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, true, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var nullBool Bool
		value, _ := nullBool.Value()

		assert.Nil(t, value)
	})
}

func TestBool_Scan(t *testing.T) {
	nb1, _ := NewBool(true)
	nb2, _ := NewBool(false)
	nb3, _ := NewBool(nil)
	cases := TestCases{
		"ints": {
			{in: 1, va: true, iv: true, ie: false},
			{in: int8(1), va: true, iv: true, ie: false},
			{in: int16(1), va: true, iv: true, ie: false},
			{in: int32(1), va: true, iv: true, ie: false},
			{in: int64(1), va: true, iv: true, ie: false},

			{in: 0, va: false, iv: true, ie: false},
			{in: int8(0), va: false, iv: true, ie: false},
			{in: int16(0), va: false, iv: true, ie: false},
			{in: int32(0), va: false, iv: true, ie: false},
			{in: int64(0), va: false, iv: true, ie: false},

			{in: 5, va: false, iv: false, ie: false},
			{in: -5, va: false, iv: false, ie: false},
		},
		"strings": {
			{in: "1", va: true, iv: true, ie: false},
			{in: "t", va: true, iv: true, ie: false},
			{in: "T", va: true, iv: true, ie: false},
			{in: "true", va: true, iv: true, ie: false},
			{in: "TRUE", va: true, iv: true, ie: false},
			{in: "True", va: true, iv: true, ie: false},
			{in: "y", va: true, iv: true, ie: false},
			{in: "Y", va: true, iv: true, ie: false},
			{in: "YES", va: true, iv: true, ie: false},
			{in: "Yes", va: true, iv: true, ie: false},

			{in: "0", va: false, iv: true, ie: false},
			{in: "f", va: false, iv: true, ie: false},
			{in: "F", va: false, iv: true, ie: false},
			{in: "false", va: false, iv: true, ie: false},
			{in: "False", va: false, iv: true, ie: false},
			{in: "FALSE", va: false, iv: true, ie: false},
			{in: "na", va: false, iv: true, ie: false},
			{in: "N", va: false, iv: true, ie: false},
			{in: "NO", va: false, iv: true, ie: false},
			{in: "No", va: false, iv: true, ie: false},
			{in: "some string", va: false, iv: false, ie: false},
		},

		"booleans": {
			{in: true, va: true, iv: true, ie: false},
			{in: false, va: false, iv: true, ie: false},
			{in: nb1, va: true, iv: true, ie: false},
			{in: nb2, va: false, iv: true, ie: false},
			{in: nb3, va: false, iv: false, ie: false},
		},

		"byte slice": {
			{na: "bytes for true", in: makeBytes(true), va: true, iv: true, ie: false},
			{na: "bytes for false", in: makeBytes(false), va: false, iv: true, ie: false},
			{na: "bytes for nil", in: makeBytes(nil), va: false, iv: false, ie: false},
		},
		"nil": {
			{in: nil, va: false, iv: false, ie: false},
		},

		"errors": {},
	}
	checkCases(cases, t, Bool{})
}

func BenchmarkBool_Scan(b *testing.B) {
	var nb Bool
	for i := 0; i < b.N; i++ {
		err := nb.Scan(i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		b, _ := json.Marshal(true)
		jb, err := nullBool.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nb, err := NewBool(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, _ := nb.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func BenchmarkBool_MarshalJSON(b *testing.B) {
	nb, _ := NewBool("true")
	for i := 0; i < b.N; i++ {
		_, err := nb.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	cases := []map[string]interface{}{
		{in: []byte("true"), va: true, iv: true},
		{in: []byte("1"), va: true, iv: true},
		{in: []byte("t"), va: true, iv: true},
		{in: []byte("T"), va: true, iv: true},
		{in: []byte("TRUE"), va: true, iv: true},
		{in: []byte("True"), va: true, iv: true},

		{in: []byte("0"), va: false, iv: true},
		{in: []byte("f"), va: false, iv: true},
		{in: []byte("F"), va: false, iv: true},
		{in: []byte("false"), va: false, iv: true},
		{in: []byte("False"), va: false, iv: true},
		{in: []byte("FALSE"), va: false, iv: true},

		{in: []byte("not bool"), va: false, iv: false},
		{in: []byte("null"), va: false, iv: false},
	}

	for _, testCase := range cases {
		var nullBool Bool
		err := nullBool.UnmarshalJSON(testCase[in].([]byte))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, testCase[va], nullBool.Bool, "va param for intput %+v: %+v", testCase[in], testCase[va])
		assert.Equal(t, testCase[iv], nullBool.Valid, "iv param for intput %+v: %+v", testCase[in], testCase[iv])
	}
}

func BenchmarkBool_UnmarshalJSON(b *testing.B) {
	bs := "\"true\""
	bytes := []byte(bs)
	var nb Bool
	for i := 0; i < b.N; i++ {
		err := nb.UnmarshalJSON(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

}

package null

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"github.com/smgladkovskiy/structs"
	"strconv"
	"strings"
)

// Int64 Реализация Int64
type Int64 struct {
	Int64 int64
	Valid bool
}

func NewInt64(v interface{}) (*Int64, error) {
	var ni Int64
	err := ni.Scan(v)
	return &ni, err
}

func (ni *Int64) Scan(value interface{}) error {
	if value == nil {
		ni.Int64, ni.Valid = 0, false
		return nil
	}
	var err error

	ni.Valid = false
	switch v := value.(type) {
	case string:
		ni.Int64, err = strconv.ParseInt(v, 10, 64)
		if err == nil {
			ni.Valid = true
		}
		return err
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		if v == 0 {
			ni.Int64 = 0
			return nil
		}
		i, ok := v.(int)
		if ok {
			ni.Int64, ni.Valid = int64(i), true
			return nil
		}
		ui, ok := v.(uint)
		if ok {
			ni.Int64, ni.Valid = int64(ui), true
			return nil
		}
		i8, ok := v.(int8)
		if ok {
			ni.Int64, ni.Valid = int64(i8), true
			return nil
		}
		ui8, ok := v.(uint8)
		if ok {
			ni.Int64, ni.Valid = int64(ui8), true
			return nil
		}
		i16, ok := v.(int16)
		if ok {
			ni.Int64, ni.Valid = int64(i16), true
			return nil
		}
		ui16, ok := v.(uint16)
		if ok {
			ni.Int64, ni.Valid = int64(ui16), true
			return nil
		}
		i32, ok := v.(int32)
		if ok {
			ni.Int64, ni.Valid = int64(i32), true
			return nil
		}
		ui32, ok := v.(uint32)
		if ok {
			ni.Int64, ni.Valid = int64(ui32), true
			return nil
		}
		i64, ok := v.(int64)
		if ok {
			ni.Int64, ni.Valid = i64, true
			return nil
		}
		ui64, ok := v.(uint64)
		if ok {
			ni.Int64, ni.Valid = int64(ui64), true
			return nil
		}
	case []byte:
		i := int64(binary.BigEndian.Uint64(v))
		ni.Int64, ni.Valid = i, i > 0
		return nil
	case Int64:
		ni.Int64, ni.Valid = v.Int64, v.Valid
		return nil
	case *Int64:
		ni.Int64, ni.Valid = v.Int64, v.Valid
		return nil
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: ni}
}

// va implements the driver Valuer interface.
func (ni Int64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int64, nil
}

// MarshalJSON correctly serializes a Int64 to JSON
func (ni Int64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return structs.NullString, nil
}

func (ni *Int64) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		ni.Int64 = 0
		return
	}

	ni.Int64, err = strconv.ParseInt(s, 10, 64)
	if err == nil {
		ni.Valid = true
	}
	return
}

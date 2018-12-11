package null

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/smgladkovskiy/structs"
	"strings"
	"time"
)

type Date struct {
	Time  time.Time
	Valid bool // iv is true if Time is not NULL
}

// NewDate Создание Date переменной
func NewDate(v interface{}) (*Date, error) {
	var nt Date
	err := nt.Scan(v)
	return &nt, err
}

// Scan implements the Scanner interface for Date
func (nd *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		t, err := time.Parse(structs.DateFormat(), v)
		if err != nil {
			*nd = Date{Time: time.Time{}, Valid: false}
			return err
		}
		nd.Time, nd.Valid = t, true
		return nil
	case time.Time:
		if v.IsZero() {
			*nd = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		nd.Time, nd.Valid = v, true

		return nil
	case *time.Time:
		if v.IsZero() {
			*nd = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		nd.Time, nd.Valid = *v, true

		return nil
	case Date:
		*nd = v
		return nil
	case *Date:
		if v.Time.IsZero() {
			return nil
		}

		nd.Time, nd.Valid = v.Time, v.Valid

		return nil
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: nd}
}

// va implements the driver Valuer interface.
func (nd Date) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Time, nil
}

func (nd *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nd.Time = time.Time{}
		return
	}
	nd.Time, err = time.Parse(structs.DateFormat(), s)
	if err == nil {
		nd.Valid = true
	}
	return
}

func (nd Date) MarshalJSON() ([]byte, error) {
	if nd.Valid {
		return json.Marshal(nd.Time.Format(structs.DateFormat()))
	}

	return structs.NullString, nil
}

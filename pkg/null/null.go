package null

import (
	"database/sql/driver"
	"errors"
)

type String string

func (s *String) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return errors.New("column is not a string")
	}

	*s = String(strVal)

	return nil
}

func (s String) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}

	return string(s), nil
}

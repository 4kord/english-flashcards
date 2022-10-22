package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable Bool that supports JSON
type NullBool sql.NullBool

func (nb NullBool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return json.Marshal(nil)
}

func (nb *NullBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		nb.Valid = true
		nb.Bool = *b
	} else {
		nb.Valid = false
	}
	return nil
}

// Nullable Float64 that supports JSON
type NullFloat64 sql.NullFloat64

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	}
	return json.Marshal(nil)
}

func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	var f *float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	if f != nil {
		nf.Valid = true
		nf.Float64 = *f
	} else {
		nf.Valid = false
	}
	return nil
}

// Nullable Int32 that supports JSON
type NullInt32 sql.NullInt32

func (ni NullInt32) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int32)
	}
	return json.Marshal(nil)
}

func (ni *NullInt32) UnmarshalJSON(data []byte) error {
	var i *int32
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Valid = true
		ni.Int32 = *i
	} else {
		ni.Valid = false
	}
	return nil
}

// Nullable String that supports JSON
type NullString sql.NullString

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}

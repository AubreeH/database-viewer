package sqlTypes

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

type NullBool struct {
	sql.NullBool
}

type NullByte struct {
	sql.NullByte
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullInt16 struct {
	sql.NullInt16
}

type NullInt32 struct {
	sql.NullInt32
}

type NullInt64 struct {
	sql.NullInt64
}

type NullTime struct {
	sql.NullTime
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}
	return []byte(`null`), nil
}

func (b NullByte) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Byte)
	}
	return []byte(`null`), nil
}

func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	}
	return []byte(`null`), nil
}

func (i NullInt16) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int16)
	}
	return []byte(`null`), nil
}

func (i NullInt32) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int32)
	}
	return []byte(`null`), nil
}

func (i NullInt64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	}
	return []byte(`null`), nil
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return []byte(`null`), nil
}

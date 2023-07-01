package types

import (
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (v *UUID) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	if s == "" {
		return nil
	}

	n, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid integer string '%s' => %w", s, err)
	}

	*v = UUID(n)

	return nil

}

func (v *UUID) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	n, err := uuid.Parse(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid integer string '%s' => %w", attr.Value, err)
	}

	*v = UUID(n)

	return nil
}

func (v *UUID) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", *v)), nil
}

func (v *UUID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if v == nil {
		return xml.Attr{Name: name, Value: ""}, nil
	}

	return xml.Attr{Name: name, Value: fmt.Sprintf("%s", *v)}, nil
}

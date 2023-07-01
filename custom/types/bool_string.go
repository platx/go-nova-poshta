package types

import (
	"encoding/xml"
	"fmt"
)

type BoolString bool

func (v *BoolString) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	switch s {
	case "true", "1":
		*v = true
	case "false", "0":
		*v = false
	default:
		return fmt.Errorf("invalid boolean string: %s", s)
	}

	return nil
}

func (v *BoolString) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "true", "1":
		*v = true
	case "false", "0":
		*v = false
	default:
		return fmt.Errorf("invalid boolean string: %s", attr.Value)
	}

	return nil
}

func (v *BoolString) MarshalJSON() ([]byte, error) {
	if *v {
		return []byte("\"1\""), nil
	}

	return []byte("\"0\""), nil
}

func (v *BoolString) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if *v {
		return xml.Attr{Name: name, Value: "1"}, nil
	}

	return xml.Attr{Name: name, Value: "0"}, nil
}

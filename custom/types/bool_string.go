package types

import (
	"encoding/xml"
	"fmt"
)

type BoolString bool

func (v BoolString) String() string {
	if v {
		return "1"
	}

	return "0"
}

func (v *BoolString) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	switch s {
	case "1":
		*v = true
	case "0":
		*v = false
	default:
		return fmt.Errorf("invalid boolean string '%s'", s)
	}

	return nil
}

func (v *BoolString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	switch s {
	case "1":
		*v = true
	case "0":
		*v = false
	default:
		return fmt.Errorf("invalid boolean string '%s'", s)
	}

	return nil
}

func (v BoolString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", v.String())), nil
}

func (v BoolString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(v.String(), start)
}

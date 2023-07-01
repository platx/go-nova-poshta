package types

import (
	"encoding/xml"
	"fmt"
)

type BoolInt bool

func (v *BoolInt) UnmarshalJSON(data []byte) error {
	s := string(data)

	switch s {
	case "1":
		*v = true
	case "0":
		*v = false
	default:
		return fmt.Errorf("invalid boolean integer '%s'", s)
	}

	return nil
}

func (v *BoolInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		return fmt.Errorf("invalid boolean integer '%s'", s)
	}

	return nil
}

func (v BoolInt) MarshalJSON() ([]byte, error) {
	if v {
		return []byte("1"), nil
	}

	return []byte("0"), nil
}

func (v BoolInt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if v {
		return e.EncodeElement("1", start)
	}

	return e.EncodeElement("0", start)
}

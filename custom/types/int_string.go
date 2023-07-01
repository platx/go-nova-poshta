package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type IntString int

func (v *IntString) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	if s == "" {
		*v = IntString(0)

		return nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid integer string '%s' => %w", s, err)
	}

	*v = IntString(n)

	return nil

}

func (v *IntString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*v = IntString(0)

		return nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid integer string '%s' => %w", s, err)
	}

	*v = IntString(n)

	return nil
}

func (v IntString) MarshalJSON() ([]byte, error) {
	if int(v) == 0 {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%d"`, v)), nil
}

func (v IntString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if int(v) == 0 {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(fmt.Sprintf("%d", v), start)
}

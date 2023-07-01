package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type FloatString float64

func (v *FloatString) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	if s == "" {
		*v = FloatString(0.0)

		return nil
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid float string '%s' => %w", s, err)
	}

	*v = FloatString(n)

	return nil

}

func (v *FloatString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*v = FloatString(0.0)

		return nil
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid float string '%s' => %w", s, err)
	}

	*v = FloatString(n)

	return nil
}

func (v FloatString) MarshalJSON() ([]byte, error) {
	if float64(v) == 0 {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%v"`, v)), nil
}

func (v FloatString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if float64(v) == 0 {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(fmt.Sprintf("%v", v), start)
}

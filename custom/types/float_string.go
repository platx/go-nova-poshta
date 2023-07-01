package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type FloatString float64

func (v *FloatString) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid float string '%s' => %w", s, err)
	}

	*v = FloatString(n)

	return nil

}

func (v *FloatString) UnmarshalXMLAttr(attr xml.Attr) error {
	n, err := strconv.ParseFloat(attr.Value, 64)
	if err != nil {
		return fmt.Errorf("invalid float string '%s' => %w", attr.Value, err)
	}

	*v = FloatString(n)

	return nil
}

func (v *FloatString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", *v)), nil
}

func (v *FloatString) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: fmt.Sprintf("%v", *v)}, nil
}

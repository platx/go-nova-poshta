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

func (v *IntString) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		*v = IntString(0)

		return nil
	}

	n, err := strconv.Atoi(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid integer string '%s' => %w", attr.Value, err)
	}

	*v = IntString(n)

	return nil
}

func (v *IntString) MarshalJSON() ([]byte, error) {
	if int(*v) == 0 {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%d\"", *v)), nil
}

func (v *IntString) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if int(*v) == 0 {
		return xml.Attr{Name: name, Value: ""}, nil
	}

	return xml.Attr{Name: name, Value: fmt.Sprintf("%d", *v)}, nil
}

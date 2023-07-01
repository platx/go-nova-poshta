package types

import (
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
)

type UUID uuid.UUID

var EmptyUUID = UUID(uuid.Nil)

func MustParseUUID(s string) UUID {
	return UUID(uuid.MustParse(s))
}

func (v *UUID) String() string {
	return uuid.UUID(*v).String()
}

func (v *UUID) UnmarshalJSON(data []byte) error {
	s := trimQuotes(string(data))

	if s == "" {
		*v = UUID(uuid.Nil)
		return nil
	}

	n, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid uuid string '%s' => %w", s, err)
	}

	*v = UUID(n)

	return nil

}

func (v *UUID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*v = UUID(uuid.Nil)

		return nil
	}

	n, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid uuid string '%s' => %w", s, err)
	}

	*v = UUID(n)

	return nil
}

func (v UUID) MarshalJSON() ([]byte, error) {
	cv := uuid.UUID(v)

	if cv == uuid.Nil {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", cv.String())), nil
}

func (v UUID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	cv := uuid.UUID(v)

	if cv == uuid.Nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(cv.String(), start)
}

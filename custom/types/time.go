package types

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

const (
	CustomDateFormat          = "02.01.2006"
	CustomDateTimeFormat      = "2006.01.02 15:04:05"
	ReverseDateTimeFormat     = "02-01-2006 15:04:05"
	ReverseDotDateTimeFormat  = "02.01.2006 15:04:05"
	DateHourMinuteFormat      = "2006.01.02 15:04"
	SlashDateHourMinuteFormat = "06/01/02 15:04"
)

func MustParseTime(format string, value string) time.Time {
	v, err := time.Parse(format, value)

	if err != nil {
		panic(err)
	}

	return v
}

type Date struct {
	time.Time
}

func MustParseDate(value string) Date {
	return Date{MustParseTime(time.DateOnly, value)}
}

func (t *Date) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = Date{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = Date{parsed}

	return nil
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = Date{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = Date{parsed}

	return nil
}

func (t Date) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.DateOnly))), nil
}

func (t Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(time.DateOnly), start)
}

type DateTime struct {
	time.Time
}

func MustParseDateTime(value string) DateTime {
	return DateTime{MustParseTime(time.DateTime, value)}
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = DateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(time.DateTime, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = DateTime{parsed}

	return nil
}

func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = DateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(time.DateTime, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = DateTime{parsed}

	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.DateTime))), nil
}

func (t DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(time.DateTime), start)
}

type CustomDate struct {
	time.Time
}

func MustParseCustomDate(value string) CustomDate {
	return CustomDate{MustParseTime(CustomDateFormat, value)}
}

func (t *CustomDate) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = CustomDate{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(CustomDateFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = CustomDate{parsed}

	return nil
}

func (t *CustomDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = CustomDate{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(CustomDateFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = CustomDate{parsed}

	return nil
}

func (t CustomDate) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(CustomDateFormat))), nil
}

func (t CustomDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(CustomDateFormat), start)
}

type CustomDateTime struct {
	time.Time
}

func MustParseCustomDateTime(value string) CustomDateTime {
	return CustomDateTime{MustParseTime(CustomDateTimeFormat, value)}
}

func (t *CustomDateTime) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = CustomDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(CustomDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = CustomDateTime{parsed}

	return nil
}

func (t *CustomDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = CustomDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(CustomDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = CustomDateTime{parsed}

	return nil
}

func (t CustomDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(CustomDateTimeFormat))), nil
}

func (t CustomDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(CustomDateTimeFormat), start)
}

type ReverseDateTime struct {
	time.Time
}

func MustParseReverseDateTime(value string) ReverseDateTime {
	return ReverseDateTime{MustParseTime(ReverseDateTimeFormat, value)}
}

func (t *ReverseDateTime) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = ReverseDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(ReverseDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = ReverseDateTime{parsed}

	return nil
}

func (t *ReverseDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = ReverseDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(ReverseDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = ReverseDateTime{parsed}

	return nil
}

func (t ReverseDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(ReverseDateTimeFormat))), nil
}

func (t ReverseDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(ReverseDateTimeFormat), start)
}

type ReverseDotDateTime struct {
	time.Time
}

func MustParseReverseDotDateTime(value string) ReverseDotDateTime {
	return ReverseDotDateTime{MustParseTime(ReverseDotDateTimeFormat, value)}
}

func (t *ReverseDotDateTime) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = ReverseDotDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(ReverseDotDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = ReverseDotDateTime{parsed}

	return nil
}

func (t *ReverseDotDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = ReverseDotDateTime{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(ReverseDotDateTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = ReverseDotDateTime{parsed}

	return nil
}

func (t ReverseDotDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(ReverseDotDateTimeFormat))), nil
}

func (t ReverseDotDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(ReverseDotDateTimeFormat), start)
}

type DateHourMinute struct {
	time.Time
}

func MustParseDateHourMinute(value string) DateHourMinute {
	return DateHourMinute{MustParseTime(DateHourMinuteFormat, value)}
}

func (t *DateHourMinute) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = DateHourMinute{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(DateHourMinuteFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = DateHourMinute{parsed}

	return nil
}

func (t *DateHourMinute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = DateHourMinute{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(DateHourMinuteFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = DateHourMinute{parsed}

	return nil
}

func (t DateHourMinute) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(DateHourMinuteFormat))), nil
}

func (t DateHourMinute) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(DateHourMinuteFormat), start)
}

type SlashDateHourMinute struct {
	time.Time
}

func MustParseSlashDateHourMinute(value string) SlashDateHourMinute {
	return SlashDateHourMinute{MustParseTime(SlashDateHourMinuteFormat, value)}
}

func (t *SlashDateHourMinute) UnmarshalJSON(data []byte) error {
	s := string(data)

	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		*t = SlashDateHourMinute{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(SlashDateHourMinuteFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = SlashDateHourMinute{parsed}

	return nil
}

func (t *SlashDateHourMinute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*t = SlashDateHourMinute{time.Time{}}

		return nil
	}

	parsed, err := time.Parse(SlashDateHourMinuteFormat, s)
	if err != nil {
		return fmt.Errorf("invalid time string '%s' => %w", s, err)
	}

	*t = SlashDateHourMinute{parsed}

	return nil
}

func (t SlashDateHourMinute) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Format(SlashDateHourMinuteFormat))), nil
}

func (t SlashDateHourMinute) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(t.Format(SlashDateHourMinuteFormat), start)
}

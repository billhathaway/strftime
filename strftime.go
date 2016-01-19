package strftime

import (
	"bytes"
	"fmt"
)

// conversions holds all the conversion characters that can be activated with '%'
var conversions = map[byte]string{
	'a': "Mon",         // day name abbreviated
	'A': "Monday",      // day name full
	'b': "Jan",         // month name abbreviated
	'C': "06",          // year/100 as a decimal number
	'd': "02",          // 2 digit day
	'D': "01/02/06",    // mm/dd/yy
	'e': "_2",          // day of the month as a decimal number (1-31); single digits are preceded by a blank
	'F': "2006-01-02",  // YYYY-MM-DD
	'H': "15",          // hours as decimal 01-24
	'I': "03",          // hours as decimal using 12 hour clock
	'M': "04",          // 2 digit minute
	'm': "01",          // month in decimal
	'n': "\n",          // newline
	'p': "PM",          // AM or PM
	'P': "pm",          // am or PM
	'r': "03:04:05 PM", // time in AM or PM notation
	'R': "15:04",       //HH:MM
	'S': "05",          // seconds as decimal
	't': "\t",          // tab
	'T': "15:04:05",    // time in 24 hour notation
	'Y': "2006",        // 4 digit year
	'z': "-0700",       // timezone offset from UTC
	'Z': "MST",         // timezone name or abbreviation
	'%': "%",
}

// New is passed a strftime(3) string and returns a string in time.Format() syntax
// error is non-nil if we don't know how to handle a conversion
func New(format string) (string, error) {
	buf := bytes.Buffer{}
	var special bool
	for i := range format {
		ch := format[i]
		if special {
			val, ok := conversions[ch]
			if !ok {
				return "", fmt.Errorf("unknown conversion specifier '%%%c'", ch)
			}
			buf.WriteString(val)
			special = false
			continue
		}
		if ch == '%' {
			special = true
			continue
		}
		buf.WriteByte(ch)
	}

	return buf.String(), nil
}

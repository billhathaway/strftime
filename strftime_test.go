package strftime

import (
	"testing"
	"time"
)

var referenceTime time.Time

func init() {
	var err error
	referenceTime, err = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		panic(err)
	}
}

func testConversion(t *testing.T, format string, expected string) {
	t.Logf("format = [%s]\n", format)
	formatted := referenceTime.Format(format)
	if formatted != expected {
		t.Errorf("formatted string does not match expected value [%s] != [%s]", formatted, expected)
	}
}

// TestConversions runs a test for each conversion handled by the package
func TestConversions(t *testing.T) {
	tests := map[string]string{
		"%a": "Mon",         // day name abbreviated
		"%A": "Monday",      // day name full
		"%b": "Jan",         // month name abbreviated
		"%C": "06",          // year/100 as a decimal number
		"%d": "02",          // 2 digit day
		"%D": "01/02/06",    // mm/dd/yy
		"%e": " 2",          // day of the month as a decimal number (1-31); single digits are preceded by a blank
		"%F": "2006-01-02",  // YYYY-MM-DD
		"%H": "15",          // hours as decimal 01-24
		"%I": "03",          // hours as decimal using 12 hour clock
		"%M": "04",          // 2 digit minute
		"%m": "01",          // month in decimal
		"%n": "\n",          // newline
		"%p": "PM",          // AM or PM
		"%P": "pm",          // am or PM
		"%r": "03:04:05 PM", // time in AM or PM notation
		"%R": "15:04",       // HH:MM
		"%S": "05",          // seconds as decimal
		"%t": "\t",          // tab
		"%T": "15:04:05",    // time in 24 hour notation
		"%Y": "2006",        // 4 digit year
		"%z": "+0000",       // timezone offset from UTC
		"%Z": "UTC",         // timezone name or abbreviation
		"%%": "%",
		"AB": "AB", // generic characters passed through
	}
	for k, expected := range tests {
		formatted, err := New(k)
		if err != nil {
			t.Error(err)
			continue
		}
		testConversion(t, formatted, expected)
	}

}

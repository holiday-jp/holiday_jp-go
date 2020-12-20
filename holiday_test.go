package holiday

import (
	"reflect"
	"testing"
	"time"
)

func TestGenDateStr(t *testing.T) {
	tests := []struct {
		time time.Time
		want string
	}{
		{
			time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			"1970-01-01",
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			"1970-01-01",
		},
		{
			time.Date(1991, 12, 19, 0, 0, 0, 0, time.UTC),
			"1991-12-19",
		},
	}

	for _, test := range tests {
		got := genDateStr(test.time)
		if test.want != got {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

func TestIsHoliday(t *testing.T) {
	tests := []struct {
		time time.Time
		want bool
	}{
		{
			time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			true,
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			true,
		},
		{
			time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
			false,
		},
		{
			time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC),
			false,
		},
	}

	for _, test := range tests {
		got := IsHoliday(test.time)
		if test.want != got {
			t.Fatalf("want '%t', but '%t':", test.want, got)
		}
	}
}

func TestHolidayName(t *testing.T) {
	tests := []struct {
		time time.Time
		want string
		err  bool
	}{
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			"元日",
			false,
		},
		{
			time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
			"",
			true,
		},
	}

	for _, test := range tests {
		got, err := HolidayName(test.time)
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", test.time, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", test.time)
		}
		if got != "" && got != test.want {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		t0   time.Time
		t1   time.Time
		want Holidays
	}{
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			Holidays{
				"1970-01-01": Holiday{
					"date":    "1970-01-01",
					"week":    "木",
					"week_en": "Thursday",
					"name":    "元日",
					"name_en": "New Year's Day",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
			Holidays{
				"1970-01-01": Holiday{
					"date":    "1970-01-01",
					"week":    "木",
					"week_en": "Thursday",
					"name":    "元日",
					"name_en": "New Year's Day",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 14, 1, 1, 1, 1, time.UTC),
			Holidays{
				"1970-01-01": Holiday{
					"date":    "1970-01-01",
					"week":    "木",
					"week_en": "Thursday",
					"name":    "元日",
					"name_en": "New Year's Day",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 15, 1, 1, 1, 1, time.UTC),
			Holidays{
				"1970-01-01": Holiday{
					"date":    "1970-01-01",
					"week":    "木",
					"week_en": "Thursday",
					"name":    "元日",
					"name_en": "New Year's Day",
				},
				"1970-01-15": Holiday{
					"date":    "1970-01-15",
					"week":    "木",
					"week_en": "Thursday",
					"name":    "成人の日",
					"name_en": "Coming of Age Da",
				},
			},
		},
	}

	for _, test := range tests {
		got := Between(test.t0, test.t1)

		if len(got) != len(test.want) {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
		if reflect.DeepEqual(got, test.want) {
			for k, v := range got {
				vw := test.want[k]
				for k, v := range v {
					if vw[k] != v {
						t.Fatalf("\nwant %q,\nbut  %q:", test.want, got)
					}
				}
			}
		}
	}
}

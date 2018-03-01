package holiday

import (
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
			t.Fatal("want %q, but %q:", test.want, got)
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
			t.Fatal("want %q, but %q:", test.want, got)
		}
	}
}

func TestHolidayName(t *testing.T) {
	tests := []struct {
		time time.Time
		want *Holiday
		err  bool
	}{
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			&Holiday{
				time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
				"元日",
			},
			false,
		},
		{
			time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
			&Holiday{
				time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
				"",
			},
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
		if got != nil && *got != *test.want {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		t0   time.Time
		t1   time.Time
		want []*Holiday
	}{
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			[]*Holiday{
				&Holiday{
					time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
					"元日",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 2, 1, 1, 1, 1, time.UTC),
			[]*Holiday{
				&Holiday{
					time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
					"元日",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 14, 1, 1, 1, 1, time.UTC),
			[]*Holiday{
				&Holiday{
					time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
					"元日",
				},
			},
		},
		{
			time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
			time.Date(1970, 1, 15, 1, 1, 1, 1, time.UTC),
			[]*Holiday{
				&Holiday{
					time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC),
					"元日",
				},
				&Holiday{
					time.Date(1970, 1, 15, 1, 1, 1, 1, time.UTC),
					"成人の日",
				},
			},
		},
	}

	for _, test := range tests {
		got := Between(test.t0, test.t1)

		if len(got) != len(test.want) {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
		for i, gotV := range got {
			if gotV == nil {
				t.Fatalf("want %q, but nil:", *test.want[i])
			}
			if *gotV != *test.want[i] {
				t.Fatalf("want %v, but %v:", *test.want[i], *gotV)
			}
		}
	}
}

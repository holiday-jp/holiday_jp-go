package holiday

import (
	"errors"
	"fmt"
	"time"
)

// Holiday holds holiday info.
type Holiday map[string]string

// New create a new Holiday.
// If `t` is not holiday, when return nil and error.
func New(t time.Time) (*Holiday, error) {
	holiday, ok := holidays[genDateStr(t)]
	if !ok {
		return nil, errors.New("There is no applicable holiday")
	}
	return &holiday, nil
}

// Date returns the day of the holiday.
func (h *Holiday) Date() (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, (*h)["date"]+" 00:00:00 +0000 UTC")
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Name returns name of the holiday.
// It behaves in the same way as the String function.
func (h *Holiday) Name() string {
	if h == nil {
		return ""
	}
	return (*h)["name"]
}

// NameEn returns english name of the holiday.
func (h *Holiday) NameEn() string {
	if h == nil {
		return ""
	}
	return (*h)["name_en"]
}

func (h *Holiday) Week() string {
	if h == nil {
		return ""
	}
	return (*h)["week"]
}

func (h *Holiday) WeekEn() string {
	if h == nil {
		return ""
	}
	return (*h)["week_en"]
}

// String returns name of the holiday.
func (h *Holiday) String() string {
	if h == nil {
		return ""
	}
	return (*h)["name"]
}

type Holidays map[string]Holiday

func (h *Holidays) Names() []string {
	if h == nil {
		return nil
	}
	ret := make([]string, 0, len(*h))
	for _, v := range *h {
		ret = append(ret, v.Name())
	}
	return ret
}

func genDateStr(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

// IsHoliday function checks whether the specified date is a holiday.
func IsHoliday(t time.Time) bool {
	_, ok := holidays[genDateStr(t)]
	return ok
}

// HolidayName function returns Holiday name string.
func HolidayName(t time.Time) (string, error) {
	holiday, ok := holidays[genDateStr(t)]
	if !ok {
		return "", errors.New("There is no applicable holiday")
	}
	return holiday.Name(), nil
}

// Between acquires the holiday of the designated section.
func Between(t0, t1 time.Time) Holidays {
	ret := Holidays{}
	for {
		if !t1.After(t0) && !t0.Equal(t1) {
			break
		}
		n, err := New(t0)
		if err != nil {
			t0 = t0.AddDate(0, 0, 1)
			continue
		}
		ret[(*n)["date"]] = *n
		t0 = t0.AddDate(0, 0, 1)
	}
	return ret
}

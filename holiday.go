package holiday

// go:generate statik -src ./datasheet/holidays.yml

import (
	"errors"
	"fmt"
	"time"

	// embed datasheet/holidays.yml in go code.
	_ "github.com/ieee0824/holiday_jp-go/statik"

	"github.com/go-yaml/yaml"
	"github.com/rakyll/statik/fs"
)

// Holiday struct holds holiday info.
// This struct is read-only.
type Holiday struct {
	t    time.Time
	name string
}

func new(name string, t time.Time) *Holiday {
	return &Holiday{
		t:    t,
		name: name,
	}
}

// Date returns the day of the holiday.
func (h *Holiday) Date() *time.Time {
	if h == nil {
		return nil
	}
	return &h.t
}

// Name returns name of the holiday.
// It behaves in the same way as the String function.
func (h *Holiday) Name() string {
	if h == nil {
		return ""
	}
	return h.name
}

// String returns name of the holiday.
func (h *Holiday) String() string {
	if h == nil {
		return ""
	}
	return h.name
}

// holidays holds the parse result of datasheet/holidays.yml.
var holidays = map[string]string{}

func init() {
	fs, err := fs.New()
	if err != nil {
		panic(err)
	}

	f, err := fs.Open("/.")
	if err != nil {
		panic(err)
	}

	if err := yaml.NewDecoder(f).Decode(holidays); err != nil {
		panic(err)
	}
}

func genDateStr(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

// IsHoliday function checks whether the specified date is a holiday.
func IsHoliday(t time.Time) bool {
	return holidays[genDateStr(t)] != ""
}

// HolidayName function returns Holiday struct pointer.
func HolidayName(t time.Time) (*Holiday, error) {
	name, ok := holidays[genDateStr(t)]
	if !ok {
		return nil, errors.New("There is no applicable holiday")
	}
	return new(name, t), nil
}

// Between acquires the holiday of the designated section.
func Between(t0, t1 time.Time) []*Holiday {
	ret := []*Holiday{}
	for {
		if !t1.After(t0) && !t0.Equal(t1) {
			break
		}
		n, err := HolidayName(t0)
		if err != nil {
			t0 = t0.AddDate(0, 0, 1)
			continue
		}
		ret = append(ret, n)
		t0 = t0.AddDate(0, 0, 1)
	}
	return ret
}

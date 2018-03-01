package holiday

// go:generate statik -src ./datasheet/holidays.yml

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-yaml/yaml"
	_ "github.com/ieee0824/holiday_jp-go/statik"
	"github.com/rakyll/statik/fs"
)

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

func IsHoliday(t time.Time) bool {
	return holidays[genDateStr(t)] != ""
}

func HolidayName(t time.Time) (string, error) {
	name, ok := holidays[genDateStr(t)]
	if !ok {
		return "", errors.New("There is no applicable holiday")
	}
	return name, nil
}

func Between(t0, t1 time.Time) []string {
	ret := []string{}
	for {
		if !t1.After(t0) {
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

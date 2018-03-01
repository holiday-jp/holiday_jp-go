# holiday_jp-go

[![GoDoc](https://godoc.org/github.com/ieee0824/holiday_jp-go?status.svg)](https://godoc.org/github.com/ieee0824/holiday_jp-go)

🎌 Japanese holiday for Go

## Requirements
* go 1.9 or later

## Installing

```
$ go get github.com/ieee0824/holiday_jp-go
```

## Example

```
import "github.com/ieee0824/holiday_jp-go"

if holiday.IsHoliday(time.Now()) {
    fmt.Println("today is holiday!")
}
```

## API

### `New(t time.Time) (*Holiday, error)`

New create a new Holiday.
If `t` is not holiday, when return nil and error.

### `IsHoliday(t time.Time) bool`

IsHoliday function checks whether the specified date is a holiday.

### `HolidayName(t time.Time) (string, error)`

HolidayName function returns Holiday name string.

### `Between(t0, t1 time.Time) Holidays`

Between acquires the holiday of the designated section.

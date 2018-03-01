# holiday_jp-go

🎌 Japanese holiday for Go

## Requirements
* go 1.9 or later

## Installing

```
& go get github.com/ieee0824/holiday_jp-go
```

## Example

```
import "github.com/ieee0824/holiday_jp-go"

if holiday.IsHoliday(time.Now()) {
    fmt.Println("today is holiday!")
}
```

## API

### `IsHoliday(t time.Time) bool`

IsHoliday function checks whether the specified date is a holiday.

### `HolidayName(t time.Time) (*Holiday, error)`

HolidayName function returns Holiday struct pointer.

### `Between(t0, t1 time.Time) []*Holiday`

Between acquires the holiday of the designated section.
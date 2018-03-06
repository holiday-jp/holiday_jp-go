# holiday_jp-go

[![GoDoc](https://godoc.org/github.com/holiday-jp/holiday_jp-go?status.svg)](https://godoc.org/github.com/holiday-jp/holiday_jp-go)

🎌 Japanese holiday for Go

## Requirements
* go 1.9 or later

## Installing

```bash
$ go get github.com/holiday-jp/holiday_jp-go
```

## Example

```go
import "github.com/holiday-jp/holiday_jp-go"

if holiday.IsHoliday(time.Now()) {
    fmt.Println("today is holiday!")
}
```

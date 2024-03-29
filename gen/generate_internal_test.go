package main

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	data, err := parse(filepath.Join("testdata", "test.yml"))
	if err != nil {
		t.Fatalf("parse failed, %v", err)
	}

	want := map[string]map[string]string{
		"1970-01-01": {
			"date":    "1970-01-01",
			"week":    "木",
			"week_en": "Thursday",
			"name":    "元日",
			"name_en": "New Year's Day",
		},
		"1970-11-23": {
			"date":    "1970-11-23",
			"week":    "月",
			"week_en": "Monday",
			"name":    "勤労感謝の日",
			"name_en": "Labor Thanksgiving Day",
		},
	}

	if !reflect.DeepEqual(data, want) {
		t.Errorf("unmatch result, want %+v, got %+v", want, data)
	}
}

func TestGenerate(t *testing.T) {
	in := map[string]map[string]string{
		"1970-01-01": {
			"date":    "1970-01-01",
			"week":    "木",
			"week_en": "Thursday",
			"name":    "元日",
			"name_en": "New Year's Day",
		},
		"1970-11-23": {
			"date":    "1970-11-23",
			"week":    "月",
			"week_en": "Monday",
			"name":    "勤労感謝の日",
			"name_en": "Labor Thanksgiving Day",
		},
	}

	var buf strings.Builder

	if err := generate(in, &buf); err != nil {
		t.Fatalf("generate failed, %v", err)
	}

	want := `
// Code generated by gen/generate.go. DO NOT EDIT.
// Generate from datasheet/holidays_detailed.yml

package holiday

// holidays holds the parse result of datasheet/holidays_detailed.yml
var holidays = Holidays{
	"1970-01-01": Holiday{
		"date":    "1970-01-01",
		"week":    "木",
		"week_en": "Thursday",
		"name":    "元日",
		"name_en": "New Year's Day",
	},
	"1970-11-23": Holiday{
		"date":    "1970-11-23",
		"week":    "月",
		"week_en": "Monday",
		"name":    "勤労感謝の日",
		"name_en": "Labor Thanksgiving Day",
	},
}
`

	if buf.String() != want {
		t.Errorf("unmatch result, want %q, got %q", want, buf.String())
	}
}

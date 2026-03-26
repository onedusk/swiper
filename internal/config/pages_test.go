package config

import (
	"reflect"
	"strings"
	"testing"
)

func TestParsePageRanges_SingleRange(t *testing.T) {
	ranges, err := ParsePageRanges("1-10")
	if err != nil {
		t.Fatal(err)
	}
	if len(ranges) != 1 || ranges[0].Start != 1 || ranges[0].End != 10 {
		t.Fatalf("expected [{1,10}], got %v", ranges)
	}
}

func TestParsePageRanges_MultipleRanges(t *testing.T) {
	ranges, err := ParsePageRanges("1-10,50-60,75")
	if err != nil {
		t.Fatal(err)
	}
	expected := []PageRange{{1, 10}, {50, 60}, {75, 75}}
	if !reflect.DeepEqual(ranges, expected) {
		t.Fatalf("expected %v, got %v", expected, ranges)
	}
}

func TestParsePageRanges_SinglePage(t *testing.T) {
	ranges, err := ParsePageRanges("5")
	if err != nil {
		t.Fatal(err)
	}
	if len(ranges) != 1 || ranges[0].Start != 5 || ranges[0].End != 5 {
		t.Fatalf("expected [{5,5}], got %v", ranges)
	}
}

func TestParsePageRanges_Empty(t *testing.T) {
	ranges, err := ParsePageRanges("")
	if err != nil {
		t.Fatal(err)
	}
	if ranges != nil {
		t.Fatalf("expected nil, got %v", ranges)
	}
}

func TestParsePageRanges_InvalidText(t *testing.T) {
	_, err := ParsePageRanges("abc")
	if err == nil {
		t.Fatal("expected error for 'abc'")
	}
}

func TestParsePageRanges_ReversedRange(t *testing.T) {
	_, err := ParsePageRanges("10-5")
	if err == nil {
		t.Fatal("expected error for reversed range")
	}
	if !strings.Contains(err.Error(), "end") {
		t.Fatalf("expected end < start error, got: %v", err)
	}
}

func TestParsePageRanges_Sorted(t *testing.T) {
	ranges, err := ParsePageRanges("50,1-10,25")
	if err != nil {
		t.Fatal(err)
	}
	if ranges[0].Start != 1 || ranges[1].Start != 25 || ranges[2].Start != 50 {
		t.Fatalf("expected sorted by start, got %v", ranges)
	}
}

func TestExpandPages_WithRanges(t *testing.T) {
	ranges := []PageRange{{1, 3}, {5, 5}}
	pages := ExpandPages(ranges, 10)
	expected := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(pages, expected) {
		t.Fatalf("expected %v, got %v", expected, pages)
	}
}

func TestExpandPages_NilReturnsAll(t *testing.T) {
	pages := ExpandPages(nil, 5)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(pages, expected) {
		t.Fatalf("expected %v, got %v", expected, pages)
	}
}

func TestExpandPages_CappedAtMax(t *testing.T) {
	ranges := []PageRange{{1, 100}}
	pages := ExpandPages(ranges, 5)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(pages, expected) {
		t.Fatalf("expected %v, got %v", expected, pages)
	}
}

func TestParsePageRanges_ZeroPage(t *testing.T) {
	_, err := ParsePageRanges("0")
	if err == nil {
		t.Fatal("expected error for page 0")
	}
}

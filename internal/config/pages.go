package config

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// PageRange represents an inclusive range of page numbers.
type PageRange struct {
	Start int
	End   int
}

// ParsePageRanges parses a page range string like "1-10,50-60,75".
// Returns nil, nil for empty input.
func ParsePageRanges(s string) ([]PageRange, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	parts := strings.Split(s, ",")
	ranges := make([]PageRange, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			start, err := strconv.Atoi(strings.TrimSpace(bounds[0]))
			if err != nil {
				return nil, fmt.Errorf("invalid page number %q: %w", bounds[0], err)
			}
			end, err := strconv.Atoi(strings.TrimSpace(bounds[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid page number %q: %w", bounds[1], err)
			}
			if end < start {
				return nil, fmt.Errorf("invalid page range %q: end (%d) < start (%d)", part, end, start)
			}
			if start < 1 {
				return nil, fmt.Errorf("invalid page number: %d (must be >= 1)", start)
			}
			ranges = append(ranges, PageRange{Start: start, End: end})
		} else {
			page, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid page number %q: %w", part, err)
			}
			if page < 1 {
				return nil, fmt.Errorf("invalid page number: %d (must be >= 1)", page)
			}
			ranges = append(ranges, PageRange{Start: page, End: page})
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	return ranges, nil
}

// ExpandPages expands page ranges into a sorted slice of individual page numbers,
// capped at maxPage. If ranges is nil, returns all pages [1..maxPage].
func ExpandPages(ranges []PageRange, maxPage int) []int {
	if ranges == nil {
		pages := make([]int, maxPage)
		for i := 0; i < maxPage; i++ {
			pages[i] = i + 1
		}
		return pages
	}

	seen := make(map[int]bool)
	for _, r := range ranges {
		end := r.End
		if end > maxPage {
			end = maxPage
		}
		for p := r.Start; p <= end; p++ {
			if p >= 1 {
				seen[p] = true
			}
		}
	}

	pages := make([]int, 0, len(seen))
	for p := range seen {
		pages = append(pages, p)
	}
	sort.Ints(pages)
	return pages
}

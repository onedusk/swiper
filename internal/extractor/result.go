package extractor

import (
	"fmt"
	"strings"
	"time"
)

// PageResult captures the complete outcome of processing a single PDF page.
type PageResult struct {
	Page        int
	Text        string
	TextErr     error
	Images      []string
	ImageErrors []error
	Duration    time.Duration
}

// Success returns true if both text and image extraction succeeded fully.
func (pr *PageResult) Success() bool {
	return pr.TextErr == nil && len(pr.ImageErrors) == 0
}

// PartialSuccess returns true if at least some content was extracted
// despite errors in either text or image extraction.
func (pr *PageResult) PartialSuccess() bool {
	hasContent := strings.TrimSpace(pr.Text) != "" || len(pr.Images) > 0
	hasErrors := pr.TextErr != nil || len(pr.ImageErrors) > 0
	return hasContent && hasErrors
}

// ErrorSummary returns a human-readable summary of all errors, or empty string if none.
func (pr *PageResult) ErrorSummary() string {
	var parts []string
	if pr.TextErr != nil {
		parts = append(parts, fmt.Sprintf("text: %v", pr.TextErr))
	}
	for i, err := range pr.ImageErrors {
		if err != nil {
			parts = append(parts, fmt.Sprintf("image[%d]: %v", i, err))
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return fmt.Sprintf("page %d errors: %s", pr.Page, strings.Join(parts, "; "))
}

// MarkdownErrorComment returns an HTML comment suitable for embedding
// at the top of the page markdown file. Returns empty string if no errors.
func (pr *PageResult) MarkdownErrorComment() string {
	summary := pr.ErrorSummary()
	if summary == "" {
		return ""
	}
	return fmt.Sprintf("<!-- EXTRACTION WARNING: %s -->\n\n", summary)
}

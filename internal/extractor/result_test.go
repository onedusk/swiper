package extractor

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPageResult_Success(t *testing.T) {
	cases := []struct {
		name string
		pr   PageResult
		want bool
	}{
		{"full success", PageResult{Page: 1, Text: "hello"}, true},
		{"text error", PageResult{Page: 1, TextErr: fmt.Errorf("fail")}, false},
		{"image error", PageResult{Page: 1, ImageErrors: []error{fmt.Errorf("x")}}, false},
		{"both errors", PageResult{Page: 1, TextErr: fmt.Errorf("a"), ImageErrors: []error{fmt.Errorf("b")}}, false},
		{"empty but no errors", PageResult{Page: 1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.pr.Success(); got != tc.want {
				t.Fatalf("Success() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPageResult_PartialSuccess(t *testing.T) {
	cases := []struct {
		name string
		pr   PageResult
		want bool
	}{
		{"full success is not partial", PageResult{Page: 1, Text: "hello"}, false},
		{"text ok + image errors", PageResult{Page: 1, Text: "hello", Images: []string{"a.jpg"}, ImageErrors: []error{fmt.Errorf("x")}}, true},
		{"text err + has images", PageResult{Page: 1, TextErr: fmt.Errorf("fail"), Images: []string{"a.jpg"}}, true},
		{"total failure no content", PageResult{Page: 1, TextErr: fmt.Errorf("fail")}, false},
		{"no errors no content", PageResult{Page: 1}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.pr.PartialSuccess(); got != tc.want {
				t.Fatalf("PartialSuccess() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPageResult_ErrorSummary(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		pr := PageResult{Page: 1, Text: "ok"}
		if got := pr.ErrorSummary(); got != "" {
			t.Fatalf("expected empty, got %q", got)
		}
	})

	t.Run("text error only", func(t *testing.T) {
		pr := PageResult{Page: 2, TextErr: fmt.Errorf("pdftotext failed")}
		got := pr.ErrorSummary()
		if !strings.Contains(got, "page 2") {
			t.Fatalf("expected 'page 2', got %q", got)
		}
		if !strings.Contains(got, "text: pdftotext failed") {
			t.Fatalf("expected text error, got %q", got)
		}
	})

	t.Run("image errors only", func(t *testing.T) {
		pr := PageResult{Page: 3, ImageErrors: []error{fmt.Errorf("copy failed"), fmt.Errorf("timeout")}}
		got := pr.ErrorSummary()
		if !strings.Contains(got, "image[0]: copy failed") {
			t.Fatalf("expected image[0] error, got %q", got)
		}
		if !strings.Contains(got, "image[1]: timeout") {
			t.Fatalf("expected image[1] error, got %q", got)
		}
	})

	t.Run("combined errors", func(t *testing.T) {
		pr := PageResult{Page: 1, TextErr: fmt.Errorf("t"), ImageErrors: []error{fmt.Errorf("i")}}
		got := pr.ErrorSummary()
		if !strings.Contains(got, "text: t") || !strings.Contains(got, "image[0]: i") {
			t.Fatalf("expected combined errors, got %q", got)
		}
	})
}

func TestPageResult_MarkdownErrorComment(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		pr := PageResult{Page: 1, Text: "ok"}
		if got := pr.MarkdownErrorComment(); got != "" {
			t.Fatalf("expected empty, got %q", got)
		}
	})

	t.Run("with errors", func(t *testing.T) {
		pr := PageResult{Page: 1, TextErr: fmt.Errorf("fail")}
		got := pr.MarkdownErrorComment()
		if !strings.HasPrefix(got, "<!-- EXTRACTION WARNING:") {
			t.Fatalf("expected HTML comment, got %q", got)
		}
		if !strings.HasSuffix(got, "-->\n\n") {
			t.Fatalf("expected comment ending, got %q", got)
		}
	})
}

func TestPageResult_Duration(t *testing.T) {
	pr := PageResult{Page: 1, Duration: 500 * time.Millisecond}
	if pr.Duration != 500*time.Millisecond {
		t.Fatalf("expected 500ms, got %v", pr.Duration)
	}
}

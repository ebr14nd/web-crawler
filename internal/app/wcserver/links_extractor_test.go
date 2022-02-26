package wcserver

import (
	"testing"
)

func TestExtractLinks(t *testing.T) {
	data, err := readTestFile("simple-link.html")
	if err != nil {
		t.Error(err)
	}
	got, err := extractLinks(data)
	want := "/testdata"

	if err != nil {
		t.Error(err)
	}

	if !contains(got, want) {
		t.Errorf(" %q not found in extracted links %v", want, got)
	}
}

func TestExtractLinksExternal(t *testing.T) {
	data, err := readTestFile("simple-external-link.html")
	if err != nil {
		t.Error(err)
	}
	got, err := extractLinks(data)
	want := "http://example.com"

	if err != nil {
		t.Error(err)
	}

	if !contains(got, want) {
		t.Errorf(" %q not found in extracted links %v", want, got)
	}
}

func TestExtractLinksNotHtml(t *testing.T) {
	data, err := readTestFile("hello-world.html")
	if err != nil {
		t.Error(err)
	}
	got, err := extractLinks(data)

	if err != nil {
		t.Error(err)
	}

	if len(got) != 0 {
		t.Error("Links found in a non HTML file")
	}
}

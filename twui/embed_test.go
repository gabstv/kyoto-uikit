package twui

import "testing"

func TestEmbed(t *testing.T) {
	if err := ExtractTemplates("."); err != nil {
		t.Fatal(err)
	}
}

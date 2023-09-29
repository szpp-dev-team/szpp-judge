package langs

import (
	"strings"
	"testing"
)

func TestLangIDShouldBeUnique(t *testing.T) {
	ids := map[LangID]struct{}{}

	for _, m := range langMetas {
		if _, used := ids[m.ID]; used {
			t.Fatalf("Dupulicated LangID: %q", m.ID)
		}
		ids[m.ID] = struct{}{}
	}
}

func TestLangNameShouldBeUnique(t *testing.T) {
	ids := map[string]struct{}{}

	for _, m := range langMetas {
		if _, used := ids[m.Name]; used {
			t.Fatalf("Dupulicated Name: %q", m.Name)
		}
		ids[m.Name] = struct{}{}
	}
}

func TestLangIDFormat(t *testing.T) {
	for _, m := range langMetas {
		if strings.Count(string(m.ID), "/") != 2 {
			t.Fatalf("Format of LangID must be '{languageName}/{version}/{implementation}', but got %q", m.ID)
		}
	}
}

func TestDockerImageName(t *testing.T) {
	for _, m := range langMetas {
		if !strings.HasPrefix(string(m.DockerImage), ImagePrefix) {
			t.Fatalf("DockerImage of %s must have prefix %q", m.ID, ImagePrefix)
		}
	}
}

func TestNonEmptyField(t *testing.T) {
	checkStringField := func(t *testing.T, id LangID, value, field string) {
		t.Helper()
		if strings.TrimSpace(value) != value {
			t.Fatalf("%s of %s contains extra whitespaces: %q", field, id, value)
		}
		if len(value) == 0 {
			t.Fatalf("%s of %s is empty", field, id)
		}
		if strings.ContainsRune(value, '　') {
			t.Fatalf("%s of %s contains zenkaku-space", field, id)
		}
	}

	for _, m := range langMetas {
		checkStringField(t, m.ID, string(m.ID), "ID")
		checkStringField(t, m.ID, string(m.Name), "Name")
		checkStringField(t, m.ID, string(m.DockerImage), "DockerImage")
		checkStringField(t, m.ID, string(m.SourceFile), "SourceFile")

		if len(m.CompileCmd) == 0 {
			t.Fatalf("CompileCmd of %s is empty", m.ID)
		}
		if len(m.ExecCmd) == 0 {
			t.Fatalf("ExecCmd of %s is empty", m.ID)
		}
	}
}

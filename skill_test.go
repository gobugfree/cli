package cli

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"testing"
)

// The skill names operations by id (the gbf run vocabulary) as examples;
// every id must exist in the platform document the client was generated
// from, so the skill can never teach an operation that is gone.
func TestSkillNamesLiveOperations(t *testing.T) {
	if !strings.HasPrefix(Skill, "---\nname: "+SkillName+"\n") {
		t.Fatalf("SKILL.md must start with frontmatter naming %s", SkillName)
	}
	data, err := os.ReadFile("openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Paths map[string]map[string]struct {
			OperationId string `json:"operationId"`
		} `json:"paths"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	live := map[string]bool{}
	for _, item := range doc.Paths {
		for _, op := range item {
			live[op.OperationId] = true
		}
	}
	ids := regexp.MustCompile("`((?:get|post|put|patch|delete)[A-Z][A-Za-z]*)`").FindAllStringSubmatch(Skill, -1)
	if len(ids) == 0 {
		t.Fatal("the skill names no operation ids")
	}
	for _, match := range ids {
		if !live[match[1]] {
			t.Errorf("the skill names %s but openapi.json does not serve it", match[1])
		}
	}
}

func TestSkillVersion(t *testing.T) {
	v := SkillVersion("v3.25.0")
	if !strings.HasPrefix(v, "v3.25.0+") || len(v) != len("v3.25.0+")+12 {
		t.Fatalf("version %q", v)
	}
}

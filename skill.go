// Package cli holds the agent skill that ships with gbf: an installable
// SKILL.md (frontmatter with name and description, then the instructions)
// that teaches an AI agent how to work the platform through the command.
package cli

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
)

// Skill is the SKILL.md as shipped; gbf skill prints or installs it.
//
//go:embed SKILL.md
var Skill string

// SkillName is the skill's name from its frontmatter, and the directory an
// agent installs it into.
const SkillName = "gbf"

// SkillVersion stamps the skill with the gbf version plus a short SHA-256 of
// the content, so an agent notices both a new binary and a changed text.
func SkillVersion(cliVersion string) string {
	sum := sha256.Sum256([]byte(Skill))
	return cliVersion + "+" + hex.EncodeToString(sum[:])[:12]
}

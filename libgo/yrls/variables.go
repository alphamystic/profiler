package yrls

var ruleTemplate := `
rule {{ .RuleName }} {
	meta:
		description = "{{ .Description }}"
		author = "Your Name"
		reference = "Any relevant reference"

	strings:
		$string1 = "example_string" // Define strings or patterns to match

	condition:
		$string1 // Define the condition for the rule
}
`

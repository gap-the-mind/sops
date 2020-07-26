package toml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mozilla.org/sops/v3"
)

// Example form https://en.wikipedia.org/wiki/TOML
var PLAIN = []byte(`
# This is a TOML document.

title = "TOML Example"

[owner]
name = "Tom Preston-Werner"
dob = 1979-05-27T07:32:00-08:00 # First class dates

[database]
server = "192.168.1.1"
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # Indentation (tabs and/or spaces) is allowed but not required
  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

[clients]
data = [ ["gamma", "delta"], [1, 2] ]

# Line breaks are OK when inside arrays
hosts = [
  "alpha",
  "omega"
]
`)

var BRANCHES = sops.TreeBranches{sops.TreeBranch{
	sops.TreeItem{
		Key:   "key1",
		Value: "value",
	},
	sops.TreeItem{
		Key:   "key1_a",
		Value: "value",
	},
	sops.TreeItem{
		Key:   sops.Comment{Value: " ^ comment 1"},
		Value: nil,
	},
},
	sops.TreeBranch{
		sops.TreeItem{
			Key:   "key2",
			Value: "value2",
		},
	},
}

var plainNested = []byte(`
[glossary]
title = "example glossary"

[glossary.GlossDiv]
title="S"

[glossary.GlossDiv.GlossList.GlossEntry]
ID = "SGML"
SortAs = "SGML"
GlossTerm = "Standard Generalized Markup Language"
Acronym ="SGML"
Abbrev ="ISO 8879:1986"
GlossSee = "markup"
               

[glossary.GlossDiv.GlossList.GlossEntry.GlossDef] 
para = "A meta-markup language, used to create markup languages such as DocBook."
GlossSeeAlso =[
	"GML", 
	"XML"
]
`)

var branchesNested = sops.TreeBranches{
	sops.TreeBranch{
		sops.TreeItem{
			Key: "glossary",
			Value: sops.TreeBranch{
				sops.TreeItem{
					Key:   "title",
					Value: "example glossary",
				},
				sops.TreeItem{
					Key: "GlossDiv",
					Value: sops.TreeBranch{
						sops.TreeItem{
							Key:   "title",
							Value: "S",
						},
						sops.TreeItem{
							Key: "GlossList",
							Value: sops.TreeBranch{
								sops.TreeItem{
									Key: "GlossEntry",
									Value: sops.TreeBranch{
										sops.TreeItem{
											Key:   "ID",
											Value: "SGML",
										},
										sops.TreeItem{
											Key:   "SortAs",
											Value: "SGML",
										},
										sops.TreeItem{
											Key:   "GlossTerm",
											Value: "Standard Generalized Markup Language",
										},
										sops.TreeItem{
											Key:   "Acronym",
											Value: "SGML",
										},
										sops.TreeItem{
											Key:   "Abbrev",
											Value: "ISO 8879:1986",
										},
										sops.TreeItem{
											Key:   "GlossSee",
											Value: "markup",
										},
										sops.TreeItem{
											Key: "GlossDef",
											Value: sops.TreeBranch{
												sops.TreeItem{
													Key:   "para",
													Value: "A meta-markup language, used to create markup languages such as DocBook.",
												},
												sops.TreeItem{
													Key: "GlossSeeAlso",
													Value: []interface{}{
														"GML",
														"XML",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestLoadPlainFile(t *testing.T) {
	branches, err := (&Store{}).LoadPlainFile(PLAIN)
	assert.Nil(t, err)
	assert.Equal(t, BRANCHES, branches)
}

func TestLoadPlainCompledFile(t *testing.T) {
	branches, err := (&Store{}).LoadPlainFile(plainNested)
	assert.Nil(t, err)
	assert.Equal(t, branchesNested, branches)
}

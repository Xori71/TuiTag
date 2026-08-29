package parser

import (
	"github.com/pelletier/go-toml/v2"
)

// Maps tag fields to single/multi line bubble widgets
var inputFieldKindMap = map[string]int{
	"title":       0,
	"album":       0,
	"artist":      0,
	"albumartist": 0,
	"composer":    0,
	"year":        0,
	"genre":       0,
	"tracknumber": 0,
	"tracktotal":  0,
	"discnumber":  0,
	"disctotal":   0,
	"lyrics":      1,
	"comment":     1,
}

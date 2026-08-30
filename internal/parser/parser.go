package parser

import (
	"os"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	toml "github.com/pelletier/go-toml/v2"
)

// Maps tag fields to single/multi line bubble widgets
var fieldToWidgetType = map[string]int{
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

var fieldsToPrettyFields = map[string]string{
	"title":       "Title",
	"album":       "Album",
	"artist":      "Artist",
	"albumartist": "Album Artist",
	"composer":    "Composer",
	"year":        "Year",
	"genre":       "Genre",
	"tracknumber": "Track Number",
	"tracktotal":  "Track Total",
	"discnumber":  "Disc Number",
	"disctotal":   "Disc Total",
	"lyrics":      "Lyrics",
	"comment":     "Comment",
}

const configDir = "../../config.toml"

type Field struct {
	Title string
	Input textinput.Model // Active if widget type == 0
	Area  textarea.Model  // Active if widget type == 1
}

type Config struct {
	DefaultDir string
	TagOrder   []string
}

func ParseConfigFile() (Config, error) {
	data, err := os.ReadFile(configDir)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = toml.Unmarshal([]byte(data), &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetInputFields(fields []string) []Field {
	var inputFields []Field
	for _, field := range fields {
		switch fieldToWidgetType[field] {
		case 0:
			inputFields = append(inputFields, Field{
				Title: fieldsToPrettyFields[field],
				Input: textinput.New(),
			})
			break
		case 1:
			inputFields = append(inputFields, Field{
				Title: fieldsToPrettyFields[field],
				Area:  textarea.New(),
			})
			break
		}
	}

	return nil
}

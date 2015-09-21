package publish_prod

import (
	"laatoocore"
)

const (
	ENTITY_WORD_NAME = "Word"
)

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_WORD_NAME, NewWord)
}

type Word struct {
	WordGur  string `json:"WordGur" bson:"WordGur"`
	WordEng  string `json:"WordEng" bson:"WordEng"`
	Reviewed bool   `json:"Reviewed" bson:"Reviewed"`
}

func NewWord(conf map[string]interface{}) (interface{}, error) {
	word := &Word{Reviewed: false}
	return word, nil
}

func (ent *Word) GetId() string {
	return ent.WordGur
}
func (ent *Word) SetId(id string) {
	ent.WordGur = id
}

func (ent *Word) GetIdField() string {
	return "WordGur"
}

func (ent *Word) PreSave() error {
	return nil
}
func (ent *Word) PostLoad() error {
	return nil
}

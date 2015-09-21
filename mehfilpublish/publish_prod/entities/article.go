package publish_prod

import (
	"github.com/twinj/uuid"
	"laatoocore"
)

const (
	ENTITY_ARTICLE_NAME = "Article"
)

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_ARTICLE_NAME, NewArticle)
}

type Article struct {
	Id         string `json:"Id" bson:"Id"`
	CreatedBy  string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy  string `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn  string `json:"UpdatedOn" bson:"UpdatedOn"`
	Title      string `json:"Title" bson:"Title"`
	BodyGur    string `json:"BodyGur" bson:"BodyGur"`
	SummaryGur string `json:"SummaryGur" bson:"SummaryGur"`
	BodyEng    string `json:"BodyEng" bson:"BodyEng"`
	SummaryEng string `json:"SummaryEng" bson:"SummaryEng"`
	TitleEng   string `json:"TitleEng" bson:"TitleEng"`
	Image      string `json:"Image" bson:"Image"`
	Type       []string
}

func NewArticle(conf map[string]interface{}) (interface{}, error) {
	art := &Article{}
	art.Id = uuid.NewV4().String()
	return art, nil
}

func (ent *Article) GetId() string {
	return ent.Id
}
func (ent *Article) SetId(id string) {
	ent.Id = id
}

func (ent *Article) GetIdField() string {
	return "Id"
}

func (ent *Article) PreSave() error {
	return nil
}
func (ent *Article) PostLoad() error {
	return nil
}

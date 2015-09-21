package publish_prod

import (
	"github.com/twinj/uuid"
	"laatoocore"
)

const (
	ENTITY_VIDEO_NAME = "Video"
)

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_VIDEO_NAME, NewVideo)
}

type Video struct {
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
	Video      string `json:"Video" bson:"Video"`
}

func NewVideo(conf map[string]interface{}) (interface{}, error) {
	art := &Video{}
	art.Id = uuid.NewV4().String()
	return art, nil
}

func (ent *Video) GetId() string {
	return ent.Id
}
func (ent *Video) SetId(id string) {
	ent.Id = id
}

func (ent *Video) GetIdField() string {
	return "Id"
}

func (ent *Video) PreSave() error {
	return nil
}
func (ent *Video) PostLoad() error {
	return nil
}

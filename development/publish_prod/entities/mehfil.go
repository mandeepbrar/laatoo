package publish_prod

import (
	"github.com/twinj/uuid"
	"laatoocore"
)

const (
	ENTITY_MEHFIL_NAME = "Mehfil"
)

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_MEHFIL_NAME, NewMehfil)
}

type Mehfil struct {
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
}

func NewMehfil(conf map[string]interface{}) (interface{}, error) {
	art := &Mehfil{}
	art.Id = uuid.NewV4().String()
	return art, nil
}

func (ent *Mehfil) GetId() string {
	return ent.Id
}
func (ent *Mehfil) SetId(id string) {
	ent.Id = id
}

func (ent *Mehfil) GetIdField() string {
	return "Id"
}

func (ent *Mehfil) PreSave() error {
	return nil
}
func (ent *Mehfil) PostLoad() error {
	return nil
}

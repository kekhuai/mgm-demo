package book

import "github.com/Kamva/mgm/v2"

type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func (model *Book) Creating() error {
	if err := model.DefaultModel.Creating(); nil != err {
		return err
	}
	model.Name = "before save"
	return nil
}

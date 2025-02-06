package dao

import (
	"errors"

	"github.com/uidea/artwork-backend/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) UpsertAbout(
	Content string,
	Cover string,
	CoverBlur string,
) error {
	about := model.About{}
	_, err := about.First(d.engine)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			about.Content = Content
			about.Cover = Cover
			about.CoverBlur = CoverBlur
			return about.Create(d.engine)
		}
		return err
	}

	about.Content = Content
	about.Cover = Cover
	about.CoverBlur = CoverBlur

	return about.Save(d.engine, &about)
}

func (d *Dao) GetAbout() (*model.About, error) {
	about := model.About{}
	return about.Get(d.engine)
}

func (d *Dao) UpdateAbout(
	id uint32,
	content string,
	cover string,
	coverBlur string,
) error {
	about := model.About{
		ID: id,
	}
	values := map[string]interface{}{
		"content":    content,
		"cover":      cover,
		"cover_blur": coverBlur,
	}
	return about.Update(d.engine, values)
}

// func (d *Dao) GetAboutFirst() (*model.About, error) {
// 	about := model.About{}
// 	return about.Get(d.engine)
// }

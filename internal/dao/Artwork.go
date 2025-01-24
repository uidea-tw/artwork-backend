package dao

import (
	"github.com/uidea/artwork-backend/internal/model"
)

func (d *Dao) CreateArtwork(
	Title string,
	Cover string,
	CoverBlur string,
	Description string,
	Content string,
	Author string,
	Year uint32,
	Length uint32,
	Width uint32,
	Height uint32,
	Price uint32,
) error {
	artwork := model.Artwork{
		Title:       Title,
		Cover:       Cover,
		CoverBlur:   CoverBlur,
		Description: Description,
		Content:     Content,
		Author:      Author,
		Year:        Year,
		Length:      Length,
		Width:       Width,
		Height:      Height,
		Price:       Price,
	}
	return artwork.Create(d.engine)
}

func (d *Dao) GetArtwork(id uint32) (*model.Artwork, error) {
	artwork := model.Artwork{ID: id}
	return artwork.Get(d.engine)
}

func (d *Dao) DeleteArtwork(id uint32) error {
	artwork := model.Artwork{ID: id}
	return artwork.Delete(d.engine)
}

func (d *Dao) ListArtWork() ([]*model.Artwork, error) {
	artwork := model.Artwork{}
	return artwork.List(d.engine)
}

func (d *Dao) UpdateArtwork(
	id uint32,
	title string,
	cover string,
	coverBlur string,
	description string,
	content string,
	author string,
	year uint32,
	length uint32,
	width uint32,
	height uint32,
	price uint32,
) error {
	artwork := model.Artwork{
		ID: id,
	}
	values := map[string]interface{}{
		"title":       title,
		"cover":       cover,
		"cover_blur":  coverBlur,
		"description": description,
		"content":     content,
		"author":      author,
		"year":        year,
		"length":      length,
		"width":       width,
		"height":      height,
		"price":       price,
	}
	return artwork.Update(d.engine, values)
}

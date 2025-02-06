package dao

import "github.com/uidea/artwork-backend/internal/model"

func (d *Dao) GetUploadFile(uuid string) (model.File, error) {
	file := model.File{ID: uuid}
	return file.Get(d.engine)
}

func (d *Dao) CreateUploadFile(
	ID string,
	Filename string,
	StorageName string,
	Bucket string,
	ContentType string,
	Size int64,
) error {
	file := model.File{
		ID:          ID,
		Filename:    Filename,
		StorageName: StorageName,
		Bucket:      Bucket,
		ContentType: ContentType,
		Size:        Size,
	}
	return file.Create(d.engine)
}

func (d *Dao) DeleteUploadFile(uuid string) error {
	file := model.File{ID: uuid}
	return file.Delete(d.engine)
}

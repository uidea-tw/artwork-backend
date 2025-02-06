package service

type UploadFile struct {
	ID          string `json:"ID"`
	Filename    string `json:"filename"`
	StorageName string `json:"storage_name"`
	Bucket      string `json:"bucket"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}

func (svc *Service) GetUploadFile(uuid string) (*UploadFile, error) {
	file, err := svc.dao.GetUploadFile(uuid)
	if err != nil {
		return nil, err
	}

	return &UploadFile{
		ID:          file.ID,
		Filename:    file.Filename,
		StorageName: file.StorageName,
		Bucket:      file.Bucket,
		ContentType: file.ContentType,
		Size:        file.Size,
	}, nil
}

func (svc *Service) CreateUploadFile(param UploadFile) error {
	return svc.dao.CreateUploadFile(
		param.ID,
		param.Filename,
		param.StorageName,
		param.Bucket,
		param.ContentType,
		param.Size,
	)
}

func (svc *Service) DeleteUploadFile(uuid string) error {
	return svc.dao.DeleteUploadFile(uuid)
}

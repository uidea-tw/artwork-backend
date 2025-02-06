package app

func FileTypeFormat(ContentType string) string {
	var ext string

	switch ContentType {
	case "image/jpeg", "image/jpg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	case "image/webp":
		ext = ".webp"
	case "image/svg+xml":
		ext = ".svg"
	case "application/pdf":
		ext = ".pdf"
	case "application/json":
		ext = ".json"
	case "application/xml":
		ext = ".xml"
	case "text/plain":
		ext = ".txt"
	case "text/html":
		ext = ".html"
	case "audio/mpeg":
		ext = ".mp3"
	case "video/mp4":
		ext = ".mp4"
	default:
		ext = ""
	}

	return ext
}

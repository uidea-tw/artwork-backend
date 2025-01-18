package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "取得標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")

	ErrorGetArticleFail    = NewError(20020001, "取得單篇文章失敗")
	ErrorGetArticlesFail   = NewError(20020002, "取得多篇文章失敗")
	ErrorCreateArticleFail = NewError(20020003, "建立文章失敗")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失敗")
	ErrorDeleteArticleFail = NewError(20020005, "刪除文章失敗")

	ErrorUploadFileFail = NewError(20030001, "上傳檔案失敗")
)

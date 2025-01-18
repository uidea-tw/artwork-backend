package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "伺服器錯誤")
	InvalidParams             = NewError(1001, "帶入參數錯誤")
	NotFound                  = NewError(1002, "找不到相關資訊")
	UnauthorizedAuthNotExist  = NewError(1003, "驗證失敗，找不報對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1004, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(1005, "驗證失敗，Token逾期")
	UnauthorizedTokenGenerate = NewError(1006, "驗證失敗，Token產生失敗")
	NotFoundRegisterAccount   = NewError(1007, "找尋不到註冊帳號")
	TooManyRequests           = NewError(1008, "請求過多")
	AlreadyExistsRecord       = NewError(1009, "該資料已存在")
)

package e

type ControllerError struct {
	Code     int    `json:"code"`
	Message  string `json:"msg"`
	Moreinfo string `json:"moreinfo"`
}

var (
	Err404           = &ControllerError{404, "页面没有找到", ""}
	ErrInputData     = &ControllerError{10001, "数据输入错误", ""}
	ErrDatabase      = &ControllerError{10002, "服务器错误", ""}
	ErrValidation    = &ControllerError{13011, "请求参数验证失败", ""}
	ErrAddFail       = &ControllerError{11000, "创建失败", ""}
	ErrEditFail      = &ControllerError{11001, "更新失败", ""}
	ErrDelFail       = &ControllerError{11002, "删除失败", ""}
	ErrInvalidParams = &ControllerError{11003, "验证失败", ""}
	ErrIdData        = &ControllerError{10016, "此ID无数据记录", ""}
	OtherTaskRunning = &ControllerError{12000, "有其他任务在执行", ""}
	ErrUploadCover   = &ControllerError{12001, "上传封面失败", ""}
)

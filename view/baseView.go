package view

type BaseView struct {
	// 返回码   正常： 0  异常： -1
	RetCode int32 `json:"retCode"`

	// 返回数据
	Data interface{} `json:"data"`

	// 返回提示信息
	Message string `json:"message"`
}

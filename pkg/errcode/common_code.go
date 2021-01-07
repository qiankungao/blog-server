package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFund       = NewError(10000002, "没找到")

	TooManyRequests = NewError(10000007, "请求过多")
)

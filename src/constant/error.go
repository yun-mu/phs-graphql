package constant

import "errors"

const (
	ErrorMsgParamWrong = "param wrong"
	ErrorMsgUserCreate = "创建用户错误"
	ErrorMsgUnAuth     = "un auth"
)

var (
	ErrorOutOfRange    = errors.New("out of range")
	ErrorIDFormatWrong = errors.New("id format is wrong")
	ErrorNotFound      = errors.New("not found")
	ErrorHasExist      = errors.New("has exist")
	ErrorNotExist      = errors.New("not exist")
	ErrorParamWrong    = errors.New("param is wrong")
	ErrorUnAuth        = errors.New("un auth")
	ErrorEmpty         = errors.New("empty error")
	ErrorUnFollow      = errors.New("你还没有关注公众号")
	ErrorBadGateway    = errors.New("服务器错误")
)

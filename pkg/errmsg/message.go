package errmsg

type ErrCode uint

const (
	ErrorMsgCantUnmarshal ErrCode = iota + 1
)

var e = map[ErrCode]string{
	ErrorMsgCantUnmarshal: "can't unmarshal",
}

func Code(msg error) ErrCode {
	if msg == nil {
		return 0
	}

	switch msg.Error() {
	case ErrorMsgCantUnmarshal.String():
		return ErrorMsgCantUnmarshal
	default:
		return ErrCode(0)
	}
}

func (ec ErrCode) Error() string {
	return ec.String()
}

func (ec ErrCode) String() string {
	return e[ec]
}

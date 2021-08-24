package status

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(status int, msg string) Status {
	return Status{
		Code: status,
		Msg:  msg,
	}
}

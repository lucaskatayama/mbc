package status

// Status represents a status
type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New creates a status
func New(status int, msg string) Status {
	return Status{
		Code: status,
		Msg:  msg,
	}
}

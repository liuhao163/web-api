package exception

type Exception struct {
	ErrorNo  int64
	ErrorMsg string
}

var PARAM_ILLEGL = Exception{10000, "param_illegl"}

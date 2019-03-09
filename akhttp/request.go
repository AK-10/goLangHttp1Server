package akhttp

type request struct {
	root string
	method string
	headQuery map[string]string
	field map[string]*string
}

func MakeRequest(buf []byte) request {
	return request{}
}


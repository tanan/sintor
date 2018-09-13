package domain

type RequestID int

type Request struct {
	ID     RequestID
	URL    string
	Method Method
	Body   string
	Header map[string]string
}

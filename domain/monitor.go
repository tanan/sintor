package domain

type MonitorID int

type Method int

type Monitor struct {
	ID     MonitorID
	URL    string
	Method Method
	Body   string
	Header map[string]string
}

type MonitorResult struct {
	ID       MonitorID
	Status   int
	Contents string
}

package monitor

import (
	"github.com/tanan/sintor/core/apperr"
	"github.com/tanan/sintor/domain"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Requester struct {
	URL    string
	Method string
	Header map[string]interface{}
	Body   string
}

func (r Requester) get(url string) (domain.MonitorResult, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println()
		return domain.MonitorResult{}, apperr.AppError{
			Msg: "response error : " + url,
			Err: err,
		}
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 1024))
	if err != nil {
		return domain.MonitorResult{}, apperr.AppError{
			Msg: "body parse error : " + url,
			Err: err,
		}
	}
	return domain.MonitorResult{
		Response: domain.Response{
			Status:       res.StatusCode,
			ResponseTime: 1,
			Contents:     string(body),
		},
	}, nil
}

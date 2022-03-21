package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RemoteHandler struct {
}

type Remote struct {
	Url string `json:"url"`
}

func (rh *RemoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	remote := Remote{}
	json.Unmarshal(b, &remote)
	resp, err := http.Get(remote.Url)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	log.Printf("http get url : %+v, code: %d", remote.Url, resp.StatusCode)
	w.WriteHeader(resp.StatusCode)
}

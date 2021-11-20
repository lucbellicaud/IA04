package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	rad "gitlab.utc.fr/lagruesy/restagentdemo"
)

type RestClientAgent struct {
	id   string
	url  string
	pref []rad.AgentID
}

func NewRestClientAgent(id string, url string, pref []rad.AgentID) *RestClientAgent {
	return &RestClientAgent{id, url, pref}
}

func (rca *RestClientAgent) treatResponse(r *http.Response) rad.AgentID {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp rad.Response
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.Result
}

func (rca *RestClientAgent) doRequest() (res rad.AgentID, err error) {
	req := rad.Request{
		Vote: rca.pref[0],
	}

	// sérialisation de la requête
	url := rca.url + "/calculator"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	res = rca.treatResponse(resp)

	return
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)
	res, err := rca.doRequest()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		log.Printf("Vote : [%s] Résultat : %s\n", rca.pref, res)
	}
}

package restagentdemo

type AgentID string

type Request struct {
	Vote AgentID `json:"vote"`
}

type Response struct {
	Result AgentID `json:"res"`
}

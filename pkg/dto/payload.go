package dto

type NewJobRequest struct {
	Payload string `json:"payload"`
}

type NewJobResponse struct {
	ID string `json:"id"`
}

type JobResultResponse struct {
	Result string `json:"result"`
}

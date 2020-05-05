package model

import "financial-planner/constanta"

// Status for default structure code response and message
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response for default structure response message
type Response struct {
	Status Status                 `json:"status"`
	Result map[string]interface{} `json:"result"`
}

// SetCode : setter code for Response
func (r *Response) SetCode(code int) {
	r.Status.Code = code
	if code == 0 {
		r.Status.Code = constanta.SuccessCode
	}
}

// SetMessage : setter message for Response
func (r *Response) SetMessage(message string) {
	r.Status.Message = message
	if message == "" {
		r.Status.Message = constanta.SuccessMessage
	}
}

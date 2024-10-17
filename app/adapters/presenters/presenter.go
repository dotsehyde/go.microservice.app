package presenters

type (
	ErrorResponse struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

	PaginatedResponse struct {
		Data          interface{} `json:"data"`
		TotalElements int         `json:"total_elements"`
		HasMore       bool        `json:"has_more"`
	}
)

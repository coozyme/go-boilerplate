package response

type (
	PaginationInfo struct {
		Page        int `json:"page"`
		TotalPage   int `json:"total_page"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
	}

	Pagination struct {
		PaginationInfo
		Count int64 `json:"count"`
		Limit int64 `json:"limit"`
	}
)

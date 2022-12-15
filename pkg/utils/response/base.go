package response

type (
	Status struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
	}

	Meta struct {
		Pagination *Pagination `json:"pagination,omitempty"`
	}
)

var (
	Duplicate         = "duplicated"
	DataNotFound      = "data not found"
	Unauthorized      = "unauthorized"
	Validation        = "invalid parameters or payload"
	BadRequest        = "bad Request"
	InternalServer    = "internal server error"
	RouteNotFound     = "route not found"
	Server            = "something bad happened"
	OtpExpired        = "otp expired"
	OtpInvalid        = "otp invalid"
	UserAlreadyExists = "user already exists"
)

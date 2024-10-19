package requestdto

type (
	TestRequest struct {
		Name string `json:"name" validate:"required|ascii"`
		Age  int    `json:"age" validate:"required|int"`
	}
)

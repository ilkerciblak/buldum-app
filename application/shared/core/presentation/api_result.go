package presentation

// API success durumunda da bir status donebilir ama simdilik bosverelim
type ApiResult[T any] struct {
	Data           T
	ProblemDetails ProblemDetails
}

type ProblemDetails struct {
	Type   string            `json:"type"`
	Title  string            `json:"title"`
	Status int               `json:"status"`
	Detail string            `json:"detail"`
	Errors map[string]string `json:"errors"`
}

package response

type Response struct {
	StatusCode    int
	Path          string
	ContentLength int64
	LastModified  string
}

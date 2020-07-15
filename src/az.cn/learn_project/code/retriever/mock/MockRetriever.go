package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, data map[string]string) string {
	r.Contents = data["contents"]
	return "ok"
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Inspecting Retriever: %s", r.Contents)
}

func (r *Retriever) Error() string {
	return r.Contents
}

package taskService

type Task struct {
	ID          int    `json:"id"`
	Is_Done     bool   `json:"is_Done"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

/*type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
*/

package response

type CreateResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

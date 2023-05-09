package request

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	NamespaceID string `json:"namespace_id"`
	Readme      bool   `json:"initialize_with_readme"`
}

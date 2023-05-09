package response

type CreateCommit struct {
	AuthorEmail    string `json:"author_email"`
	AuthorName     string `json:"author_name"`
	AuthoredDate   string `json:"authored_date"`
	CommittedDate  string `json:"committed_date"`
	CommitterEmail string `json:"committer_email"`
	CommitterName  string `json:"committer_name"`
	ID             string `json:"id"`
	ShortID        string `json:"short_id"`
	Title          string `json:"title"`
	Message        string `json:"message"`
}

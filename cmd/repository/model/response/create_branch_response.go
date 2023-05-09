package response

type CreateBranch struct {
	Name      string       `json:"name"`
	Merged    bool         `json:"merged"`
	Protected string       `json:"protected"`
	Default   string       `json:"default"`
	WebURL    string       `json:"web_url"`
	Commit    CreateCommit `json:"commit"`
}

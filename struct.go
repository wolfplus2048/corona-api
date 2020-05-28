package corona

// Server struct
type Server struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Metadata map[string]string `json:"metadata"`
	Frontend bool              `json:"frontend"`
	Hostname string            `json:"hostname"`
}


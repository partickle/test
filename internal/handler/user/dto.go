package user

type SetIsActiveResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"isActive"`
}

type GetReviewResponse struct {
	UserID       string        `json:"user_id"`
	PullRequests []PullRequest `json:"pull_requests"`
}

type PullRequest struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}

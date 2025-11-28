package pr

import "time"

type PullRequest struct {
	PullRequestID     string
	PullRequestName   string
	AuthorID          string
	Status            PullRequestStatus
	AssignedReviewers []string
	CreatedAt         time.Time
	MergedAt          *time.Time
}

type PullRequestShort struct {
	PullRequestID   string
	PullRequestName string
	AuthorID        string
	Status          PullRequestStatus
}

type PullRequestStatus string

const (
	StatusOpen   = "open"
	StatusMerged = "merged"
)

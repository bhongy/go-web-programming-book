package data

import "time"

// Thread models a discussion thread for the forum
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}

// NumReplies returns the number of posts in a thread
func (thread *Thread) NumReplies() int {
	return 0
}

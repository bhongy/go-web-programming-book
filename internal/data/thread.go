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

// Creator returns the User who created the thread
func (t *Thread) Creator() (u User) {
	query := `
		SELECT
			id, uuid, name, email, created_at
		FROM
			users
		WHERE
			id = $1
	`
	Db.
		QueryRow(query, t.UserID).
		Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	return
}

// CreatedAtDate formats CreateAt data to display nicely on the screen
func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// NumReplies returns the number of posts in a thread
func (t *Thread) NumReplies() (count int) {
	query := "SELECT count(id) FROM posts WHERE thread_id = $1"
	Db.QueryRow(query, t.ID).Scan(&count)
	return
}

// Posts returns the posts in the thread
func (t *Thread) Posts() (posts []Post, err error) {
	query := `
		SELECT
			id, uuid, body, user_id, thread_id, created_at
		FROM
			posts
		WHERE
			thread_id = $1
	`
	rows, err := Db.Query(query, t.ID)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var p Post
		err = rows.Scan(
			&p.ID, &p.UUID, &p.Body, &p.UserID, &p.ThreadID, &p.CreatedAt,
		)
		if err != nil {
			return
		}
		posts = append(posts, p)
	}
	return
}

// Threads returns all threads in the database
// order by the most recently created ones first
func Threads() (threads []Thread, err error) {
	query := `
		SELECT
			id, uuid, topic, user_id, created_at
		FROM
			threads
		ORDER BY
			created_at DESC
	`
	rows, err := Db.Query(query)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var t Thread
		err = rows.Scan(
			&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt,
		)
		if err != nil {
			return
		}
		threads = append(threads, t)
	}
	return
}

// ThreadByUUID returns a thread given UUID
func ThreadByUUID(uuid string) (t Thread, err error) {
	query := `
		SELECT
			id, uuid, topic, user_id, created_at
		FROM
			threads
		WHERE
			uuid = $1
	`
	err = Db.
		QueryRow(query, uuid).
		Scan(&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt)
	return
}

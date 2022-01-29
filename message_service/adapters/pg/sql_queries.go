package pg

const (
	createEmail = `INSERT INTO emails (uuid, from, to, subject, body, created_at, updated_at) 
									VALUES ($1, $2, $3, $4, $5, now(), now())`
	getEmailByUUID = `SELECT uuid, 
							from, to, subject, body, created_at, updated_at 
						FROM emails
							WHERE uuid = $1`
)

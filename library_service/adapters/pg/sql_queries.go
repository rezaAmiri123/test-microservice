package pg

const (
	createArticle = `INSERT INTO articles (uuid, user_uuid, title, slug, description, body, created_at, updated_at) 
									VALUES ($1, $2, $3, $4, $5, $6, now(), now()) 
									RETURNING *`

	//getNewsByID = `SELECT n.news_id,
	//								n.title,
	//								n.content,
	//								n.updated_at,
	//								n.image_url,
	//								n.category,
	//								CONCAT(u.first_name, ' ', u.last_name) as author,
	//								u.user_id as author_id
	//						 FROM news n
	//								  LEFT JOIN users u on u.user_id = n.author_id
	//						 WHERE news_id = $1`

	getArticleBySlug = `SELECT uuid, 
							user_uuid, title, slug, description, body, created_at, updated_at 
						FROM articles
							WHERE slug = $1`
)

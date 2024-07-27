package queries

var GetGeekTodayNewsByDate = `
	SELECT uuid, post, descLink, link AS originalLink, founded
	FROM "Geek" g
	WHERE DATE(founded) = $1
	ORDER BY rank DESC
	OFFSET $2
	LIMIT $3
`


var GetGeekNewsByDate = `
	SELECT uuid, post, descLink, link AS originalLink, founded
	FROM "Geek" g
	WHERE founded IS BETWEEN $1 AND $2
	ORDER BY rank DESC
	OFFSET $3
	LIMIT $4
`

var LikeGeekNews = `
	INSERT INTO liked_news
	SET
		userUuid = $1,
		postUuid = $2,
		platform = "GEEK"
`
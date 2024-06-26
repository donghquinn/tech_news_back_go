package queries

var GetGeekTodayNewsByDate = `
	SELECT *
	FROM Geek
	WHERE DATE(founded) = $1
	OFFSET $2
	LIMIT $3
	ORDER BY rank DESC
	`


var GetGeekNewsByDate = `
	SELECT *
	FROM Geek
	WHERE founded IS BETWEEN $1 AND $2
	OFFSET $3
	LIMIT $4
	ORDER BY rank DESC
	`

var LikeGeekNews = `
	INSERT INTO liked_news
	SET
		userUuid = $1,
		postUuid = $2,
		platform = "GEEK"
`
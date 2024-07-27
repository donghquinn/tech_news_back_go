package queries

var GetTodayHackerByDate =  `
	SELECT *
	FROM Hackers
	WHERE DATE(founded) = $1
	ORDER BY rank DESC
	OFFSET $2
	LIMIT $3
`

var GetHackerByDate =  `
	SELECT *
	FROM Hackers
	WHERE founded IS BETWEEN $1 AND $2
	ORDER BY rank DESC
	OFFSET $3
	LIMIT $4
`

var LikeHackerNews = `
	INSERT INTO liked_news
	SET
		user_uuid = $1,
		post_uuid = $2,
		platform = "HACKERS"
`
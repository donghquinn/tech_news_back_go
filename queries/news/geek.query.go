package queries

var GetGeekNewsByDate = `
	SELECT *
	FROM Geek
	WHERE founded IS BETWEEN $1 AND $2
	OFFSET $3
	LIMIT $4
	ORDER BY rank DESC
	`

var LikeGeekNews = `
	INSERT INTO Hacker_Liked
	SET
		userUuid = $1,
		postUuid = $2,
		newsPlatform = "GEEK"
`
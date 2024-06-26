package queries


var GetTodayMlByDate =  `
	SELECT *
	FROM MachineNews
	WHERE DATE(founded) = $1
	OFFSET $2
	LIMIT $3
	ORDER BY founded DESC
	`

var GetMlByDate =  `
	SELECT *
	FROM MachineNews
	WHERE founded IS BETWEEN $1 AND $2
	OFFSET $3
	LIMIT $4
	ORDER BY founded DESC
	`


var LikeMlNews = `
	INSERT INTO liked_news
	SET
		userUuid = $1,
		postUuid = $2,
		platform = "ML"
`
package queries


var GetTodayMlByDate =  `
	SELECT *
	FROM MachineNews
	WHERE DATE(founded) = $1
	ORDER BY founded DESC
	OFFSET $2
	LIMIT $3
`

var GetMlByDate =  `
	SELECT *
	FROM MachineNews
	WHERE founded IS BETWEEN $1 AND $2
	ORDER BY founded DESC
	OFFSET $3
	LIMIT $4
`


var LikeMlNews = `
	INSERT INTO liked_news
	SET
		userUuid = $1,
		postUuid = $2,
		platform = "ML"
`
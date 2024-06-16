package queries

var GetMlByDate =  `
	SELECT *
	FROM MachineNews
	WHERE founded IS BETWEEN $1 AND $2
	OFFSET $3
	LIMIT $4
	ORDER BY founded DESC
	`

var LikeMlNews = `
	INSERT INTO Hacker_Liked
	SET
		userUuid = $1,
		postUuid = $2,
		newsPlatform = "ML"
`
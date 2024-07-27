package queries

// Hacker Liked News
var GetLikedHackerNewsByUserId  = `
	SELECT h.*
	FROM liked_news ln
	LEFT JOIN "Hackers" h ON h.uuid = ln.postUuid
	WHERE ln.userUuid = $1
		AND ln.platform = "HACKERS"
	OFFSET $2
	LIMIT $3
	ORDER BY h.founded DESC
`

var GetLikedMlNewsByUserId  = `
	SELECT h.*
	FROM liked_news ln
	LEFT JOIN "MachineNews" m ON m.uuid = ln.postUuid
	WHERE ln.userUuid = $1
		AND ln.platform = "ML"
	OFFSET $2
	LIMIT $3
	ORDER BY m.founded DESC
`

var GetLikedGeekNewsByUserId  = `
	SELECT h.*
	FROM liked_news ln
	LEFT JOIN "Geek" g ON g.uuid = ln.postUuid
	WHERE ln.userUuid = $1
		AND ln.platform = "GEEK"
	OFFSET $2
	LIMIT $3
	ORDER BY g.founded DESC
`

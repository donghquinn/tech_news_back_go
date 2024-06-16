package queries

// Hacker Liked News
var GetLikedHackerNewsByUserId  = `
	SELECT h.*
	FROM Hacker_Liked hl
	LEFT JOIN Hackers h ON h.uuid = hl.postUuid
	WHERE hl.userUuid = $1
	ORDER BY h.founded DESC
`

var GetLikedMlNewsByUserId  = `
	SELECT h.*
	FROM Ml_Liked ml
	LEFT JOIN MachineNews m ON m.uuid = ml.postUuid
	WHERE ml.userUuid = $1
	ORDER BY m.founded DESC
`

var GetLikedGeekNewsByUserId  = `
	SELECT h.*
	FROM Geek_Liked gl
	LEFT JOIN Geek g ON g.uuid = gl.postUuid
	WHERE gl.userUuid = $1
	ORDER BY g.founded DESC
`

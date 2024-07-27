package queries


var GetUserData = `
	SELECT c.uuid, c.email, c.password, c.user_status
	FROM "Client" c
	WHERE c.email = $1
`

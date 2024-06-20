package queries


var GetUserData = `
	SELECT uuid, email, password, user_status
	FROM Client
	WHERE email = $1
`

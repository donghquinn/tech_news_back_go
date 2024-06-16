package queries



var GetUserData = `
	SELECT *
	FROM Client
	WHERE email = $1
`

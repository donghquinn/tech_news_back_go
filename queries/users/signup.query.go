package queries

var InsertNewUser = `
	INSERT INTO Client
	SET
		email = $1,
		name = $2,
		password = $3
`

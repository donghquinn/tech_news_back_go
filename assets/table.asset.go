package assets

var UserExternalUuidFunctions = `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
`


var CreateClient = `
	CRAETE TABLE IF NOT EXISTS Client (
	
	)
`

var QueriesTransaction = []string{
	UserExternalUuidFunctions, CreateClient}
	//  CreateSession, ChatGen, ImageGen, CreateFile}
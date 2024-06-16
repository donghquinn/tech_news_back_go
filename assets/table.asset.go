package assets

var UserExternalUuidFunctions = `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
`


var CreateClient = `
	CRAETE TABLE IF NOT EXISTS Client (
		uuid	UUID			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		email   VARCHAR(150)	NOT NULL	UNIQUE,
		name	VARCHAR(50)		NOT NULL,
		password	VARCHAR(200)	NOT NULL,
		password_token	VARCHAR(100),
		is_logined	BOOLEAN NOT NULL DEFAULT FALSE,
		user_status	TINYINT(1)	DEFAULT 1,
		signed_in	timestamp NOT NULL DEFAULT NOW(),
		logined		timestamp,

		CONSTRAINT user_idx UNIQUE (user_status)
	)
`

var CreateSession = `
	CREATE TABLE IF NOT EXISTS session (
		session_id		UUID 			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		user_id 		VARCHAR(50) 	NOT NULL,
		user_ip			VARCHAR(50),
		created 		TIMESTAMP 		NOT NULL DEFAULT NOW(),

		CONSTRAINT session_idx UNIQUE (user_id)
	);
`

var CreateLikedNews = `
	RAETE TABLE IF NOT EXISTS liked_news (
		hl_seq	SERIAL	NOT NULL PRIMARY KEY,
		user_uuid	VARCHAR(50) NOT NULL,
		post_uuid	VARCHAR(50) NOT NULL,
		platform 	ENUM("HACKERS","GEEK","ML") NOT NULL,
		created 	TIMESTAMP DEFAULT NOW()
		
		CONSTRAINT news_liked_idx UNIQUE (user_uuid, platform)
	)
`

var CreateHackerNews = `
	CRAETE TABLE IF NOT EXISTS Hackers (
		uuid	UUID			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		user_uuid	VARCHAR(50) NOT NULL,
		post_uuid	VARCHAR(50) NOT NULL,
		platform 	ENUM("HACKERS","GEEK","ML") NOT NULL,
		created 	TIMESTAMP DEFAULT NOW()
		
		CONSTRAINT news_liked_idx UNIQUE (user_uuid, platform)
	)
`


var QueriesTransaction = []string{
	UserExternalUuidFunctions, CreateClient, CreateLikedNews}
	//  CreateSession, ChatGen, ImageGen, CreateFile}
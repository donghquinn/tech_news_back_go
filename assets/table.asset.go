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

		CONSTRAINT user_idx (user_status)
	)
`

var CreateSession = `
	CREATE TABLE IF NOT EXISTS session (
		session_id		UUID 			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		user_id 		VARCHAR(50) 	NOT NULL,
		user_ip			VARCHAR(50),
		created 		TIMESTAMP 		NOT NULL DEFAULT NOW(),

		CONSTRAINT session_idx (user_id)
	);
`

var CreateLikedGeekNews = `
	RAETE TABLE IF NOT EXISTS Geek_Liked (
		hl_seq	SERIAL	NOT NULL PRIMARY KEY,
		user_uuid	VARCHAR(50) NOT NULL,
		post_uuid	VARCHAR(50) NOT NULL,
		platform 	ENUM("HACKERS","GEEK","ML") NOT NULL,
		created 	TIMESTAMP DEFAULT NOW()
		
		CONSTRAINT news_liked_idx (user_uuid, platform)
	)
`

var CreateLikedHackerNews = `
	RAETE TABLE IF NOT EXISTS Hacker_Liked (
		hl_seq	SERIAL	NOT NULL PRIMARY KEY,
		user_uuid	VARCHAR(50) NOT NULL,
		post_uuid	VARCHAR(50) NOT NULL,
		platform 	ENUM("HACKERS","GEEK","ML") NOT NULL,
		created 	TIMESTAMP DEFAULT NOW()
		
		CONSTRAINT news_liked_idx (user_uuid, platform)
	)
`

var CreateHackerNews = `
	CRAETE TABLE IF NOT EXISTS Hackers (
		uuid		UUID			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		rank		TINYINT(2)		NOT NULL,
		post		TEXT			NOT NULL,
		link		TEXT			NOT NULL,
		founded 	TIMESTAMP		DEFAULT NOW()
	)
`

var CreateGeekNews = `
	CRAETE TABLE IF NOT EXISTS Geek (
		uuid		UUID			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		page		VARCHAR(10)		NOT NULL,
		rank		TINYINT(2)		NOT NULL,
		post		TEXT			NOT NULL,
		descLink	TEXT			NOT NULL,
		link		TEXT			NOT NULL,
		founded 	TIMESTAMP		DEFAULT NOW()
	)
`

var CreateMachineNews = `
	CRAETE TABLE IF NOT EXISTS MachineNews (
		uuid		UUID			NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		category	VARCHAR(10)		NOT NULL,
		title		TEXT			NOT NULL,
		link		TEXT			NOT NULL,
		founded 	TIMESTAMP		DEFAULT NOW()
	)
`

var QueriesTransaction = []string{
	UserExternalUuidFunctions, CreateClient, CreateLikedHackerNews}
	//  CreateSession, ChatGen, ImageGen, CreateFile}
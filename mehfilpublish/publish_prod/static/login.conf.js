var pageConf = {
	"AuthPage":"http://localhost:7070/login",
	"LocalAuthServer":"http://localhost:7070/auth/login",
	"ViewsServer": "http://localhost:7070/view",
	"LinksServer": "http://localhost:7070/links",
	"SuccessRedirect":"/home",
	"AuthRequired":false,
	"AuthToken":"Auth-Token",
	"entities":{
	},
	"actions": {
		"Home": {
			"url":"/",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"home.html"
		}
	}
	
};
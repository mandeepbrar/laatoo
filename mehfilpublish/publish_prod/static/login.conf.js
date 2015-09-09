var pageConf = {
	"AuthPage":"/login",
	"LocalAuthServer":"/auth/login",
	"ViewsServer": "/view",
	"LinksServer": "/links",
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
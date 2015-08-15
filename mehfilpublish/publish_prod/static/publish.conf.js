var pageConf = {
	"LocalAuthServer":"http://localhost:7070/auth/login",
	"ViewsServer": "http://localhost:7070/view",
	"LinksServer": "http://localhost:7070/links",
	"SuccessRedirect":"/home",
	"AuthToken":"Auth-Token",
	"entities":{
		"article": {
			"url": "http://localhost:7070/article",
			"tabbed": true,
			"model" : {},
			"options": {},
			"tabs": [
				{
					"label":"Gurmukhi",
					"fields": [
						{
							"key":"Title",
							"type":"input",
							"templateOptions": {
					        		"label": "Title"
							}
						},
						{
							"key":"BodyGur",
							"type":"ckeditor",
							"templateOptions": {
					        		"label": "Body "
							}
						},
						{
							"key":"SummaryGur",
							"type":"textarea",
							"templateOptions": {
					        		"label": "Summary "
							}
						}
					]					
				},
				{
					"label":"English",
					"fields": [
						{
							"key":"TitleEng",
							"type":"input",
							"templateOptions": {
					        		"label": "Title"
							}
						},
						{
							"key":"BodyEng",
							"type":"ckeditor",
							"templateOptions": {
					        		"label": "Body ",
							},
						},
						{
							"key":"SummaryEng",
							"type":"textarea",
							"templateOptions": {
					        		"label": "Summary "
							}
						}
					]					
				},
				{
					"label":"Others",
					"fields": [
						{
							"key":"Image",
							"type":"input",
							"templateOptions": {
					        		"label": "Image "
							}
						}
					]
				}
			]
		}		
	},
	"actions": {
		"Home": {
			"url":"/",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"home.html"
		},
		"View Articles": {
			"url":"/articles",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"articleslist.html"
		},
		"View Article": {
			"url":"/article/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"viewarticle.html"
		},
		"Edit Article": {
			"url":"/article/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"articleedit.html"
		},
		"Create Article": {
			"url":"/article/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createarticle.html"
		}		
	},
	"actionset": {
		"menu":{
			"Home": {
				"action":"Home",
				"label":"Home"
			},
			"Articles": {
				"action":"View Articles",
				"label":"Articles"
			}
		}
	}
	
};
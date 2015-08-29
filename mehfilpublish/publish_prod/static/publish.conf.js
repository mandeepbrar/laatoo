var pageConf = {
	"AuthPage":"http://localhost:7070/login",
	"LocalAuthServer":"http://localhost:7070/auth/login",
	"ViewsServer": "/view",
	"LinksServer": "http://localhost:7070/links",
	"SuccessRedirect":"/home",
	"AuthRequired":true,
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
							"type":"media",
							"templateOptions": {
					        		"label": "Image "
							}
						},
						{
							"key":"Type",
							"type":"ui-grid",
							"templateOptions": {
					        		"label": "Article Type",
								"griditems":[
									{
										"Type":"Boli"
									},
									{
										"Type":"Lekh"
									}
								],
								"columns":[
									{"name":"Type"}
								]
							}
						}
					]
				}
			]
		},		
		"mehfil": {
			"url": "http://localhost:7070/mehfil",
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
							"type":"media",
							"templateOptions": {
					        		"label": "Image",
								"mediatype":"image"									
							}
						}
					]
				}
			]	
		},
		"word": {
			"url": "http://localhost:7070/word",
			"model" : {},
			"options": {},
			"fields": [
				{
					"key":"WordGur",
					"type":"input",
					"templateOptions": {
			        		"label": "Gurmukhi",
					}
				},
				{
					"key":"WordEng",
					"type":"input",
					"templateOptions": {
			        		"label": "English",
					}
				}
				
			]
		},
		"default_user": {
			"url": "http://localhost:7070/user",
			"model" : {},
			"options": {},
			"fields": [
				{
					"key":"Id",
					"type":"input",
					"templateOptions": {
			        		"label": "User Id",
					}
				},
				{
					"key":"Password",
					"type":"input",
					"templateOptions": {
				        "type": "password",
			        		"label": "Password",
					}
				},
				{
					"key":"Roles",
					"type":"ui-grid",
					"templateOptions": {
			        		"label": "User Roles",
						"griditems":[
							{
								"Roles":"Admin"
							},
							{
								"Roles":"Editor"
							}
						],
						"columns":[
							{"name":"Roles"}
						]
					}
				}
				
			]
		},
		"default_role": {
			"url": "http://localhost:7070/role",
			"model" : {},
			"options": {},
			"fields": [
				{
					"key":"Role",
					"type":"input",
					"templateOptions": {
			        		"label": "Role",
					}
				},
				{
					"key":"Permissions",
					"type":"ui-grid",
					"templateOptions": {
			        		"label": "Permissions",
						"griditems":[
							{
								"Permissions":"Edit User"
							},
							{
								"Permissions":"Edit Role"
							}
						],
						"columns":[
							{"name":"Permissions"}
						]
					}
				}
				
			]
		},
		"video": {
			"url": "http://localhost:7070/video",
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
							"key":"Video",
							"type":"media",
							"templateOptions": {
					        		"label": "Video",
								"mediatype":"video"									
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
		},		
		"View Mehfils": {
			"url":"/mehfils",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"mehfilslist.html"
		},
		"View Mehfil": {
			"url":"/mehfil/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"viewmehfil.html"
		},
		"Edit Mehfil": {
			"url":"/mehfil/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"mehfiledit.html"
		},
		"Create Mehfil": {
			"url":"/mehfil/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createmehfil.html"
		},		
		"View Video": {
			"url":"/video/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"viewvideo.html"
		},
		"Edit Video": {
			"url":"/video/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"videoedit.html"
		},
		"Create Video": {
			"url":"/video/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createvideo.html"
		},		
		"View Videos": {
			"url":"/videos",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"videoslist.html"
		},
		"Edit User": {
			"url":"/user/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"useredit.html"
		},
		"Create User": {
			"url":"/user/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createuser.html"
		},		
		"View Users": {
			"url":"/users",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"userslist.html"
		},
		"Edit Role": {
			"url":"/role/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"roleedit.html"
		},
		"Create Role": {
			"url":"/role/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createrole.html"
		},		
		"View Roles": {
			"url":"/roles",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"roleslist.html"
		},
		"View Words": {
			"url":"/words",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"words.html"
		},
		"New Words": {
			"url":"/newwords",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"newwords.html"
		},
		"Review Words": {
			"url":"/reviewwords",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"reviewwords.html"
		},
		"Edit Word": {
			"url":"/word/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"wordedit.html"
		},		
		"Create Word": {
			"url":"/word/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"templatepath":"createword.html"
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
			},
			"Mehfils": {
				"action":"View Mehfils",
				"label":"Mehfils"
			},
			"Videos": {
				"action":"View Videos",
				"label":"Videos"
			},
			"Words": {
				"action":"View Words",
				"label":"Words"
			},
			"Roles": {
				"action":"View Roles",
				"label":"Roles"
			},
			"Users": {
				"action":"View Users",
				"label":"Users"
			}
		}
	}
	
};
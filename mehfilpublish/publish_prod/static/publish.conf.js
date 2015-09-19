var pageConf = {
	"AuthPage":"/login",
	"LocalAuthServer":"/auth/login",
	"ViewsServer": "/view",
	"LinksServer": "/links",
	"SuccessRedirect":"/home",
	"AuthRequired":true,
	"AuthToken":"Auth-Token",
	"entities":{
		"Article": {
			"url": "/article",
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
									{"name":"Type", "key":"Type"}
								],
								"valueField":"Type"
							}
						}
					]
				}
			]
		},		
		"Mehfil": {
			"url": "/mehfil",
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
		"Word": {
			"url": "/word",
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
			"url": "/user",
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
						"gridcallback": "/configviews?viewname=view_entities&entity=Role",
						"columns":[
							{"name":"Role", "key": "Role"}
						],
						"valueField":"Role"
					}
				}
				
			]
		},
		"default_role": {
			"url": "/role",
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
						"gridcallback": "/permissions",
						"columns":[
							{"name":"Permissions", "key": "="}
						],
						"valueField":"="
					}
				}
				
			]
		},
		"Video": {
			"url": "/video",
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
			"permission":"none",
			"templatepath":"home.html"
		},
		"View Articles": {
			"url":"/articles",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Article",
			"templatepath":"articleslist.html"
		},
		"View Article": {
			"url":"/article/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Article",
			"templatepath":"viewarticle.html"
		},
		"Edit Article": {
			"url":"/article/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Article",
			"templatepath":"articleedit.html"
		},
		"Delete Article": {
			"actiontype":"restcall",
			"url":"/article/{id}",
			"restmethod":"DELETE",
			"permission":"Delete Article"
		},		
		"Create Article": {
			"url":"/article/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create Article",
			"templatepath":"createarticle.html"
		},		
		"View Mehfils": {
			"url":"/mehfils",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Mehfil",
			"templatepath":"mehfilslist.html"
		},
		"View Mehfil": {
			"url":"/mehfil/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Mehfil",
			"templatepath":"viewmehfil.html"
		},
		"Edit Mehfil": {
			"url":"/mehfil/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Mehfil",
			"templatepath":"mehfiledit.html"
		},
		"Delete Mehfil": {
			"actiontype":"restcall",
			"url":"/mehfil/{id}",
			"restmethod":"DELETE",
			"permission":"Delete Mehfil"
		},		
		"Create Mehfil": {
			"url":"/mehfil/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create Mehfil",
			"templatepath":"createmehfil.html"
		},		
		"View Video": {
			"url":"/video/view/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Video",
			"templatepath":"viewvideo.html"
		},
		"Edit Video": {
			"url":"/video/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Video",
			"templatepath":"videoedit.html"
		},
		"Delete Video": {
			"actiontype":"restcall",
			"url":"/video/{id}",
			"restmethod":"DELETE",
			"permission":"Delete Video"
		},		
		"Create Video": {
			"url":"/video/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create Video",
			"templatepath":"createvideo.html"
		},		
		"View Videos": {
			"url":"/videos",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Video",
			"templatepath":"videoslist.html"
		},
		"Edit User": {
			"url":"/user/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit User",
			"templatepath":"useredit.html"
		},
		"Create User": {
			"url":"/user/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create User",
			"templatepath":"createuser.html"
		},		
		"Delete User": {
			"actiontype":"restcall",
			"url":"/user/{id}",
			"restmethod":"DELETE",
			"permission":"Delete User"
		},		
		"View Users": {
			"url":"/users",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View User",
			"templatepath":"userslist.html"
		},
		"Edit Role": {
			"url":"/role/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Role",
			"templatepath":"roleedit.html"
		},
		"Create Role": {
			"url":"/role/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create Role",
			"templatepath":"createrole.html"
		},		
		"Delete Role": {
			"actiontype":"restcall",
			"url":"/role/{id}",
			"restmethod":"DELETE",
			"permission":"Delete Role"
		},		
		"View Roles": {
			"url":"/roles",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Role",
			"templatepath":"roleslist.html"
		},
		"View Words": {
			"url":"/words",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"View Word",
			"templatepath":"words.html"
		},
		"New Words": {
			"url":"/newwords",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Word",
			"templatepath":"newwords.html"
		},
		"Review Words": {
			"url":"/reviewwords",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Word",
			"templatepath":"reviewwords.html"
		},
		"Edit Word": {
			"url":"/word/edit/{id}",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Edit Word",
			"templatepath":"wordedit.html"
		},		
		"Delete Word": {
			"actiontype":"restcall",
			"url":"/word/{id}",
			"restmethod":"DELETE",
			"permission":"Delete Word"
		},		
		"Create Word": {
			"url":"/word/create",
			"actiontype":"openpartialpage",
			"viewmode":"link",
			"view":"mainview",
			"permission":"Create Word",
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
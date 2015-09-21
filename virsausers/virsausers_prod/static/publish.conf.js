var pageConf = {
	"AuthPage":"/login",
	"LocalAuthServer":"/auth/login",
	"ViewsServer": "/view",
	"LinksServer": "/links",
	"SuccessRedirect":"/home",
	"AuthRequired":true,
	"AuthToken":"Auth-Token",
	"entities":{
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
						"gridcallback": "/view?viewname=view_entities&entity=Role",
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
		}
	},
	"actionset": {
		"menu":{
			"Home": {
				"action":"Home",
				"label":"Home"
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
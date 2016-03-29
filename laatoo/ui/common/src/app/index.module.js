(function() {

    'use strict';
    var mainApp = angular.module('main', ['ngAnimate', 'ui.bootstrap.tabs', 'ui.router.tabs', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'dialogs.main', 'pascalprecht.translate', 'dialogs.default-translations', 'login', 'view', 'actions', 'entity', 'media', 'formlymedia', 'formlyckeditor', 'angularFileUpload', 'smart-table', 'uigrid']);

    

	if(window.AppName == "main") {
		window.bootstrapApplication(["main"]);
	}
	

    
})();

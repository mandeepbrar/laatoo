(function() {
  'use strict';

	var formlyckeditor = angular.module('formlyckeditor', ['ui.bootstrap', 'formly', 'formlyBootstrap', 'ckeditor']);

	formlyckeditor.config(formlyConfig);

	/** @ngInject */
	function formlyConfig(formlyConfigProvider) {
		formlyConfigProvider.setType({
			name: 'ckeditor',
				extends: 'textarea',
			});
			formlyConfigProvider.templateManipulators.preWrapper.push(function(template, options, scope) {
				if(options.type == "ckeditor") {
					options.ngModelAttrs= { ckeditor: { attribute: 'ckeditor'}};
					if(!options.templateOptions) {
						options.templateOptions = {};
					}
					options.templateOptions.ckeditor="";				
			}
			return template;
		});	
	
	}


})();
(function() {
  'use strict';

   var entityapp = angular.module('entity', ['ngResource', 'ui.bootstrap', 'formly', 'formlyBootstrap', 'ckeditor']);

   entityapp.config(config);

  /** @ngInject */
  function config(formlyConfigProvider) {
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
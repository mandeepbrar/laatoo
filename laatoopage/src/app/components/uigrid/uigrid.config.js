(function() {
  'use strict';

  var uigridapp = angular
    .module('uigrid')
    .config(config);

  /** @ngInject */
  function config(formlyConfigProvider) {
	formlyConfigProvider.setType({
	  name: 'ui-grid',
      templateUrl: 'app/components/uigrid/multiselect.view.html',
	  controller: MultiselectCtrl
	});	
	formlyConfigProvider.templateManipulators.preWrapper.push(function(template, options, scope) {
		if(options.type == "ui-grid") {
			scope.haslabel = false;
			scope.field = options.key;
			scope.$watch('model', function (newValue) {
					try {
						scope.selected = newValue[options.key];
						scope.griditems = options.templateOptions.griditems;
						scope.status = [];
						scope.label = options.templateOptions.label;
						scope.columns = options.templateOptions.columns;
						if(options.templateOptions.value) {
							scope.valueField = options.templateOptions.valueField;							
						} else {
							scope.valueField = options.key;
						}
						for(var index in scope.griditems) {
							var item = scope.griditems[index];
							var valToTest = item[scope.valueField];
							scope.status[item[scope.valueField]] = (scope.selected.indexOf(item[scope.valueField])>-1);
						}
					}catch(ex){}
		    });
		}
		return template;
	});	
  }
  
  /** @ngInject */
  function MultiselectCtrl($scope) {
		$scope.field = "";
		if($scope.options.templateOptions.mediatype) {
			$scope.mediatype = $scope.options.templateOptions.mediatype;
		} else {
			$scope.mediatype = 'image';
		}
		console.log("setting up change method");
		$scope.oncheckboxchange = function(evt, val) {
			console.log($scope.model);
			var modelVal = $scope.model[$scope.options.key];
			console.log("checkbox change");
			console.log(modelVal);
			if(!modelVal || !(modelVal instanceof Array)) {
			console.log("initializing");
				modelVal = [];
			}
			console.log(modelVal);
			if(evt.target.checked) {
			console.log("pushing "+val);
				modelVal.push(val);				
			} else{				
				var index = modelVal.indexOf(val);
				if (index > -1) {
				    modelVal.splice(index, 1);
				}
			}				
			$scope.model[$scope.options.key] = modelVal;
			console.log($scope.model);
		};
		console.log($scope.oncheckboxchange);
  }


})();



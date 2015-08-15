(function() {
  'use strict';

  angular
    .module('entity')
    .directive('entityform', entityform);

  /** @ngInject */
  function entityform() {
	var scope;
    var directive = {
      restrict: 'E',
      templateUrl: 'app/components/entity/entityform.view.html',
	  scope: { },
      replace: true,
      controller: EntityformCtrl,
      controllerAs: 'entityform',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function EntityformCtrl($scope, $element, $attrs, $http) {
	  scope = $scope;
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
      var name;
      if($attrs.name) {
        name = $attrs.name;
      } else {
		throw new Error("Server error. Entity name missing");	
	  }
      var entity = window.pageConf.entities[name];
	  $scope.entity = entity;
	  $scope.id = $attrs.id;
      $scope.submitText = "Submit";
      if($attrs.submitText) {
        $scope.submitText = $attrs.submitText;
      }
	  var id = null;
      if($attrs.id) {
        id = $attrs.id;
      } 
	  $scope.onSubmit = function() {
		    var id = $scope.id;
			if(id!=null) {
				$scope.entity.model.Id = id;
				$http.put(entity.url+"/"+id, $scope.entity.model).then(
			       function(response) {
						console.log(response);
			       },
			       function(errorResponse) {
			         	console.log("error communicating with server");
			       }
			   );
			} else {
				$http.post(entity.url, $scope.entity.model).then(
			       function(response) {
					console.log(response);
			       },
			       function(errorResponse) {
					console.log(errorResponse);
			         	console.log("error communicating with server");
			       }
			   );
			}
	    };	
	  if(id != null) {
		 $attrs.$observe('id', function(passedId) {
		   $scope.id = passedId;
		   return $http.get(entity.url+"/"+passedId).then(
		       function(response) {
				var entitydata = response.data;
				$scope.entity.model = entitydata;
		       },
		       function(errorResponse) {
		         	console.log("error communicating with server");
		       }
		   );
	    });	
	  }
    }	
  }

})();
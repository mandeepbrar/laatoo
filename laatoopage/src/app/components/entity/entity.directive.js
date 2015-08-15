(function() {
  'use strict';

  angular
    .module('entity')
    .directive('entity', entity);

  /** @ngInject */
  function entity() {
    var directive = {
      restrict: 'E',
      templateUrl: 'app/components/entity/entity.view.html',
	  scope: { },
      replace: true,
      transclude: true,
      controller: EntityCtrl,
      controllerAs: 'entity',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function EntityCtrl($scope, $element, $attrs, $http ) {
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
      var name;
      if($attrs.name) {
        name = $attrs.name;
      } else {
		throw new Error("Server error. Entity name missing");	
	  }
	  var id;
      if($attrs.id) {
        id = $attrs.id;
      } else {
		throw new Error("Server error. Entity id missing");	
	  }
      var entity = window.pageConf.entities[name];	 
	  $attrs.$observe('id', function(passedId) {
	    return $http.get(entity.url+"/"+passedId).then(
	        function(response) {
				$scope.entitydata = response.data;
	        },
	        function(errorResponse) {
	          console.log("error communicating with server");
	        }
	    );
      });
	}	
  }

})();
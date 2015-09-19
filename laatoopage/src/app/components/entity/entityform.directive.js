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
    function EntityformCtrl($scope, $element, $attrs, $http, $state, dialogs) {
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
      var name;
      if($attrs.name) {
        name = $attrs.name;
      } else {
		dialogs.error('Error','Action could not be completed.');
		throw new Error("Server error. Entity name missing");	
	  }
      var entity = window.pageConf.entities[name];
	  $scope.entity = entity;
	  $scope.id = $attrs.id;
      $scope.submitText = "Submit";
      if($attrs.submitText) {
        $scope.submitText = $attrs.submitText;
      }
	  var successstate;
      if($attrs.successstate) {
        successstate = $attrs.successstate;
      }
	  var id = null;
      if($attrs.id) {
        id = $attrs.id;
      } 
	  $scope.onSubmit = function() {
		    var id = $scope.id;
			if(id!=null && id!= "") {
				$scope.entity.model.Id = id;
				$http.put(entity.url+"/"+id, $scope.entity.model).then(
			       function(response) {
						if(successstate) {
							$state.go(successstate);							
						}
			       },
			       function(errorResponse) {
			            if(errorResponse.status == 0) {
						  dialogs.error('Error','Could not connect to the website. Please check your internet connection or the website is offline.');
			               return;
			            }
						dialogs.error('Error','Action could not be completed. ' + errorResponse.statusText);
			       }
			   );
			} else {
				$http.post(entity.url, $scope.entity.model).then(
			       function(response) {
					if(successstate) {
						$state.go(successstate);							
					}
			       },
			       function(errorResponse) {
			            if(errorResponse.status == 0) {
						  dialogs.error('Error','Could not connect to the website. Please check your internet connection or the website is offline.');
			               return;
			            }
						dialogs.error('Error','Action could not be completed. ' + errorResponse.statusText);
			       }
			   );
			}
	    };	//submit ends
	  if(id != null) {
		$attrs.$observe('id', function(passedId) {
		    if(passedId) {
			    $scope.id = passedId;
			    return $http.get(entity.url+"/"+passedId).then(
			       function(response) {
					var entitydata = response.data;
					$scope.entity.model = entitydata;
			       },
			       function(errorResponse) {
			            if(errorResponse.status == 0) {
						  dialogs.error('Error','Could not connect to the website. Please check your internet connection or the website is offline.');
			               return;
			            }
						dialogs.error('Error','Action could not be completed. ' + errorResponse.statusText);
			       }
			    );
		    }
	    });	
	  } else {
		$scope.$watch(name, function(value) {
		  $scope[name].$setPristine();
		});
		$scope.entity.model = null;
	  }

	}	//controller ends
  }

})();
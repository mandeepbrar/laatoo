(function() {
    'use strict';

    angular
        .module('entity')
        .directive('entity', entity);

    /** @ngInject */
    function entity() {
        var directive = {
            restrict: 'E',
			templateUrl: function($element, $attrs) {				
				if($attrs.template) {
					return $attrs.template;
				}
	            return 'app/components/entity/entity.view.html';
			},
            scope: {},
            replace: true,
			link: function (scope, elem, attrs) {
				try {
					if(attrs.entitydata) {
						scope.entitydata = scope.$eval(attrs.entitydata);		
						scope.$watch(attrs.entitydata, function(passedval) {
							scope.entitydata = passedval;
						});						
					}
					if(scope.entitydata) {
						scope.entitydata.isOwner = (scope.entitydata.CreatedBy == scope.user);						
					}
				}catch(ex) {
					console.log(ex);
				}
			},
            transclude: true,
            controller: EntityCtrl,
            controllerAs: 'entity',
            bindToController: true
        };

        return directive;

        /** @ngInject */
        function EntityCtrl($scope, $element, $attrs, $http, dialogs, EntityDataService) {
            $scope.entitydata = {};
			$scope.user = localStorage.user;
            if ($attrs.class) {
                $scope.class = "class=" + $attrs.class;
            }
            var name;
            if ($attrs.name) {
                name = $attrs.name;
            } else {
                dialogs.error('Error', 'Action could not be completed.');
            }
            var id;
            if ($attrs.id) {
                id = $attrs.id;
            } 
			var entityFetchSuccessful = function(response){
                $scope.entitydata = response.data;
				try {
					$scope.entitydata.isOwner = ($scope.entitydata.CreatedBy == $scope.user);
				}
				catch(ex) {}
            };
			var entityFetchFailure = function(errorResponse) {
				console.log(errorResponse);
                if (errorResponse.status == 0) {
                    dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                    return;
                }
                dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
            };
            $attrs.$observe('id', function(passedId) {
				EntityDataService.GetEntity(name, passedId, entityFetchSuccessful, entityFetchFailure);
            });
        }
    }

})();
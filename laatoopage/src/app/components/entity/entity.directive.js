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
				if($attrs.viewmode) {
					var url = $attrs.name + "." + $attrs.viewmode + ".html";
					return url;
				}
	            return 'app/components/entity/entity.view.html';
			},
            scope: {},
            replace: true,
			link: function (scope, elem, attrs) {
				try {
					scope.entitydata = scope.$eval(attrs.entitydata);		
					scope.$watch(attrs.entitydata, function(passedval) {
						scope.entitydata = passedval;
					});
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
        function EntityCtrl($scope, $element, $attrs, $http, dialogs) {
            $scope.entitydata = {};
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
            var entity = window.pageConf.entities[name];
            $attrs.$observe('id', function(passedId) {
                return $http.get(entity.url + "/" + passedId).then(
                    function(response) {
                        $scope.entitydata = response.data;
                    },
                    function(errorResponse) {
                        if (errorResponse.status == 0) {
                            dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                            return;
                        }
                        dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                    }
                );
            });
        }
    }

})();
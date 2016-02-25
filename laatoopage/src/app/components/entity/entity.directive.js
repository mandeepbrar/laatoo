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
            scope: {},
            replace: true,
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
            } else {
                throw new Error("Server error. Entity id missing");
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
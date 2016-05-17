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
            scope: {},
            replace: true,
            link: function(scope, elem, attrs) {
				if(attrs.model) {
					var modelVal = scope.$eval(attrs.model);
					scope.entity.model = modelVal;
				}
            },
            controller: EntityformCtrl,
            controllerAs: 'entityform',
            bindToController: true
        };

        return directive;

        /** @ngInject */
        function EntityformCtrl($scope, $element, $attrs, EntityDataService, $state, dialogs) {
            if ($attrs.class) {
                $scope.class = "class=" + $attrs.class;
            }
            var name;
            if ($attrs.name) {
                name = $attrs.name;
            } else {
                dialogs.error('Error', 'Action could not be completed.');
                throw new Error("Server error. Entity name missing");
            }
            var entity = document.EntityForms[name];
            $scope.entity = entity;
			console.log($scope.entity);
            $scope.id = $attrs.id;
            $scope.submitText = "Submit";
            if ($attrs.submitText) {
                $scope.submitText = $attrs.submitText;
            }
            var successstate;
            if ($attrs.successstate) {
                successstate = $attrs.successstate;
            }
            var id = null;
            if ($attrs.id) {
                id = $attrs.id;
            }
            $scope.onSubmit = function() {
                var id = $scope.id;
                if (id != null && id != "") {
                    $scope.entity.model.Id = id;
					var successMethod = function(response) {
                            if (successstate) {
                                $state.go(successstate);
                            }
                        };
					var failureMethod = function(errorResponse) {
                            if (errorResponse.status == 0) {
                                dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                                return;
                            }
                            dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                        }
					EntityDataService.PutEntity(name, id, $scope.entity.model, successMethod, failureMethod);
                } else {
                    var successMethod = function(response) {
                            console.log(successstate);
                            if (successstate) {
                                $state.go(successstate);
                            }
                        };
                    var failureMethod = function(errorResponse) {
                            if (errorResponse.status == 0) {
                                dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                                return;
                            }
                            dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                        };
                    EntityDataService.SaveEntity(name, $scope.entity.model, successMethod, failureMethod);
                }
            }; //submit ends
            if (id != null) {
                $attrs.$observe('id', function(passedId) {
                    if (passedId) {
                        $scope.id = passedId;
                        var successMethod = function(response) {
                                var entitydata = response.data;
                                $scope.entity.model = entitydata;
                            };
                        var failureMethod = function(errorResponse) {
                                if (errorResponse.status == 0) {
                                    dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                                    return;
                                }
                                dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                            };
                        EntityDataService.GetEntity(name, passedId, successMethod, failureMethod);
                    }
                });
            } else {
                $scope.$watch(name, function(value) {
                    $scope[name].$setPristine();
                });
                $scope.entity.model = {};
            }

        } //controller ends
    }

})();
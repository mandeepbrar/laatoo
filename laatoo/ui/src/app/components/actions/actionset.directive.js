(function() {
    'use strict';

    var actionsapp = angular
        .module('actions');
    actionsapp.directive('actionset', actionset);

    /** @ngInject */
    function actionset() {
        var directive = {
            restrict: 'E',
            templateUrl: function($element, $attrs) {
                if ($attrs.widget === 'menu') {
                    return 'app/components/actions/menu.view.html';
                }
                if ($attrs.widget === 'tab') {
                    return 'app/components/actions/tabs.view.html';
                }
                if ($attrs.widget === 'pills') {
                    return 'app/components/actions/pills.view.html';
                }
                return 'app/components/actions/menu.view.html';
            },
            replace: true,
            scope: {},
            controller: ActionsetCtrl,
            controllerAs: 'actionset',
            bindToController: true
        };
        return directive;

        /** @ngInject */
        function ActionsetCtrl($scope, $element, $attrs, $templateCache) {
            var name = $attrs.name;
            if ($attrs.class) {
                $scope.class = "class=" + $attrs.class;
            }
            var view = name;
            if ($attrs.view) {
                view = $attrs.view;
            }
            var items = new Array();
			$scope.items = items;
			document.RunComplete.then(function(response) {
	            try {
	                var actionset = document.Actionset[name];
	                for (var key in actionset) {
	                    var value = actionset[key];
	                    var actionName = value.action;
	                    var action = document.Actions[actionName];
	                    if (action.allowed) {
	                        var item = {};
	                        item.heading = value.label;
	                        item.route = actionName;
	                        item.actiontype = action.actiontype;
	                        items.push(item);
	                    }
	                }
	            } catch (exc) {
	                console.log(exc);
	            }
			});
        }
    }

})();
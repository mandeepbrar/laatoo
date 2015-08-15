(function() {
  'use strict';

  var actionsapp = angular
    .module('actions');
  actionsapp.directive('action', action);

  /** @ngInject */
  function action() {
    var directive = {
      restrict: 'E',
      templateUrl: function($element, $attrs) {
			if($attrs.widget === 'button') {
				return 'app/components/actions/button.view.html';
			}
			return 'app/components/actions/button.view.html';
		},
      replace: true,
	  scope: { },
      controller: ActionCtrl,
      controllerAs: 'action',
      bindToController: true
    };
    return directive;

    /** @ngInject */
    function ActionCtrl($scope, $element, $attrs, $templateCache ) {
      var name = $attrs.name;
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
      if($attrs.label) {
        $scope.label = $attrs.label;
      }
	  $scope.route = name;
  	  $scope.action = window.pageConf.actions[name];
    }
  }

})();
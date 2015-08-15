(function() {
  'use strict';

  angular
    .module('view')
    .directive('view', view);

  /** @ngInject */
  function view() {
    var directive = {
      restrict: 'E',
      templateUrl: function($element, $attrs) {
			if($attrs.viewtype === 'ul') {
				return 'app/components/links/ul.view.html';
			}
			if($attrs.viewtype === 'table') {
				return 'app/components/view/table.view.html';
			}
			return 'app/components/view/view.view.html';
		},
      replace: true,
	  scope: { },
      transclude: true,
      controller: ViewCtrl,
      controllerAs: 'view',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function ViewCtrl($scope, $element, $attrs, ViewService ) {
      var name = $attrs.name;
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
	  $scope.viewrows = ViewService.query({viewname: name});	  
    }
  }

})();
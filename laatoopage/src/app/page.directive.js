(function() {
  'use strict';

  angular
    .module('main')
    .directive('page', page);

  /** @ngInject */
  function page($http) {
    var directive = {
      restrict: 'E',
      templateUrl: 'app/page.view.html',
      transclude: true,
      replace: true,
      controller: PageCtrl,
      controllerAs: 'page',
      bindToController: true
    };


    return directive;

    /** @ngInject */
    function PageCtrl($scope, $element, $attrs, $stateParams ) {
	  $scope.params = $stateParams;	
      var name = $attrs.name;
      $scope.confUrl = document.location.href + "/conf";
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
    }
  }

})();
(function() {
  'use strict';

  angular
    .module('login')
    .directive('login', login);

  /** @ngInject */
  function login() {
    var directive = {
      restrict: 'E',
      templateUrl: 'app/components/login/login.view.html',
      replace: true,
      controller: LoginCtrl,
      controllerAs: 'login',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function LoginCtrl($scope, $element, $attrs ) {
      var name = $attrs.name;
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
      var social = $attrs.social;
      if($attrs.social) {
        $scope.social = true;
      }
      var social = $attrs.signup;
      if($attrs.signup) {
        $scope.signup = true;
      }
    }
  }

})();
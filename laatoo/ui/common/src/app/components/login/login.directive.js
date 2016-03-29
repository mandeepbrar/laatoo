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
      if($attrs.social) {
        $scope.social = true;
      }
      if($attrs.google) {
        $scope.google = $attrs.google;
      }
      if($attrs.facebook) {
        $scope.facebook = $attrs.facebook;
      }
      var social = $attrs.signup;
      if($attrs.signup) {
        $scope.signup = true;
      }
    }
  }

})();
(function() {
  'use strict';

  var app = angular.module('loginui');
  app.config(routeConfig);
  app.run(runFunc);

  /** @ngInject */
  function routeConfig($stateProvider, $urlRouterProvider) {
    $stateProvider
      .state('home', {
        url: '/',
        templateUrl: 'app/main/views/index.view.html',
          data: {
            requireLogin: true
          }
      })
      .state('login', {
        url: '/login',
        templateUrl: 'app/main/views/login.view.html',
        data: {
          requireLogin: false
        }
      });

    $urlRouterProvider.otherwise('/');
  }

  /** @ngInject */
  function runFunc($rootScope, $state) {

    $rootScope.$on('$stateChangeStart', function (event, toState, toParams) {
      var requireLogin = toState.data.requireLogin;
      if (requireLogin && typeof $rootScope.currentUser === 'undefined') {
        event.preventDefault();
        $state.go('login');
      }
    });
  }

})();

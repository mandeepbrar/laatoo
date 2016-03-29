(function() {
  'use strict';

  var mainapp = angular
    .module('main')
    .config(config);

  /** @ngInject */
  function config($stateProvider, $httpProvider, dialogsProvider, $translateProvider, $provide) {
		initializeApp(mainapp, $stateProvider, $httpProvider, dialogsProvider, $translateProvider, $provide);
  }

})();

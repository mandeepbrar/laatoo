(function() {
  'use strict';

  angular
    .module('main')
    .config(config);

  /** @ngInject */
  function config($logProvider) {
    // Enable log
    //$logProvider.debugEnabled(true);
  }

})();

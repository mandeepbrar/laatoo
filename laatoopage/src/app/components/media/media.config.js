(function() {
    'use strict';

    var mediapp = angular
        .module('media')
        .config(config);

    /** @ngInject */
    function config($sceDelegateProvider) {
        $sceDelegateProvider.resourceUrlWhitelist([
            'self',
            'http://youtube.com/**',
            'http://www.youtube.com/**',
            'https://youtube.com/**',
            'https://www.youtube.com/**',
        ]);
    }

})();
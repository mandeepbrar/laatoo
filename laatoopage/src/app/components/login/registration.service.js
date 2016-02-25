(function() {
    'use strict';

    var mod = angular.module('login');

    mod.factory('RegistrationService', RegistrationService);

    /** @ngInject */
    function RegistrationService($resource) {
        var data = $resource(pageConf.LocalRegServer, {}, {
            register: {
                method: 'POST'
            }
        });
        return data;
    }



})();
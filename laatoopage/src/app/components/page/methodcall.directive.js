(function() {
    'use strict';

    angular
        .module('page')
        .directive('methodCall', methodcall);

    /** @ngInject */
    function methodcall() {
        var directive = {
            restrict: 'A',
            link: function(scope, elem, attrs) {
				var methodToCall = attrs.methodCall;
                elem.ready(function() {
	                try {
						var methodtocall = scope.$eval(methodToCall);
						methodtocall();
	                } catch (ex) {
	                    console.log(ex);
	                }
                })
            },
        };
        return directive;
    }

})();
(function() {
    'use strict';

    angular
        .module('page')
        .directive('includeReplace', includeReplace);

    /** @ngInject */
    function includeReplace() {
        var directive = {
	        require: 'ngInclude',
	        restrict: 'A', /* optional */
	        link: function (scope, el, attrs) {
				console.log(attrs);
	            el.replaceWith(el.children());
				for(var attr in attrs) {
					var val = attrs[attr];
					console.log(attr);
					if(typeof(val)=="string") {
						scope[attr] = attrs[attr];						
					}
				}
				console.log(scope);
	        }
        };
        return directive;
    }

})();
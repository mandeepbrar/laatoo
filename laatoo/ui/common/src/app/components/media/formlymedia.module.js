(function() {
  'use strict';

   	var formlymedia = angular.module('formlymedia', ['ngResource', 'formly', 'formlyBootstrap']);

	formlymedia.config(formlymediaconfig);

    /** @ngInject */
    function formlymediaconfig(formlyConfigProvider) {
        formlyConfigProvider.setType({
            name: 'media',
            templateUrl: 'app/components/media/formlymedia.view.html',
            controller: mediaEditcontroller
        });
        formlyConfigProvider.templateManipulators.preWrapper.push(function(template, options, scope) {
            if (options.type == "media") {
                scope.haslabel = false;
                scope.field = options.key;
            }
            return template;
        });
    }

    /** @ngInject */
    function mediaEditcontroller($scope, $modal) {
        $scope.field = $scope.options.key;
        if ($scope.options.templateOptions.mediatype) {
            $scope.mediatype = $scope.options.templateOptions.mediatype;
        } else {
            $scope.mediatype = 'image';
        }
        $scope.allowupload = "true";
        if ($scope.options.templateOptions.allowupload) {
            $scope.allowupload = $scope.options.templateOptions.allowupload;
        }
		$scope.allowlibrary = "false";
        if ($scope.options.templateOptions.allowlibrary) {
            $scope.allowlibrary = $scope.options.templateOptions.allowlibrary;
        }
		$scope.allowexternalurl = "false";
        if ($scope.options.templateOptions.allowexternalurl) {
            $scope.allowexternalurl = $scope.options.templateOptions.allowexternalurl;
        }
        $scope.label = $scope.options.templateOptions.label;

    }


})();
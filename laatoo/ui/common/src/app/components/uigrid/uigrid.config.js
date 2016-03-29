(function() {
    'use strict';

    var uigridapp = angular
        .module('uigrid')
        .config(config);

    /** @ngInject */
    function config(formlyConfigProvider) {
        formlyConfigProvider.setType({
            name: 'ui-grid',
            templateUrl: 'app/components/uigrid/multiselect.view.html',
            controller: MultiselectCtrl
        });
        formlyConfigProvider.templateManipulators.preWrapper.push(function(template, options, scope) {
            if (options.type == "ui-grid") {
                scope.haslabel = false;
                scope.field = options.key;
            }
            return template;
        });
    }

    /** @ngInject */
    function MultiselectCtrl($scope, $http) {
        $scope.field = "";
        if ($scope.options.templateOptions.mediatype) {
            $scope.mediatype = $scope.options.templateOptions.mediatype;
        } else {
            $scope.mediatype = 'image';
        }
        var setSelections = function() {
            for (var index in $scope.griditems) {
                var item = $scope.griditems[index];
                if ($scope.valueField != "=") {
                    var valToTest = item[$scope.valueField];
                    try {
                        $scope.status[item[$scope.valueField]] = ($scope.selected.indexOf(item[$scope.valueField]) > -1);
                    } catch (ex) {}
                } else {
                    try {
                        $scope.status[item] = ($scope.selected.indexOf(item) > -1);
                    } catch (ex) {}
                }
            }
        }
        $scope.$watch('griditems', function(newValue) {
            setSelections();
        });
        $scope.$watch('model', function(newValue) {
            try {
                $scope.status = [];
                $scope.selected = newValue[$scope.options.key];
                var gridcallback = $scope.options.templateOptions.gridcallback;
                if (gridcallback) {
                    $http.get(gridcallback).then(
                        function(response) {
                            $scope.griditems = response.data;
                        },
                        function(errorResponse) {
                            console.log("error communicating with server");
                        }
                    );
                } else {
                    $scope.griditems = $scope.options.templateOptions.griditems;
                }
                $scope.label = $scope.options.templateOptions.label;
                $scope.columns = $scope.options.templateOptions.columns;
                if ($scope.options.templateOptions.valueField) {
                    $scope.valueField = $scope.options.templateOptions.valueField;
                } else {
                    $scope.valueField = $scope.options.key;
                }
                setSelections();
            } catch (ex) {}
        });
        $scope.oncheckboxchange = function(evt, val) {
            var modelVal = $scope.model[$scope.options.key];
            if (!modelVal || !(modelVal instanceof Array)) {
                modelVal = [];
            }
            if (evt.target.checked) {
                modelVal.push(val);
            } else {
                var index = modelVal.indexOf(val);
                if (index > -1) {
                    modelVal.splice(index, 1);
                }
            }
            $scope.model[$scope.options.key] = modelVal;
        };
    }


})();
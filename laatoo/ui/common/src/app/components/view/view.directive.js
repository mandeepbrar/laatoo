(function() {
    'use strict';

    angular
        .module('view')
        .directive('view', view);

    /** @ngInject */
    function view() {
        var directive = {
            restrict: 'E',
            templateUrl: function($element, $attrs) {
                if ($attrs.viewtype === 'ul') {
                    return 'app/components/links/ul.view.html';
                }
                if ($attrs.viewtype === 'table') {
                    return 'app/components/view/table.view.html';
                }
                return 'app/components/view/view.view.html';
            },
            replace: true,
            scope: {},
            link: function(scope, elem, attrs) {
                elem.ready(function() {
                    scope.refreshView();
                })
            },
            transclude: true,
            controller: ViewCtrl,
            controllerAs: 'view',
            bindToController: true
        };

        return directive;

        /** @ngInject */
        function ViewCtrl($scope, $element, $attrs, ViewService, $http, dialogs) {
            var url = $attrs.url;
            $scope.params = {};
            $scope.modelname = 'viewrows';
            if ($attrs.class) {
                $scope.class = "class=" + $attrs.class;
            }
            if ($attrs.modelname) {
                $scope.modelname = $attrs.modelname;
            }
            $attrs.$observe('url', function(passedval) {
                url = passedval;
            });
            $attrs.$observe('args', function(passedval) {
                $scope.params.args = angular.fromJson(passedval);
            });
            if ($attrs.editable) {
                $scope.editable = ($attrs.editable == 'true');
		      	$scope.qualifier = $attrs.qualifier;
                $scope.submitText = "Save";
                $scope.onSubmit = function() {
	                var actionUrl = "";
	                if ($attrs.action) {
	                    actionUrl = $attrs.action;
	                }
					console.log(actionUrl);
					var data = [];
					var qualifiedData = [];
					if($scope.qualifier) {
						var allrows = $scope[$scope.modelname];
						for(var index in allrows) {
							var row = allrows[index];
							if(row[$scope.qualifier]) {
								qualifiedData.push(row);
							}
						}
					}
					else {
						qualifiedData = $scope[$scope.modelname];
					}	
		
					if($attrs.fields) {
						var fields = $attrs.fields.split(",");
						for (var i=0; i < qualifiedData.length; i++) {
							var newRow = {};
							for (var j=0; j< fields.length; j++) {
								newRow[fields[j]] = qualifiedData[i][fields[j]];
							}								
							data.push(newRow);
						}
					}
					else {
						data = qualifiedData;
					}
					var objToSend = {data: data, params: $scope.params}
					$http.put(actionUrl, objToSend).then(
		                        function(response) {
		                            dialogs.notify('Success', 'Action completed successfully.');
		                        },
		                        function(errorResponse) {
		                            if (errorResponse.status == 0) {
		                                dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
		                                return;
		                            }
		                            dialogs.error('Error', 'Action could not be completed. ' + errorResponse);
		                        }
		                    );
                };
            }
            $scope.pagenum = 1;
            $scope.paginate = false;
            if ($attrs.paginate) {
                $scope.paginate = ($attrs.paginate == "true");
            }
            if ($attrs.pagesize) {
                $scope.pagesize = parseInt($attrs.pagesize);
            } else {
                $scope.pagesize = 10;
            }
            $scope.refreshView = function() {
                if ($scope.paginate) {
                    $scope.params.pagesize = $scope.pagesize;
                    $scope.params.pagenum = $scope.pagenum;
                }
                $http.get(url, {
                    params: $scope.params
                }).then(
                    function(response) {
					if($scope.qualifier) {
						for (var index in response.data) {
							var row = response.data[index];
							if(!row[$scope.qualifier]) {
								row[$scope.qualifier] = false;
							}
						}
					}					
                        $scope[$scope.modelname] = response.data;
                        var headers = response.headers();
                        $scope.records = headers["records"]
                        $scope.totalrecords = headers["totalrecords"]
                    },
                    function(errorResponse, status) {
                        if (errorResponse.status == 0) {
                            dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                            return;
                        }
                        if (errorResponse.status == 401) {
                            window.location.href = window.pageConf.AuthPage;
                        } else {
                            dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                        }
                    }
                );
                //  $scope[$scope.modelname] = ViewService.query($scope.params);          
            };
            $scope.$on("refresh", $scope.refreshView);
        }
    }

})();
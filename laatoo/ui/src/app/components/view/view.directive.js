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
        function ViewCtrl($scope, $element, $attrs, DataService, RequestBuilderService, $http, dialogs) {
            var servicename = $attrs.servicename;
            $scope.params = {};
            $scope.params.args = {};
            $scope.params.url = {};
            $scope.modelname = 'viewrows';
            if ($attrs.class) {
                $scope.class = "class=" + $attrs.class;
            }
            if ($attrs.modelname) {
                $scope.modelname = $attrs.modelname;
            }
            $attrs.$observe('servicename', function(passedval) {
                servicename = passedval;
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
									console.log(errorResponse);
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
			$scope.dynamicpagination = false;
            if ($attrs.paginate) {
                $scope.paginate = ($attrs.paginate == "true");
            }
            if ($attrs.dynamicpagination) {
                $scope.dynamicpagination = ($attrs.dynamicpagination == "true");
            }
            if ($attrs.pagesize) {
                $scope.pagesize = parseInt($attrs.pagesize);
            } else {
                $scope.pagesize = 10;
            }
			if($scope.dynamicpagination) {
				$scope.nextPage = function() {
					//change the page if pagenum is positive
					if($scope.pagenum > 0) {
						$scope.pagenum = $scope.pagenum + 1;
						$scope.refreshView();						
					}
				};
			}
            $scope.refreshView = function() {
                if ($scope.paginate || $scope.dynamicpagination) {
                    $scope.params.url.pagesize = $scope.pagesize;
                    $scope.params.url.pagenum = $scope.pagenum;
                }
				var getviewsuccess = function(response) {
					if($scope.qualifier) {
						for (var index in response.data) {
							var row = response.data[index];
							if(!row[$scope.qualifier]) {
								row[$scope.qualifier] = false;
							}
						}
					}					
					if ($scope.pagenum == 1 || !$scope.dynamicpagination) {
	                    $scope[$scope.modelname] = response.data;													
					} else {
						console.log("dynamic pagination");
	                    var existingData = $scope[$scope.modelname];						
						console.log(existingData);
						var newData = response.data;
						console.log(newData);
						for(var dataItem in newData) {
							existingData.push(dataItem);
						}
					}
					if(response.info) {
	                    $scope.records = response.info["records"];
	                    $scope.totalrecords = response.info["totalrecords"];	
						if ($scope.records < $scope.pagesize ) {
							$scope.pagenum = -1;
						}										
					}
				};
				var getviewfailure = function(errorResponse) {
					console.log(errorResponse);
					console.log(servicename);
                    if (errorResponse.status == 0) {
                        dialogs.error('Error', 'Could not connect to the website. Please check your internet connection or the website is offline.');
                        return;
                    }
                    if (errorResponse.status == 401) {
                        window.location.href = window.pageConf.AuthPage;
                    } else {
                        dialogs.error('Error', 'Action could not be completed. ' + errorResponse.statusText);
                    }
				};
				var req = RequestBuilderService.URLParamsRequest($scope.params.url, $scope.params.args);
				DataService.ExecuteService(servicename, req, getviewsuccess, getviewfailure);
            };
            $scope.$on("refresh", $scope.refreshView);
        }
    }

})();
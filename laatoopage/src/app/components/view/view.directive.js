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
			if($attrs.viewtype === 'ul') {
				return 'app/components/links/ul.view.html';
			}
			if($attrs.viewtype === 'table') {
				return 'app/components/view/table.view.html';
			}
			return 'app/components/view/view.view.html';
		},
      replace: true,
	  scope: { },
      transclude: true,
      controller: ViewCtrl,
      controllerAs: 'view',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function ViewCtrl($scope, $element, $attrs, ViewService, $http) {
      var name = $attrs.name;
	  $scope.params = {};
	  $scope.modelname = 'viewrows';
      if($attrs.class) {
      	$scope.class = "class="+$attrs.class;
      }
  	  if($attrs.args) {
		$scope.params = angular.fromJson($attrs.args);
	  }
  	  if($attrs.modelname) {
		$scope.modelname = $attrs.modelname;
	  }
  	  if($attrs.viewserver) {
	    $scope.viewserver = $attrs.viewserver;
	  }
      if($attrs.editable) {
      	$scope.editable = ($attrs.editable == 'true');
		$scope.submitText = "Save";
		var actionUrl = "";
	  	if($attrs.action) {
			actionUrl = $attrs.action;
	  	}
		$scope.onSubmit = function() {
			console.log("submit view");
			$http.put(actionUrl, $scope[$scope.modelname]).then(
		       function(response) {
					console.log(response);
		       },
		       function(errorResponse) {
		         	console.log("error communicating with server");
		       }
		   );
		};
      }
	  $scope.pagenum = 1;
	  $scope.paginate = false;
	  if($attrs.paginate) {
	    $scope.paginate = ($attrs.paginate == "true");	
 	  }
      if($attrs.pagesize) {
		  $scope.pagesize = parseInt($attrs.pagesize);
	  } else {
		$scope.pagesize = 10;
	  }	
	  $scope.params['viewname'] = name;
	  $scope.refreshView = function() {
		  var url = pageConf.ViewsServer;
		  if($scope.viewserver) {
			url = $scope.viewserver;	
		  }
		  if($scope.paginate) {
			  $scope.params.pagesize = $scope.pagesize;
			  $scope.params.pagenum = $scope.pagenum;			
		  }
		  $http.get(url, { params: $scope.params }).then(
		       function(response) {
					$scope[$scope.modelname] = response.data;
					var headers = response.headers();
					$scope.records = headers["records"]
					$scope.totalrecords = headers["totalrecords"]
		       },
		       function(errorResponse) {
				  if(errorResponse.status == 401) {
					window.location.href = window.pageConf.AuthPage;								
				  } else {
		         	  console.log("error communicating with server");					
				   }
		       }
		   );
		//  $scope[$scope.modelname] = ViewService.query($scope.params);	  		
	  };
	  $scope.refreshView();
    }
  }

})();
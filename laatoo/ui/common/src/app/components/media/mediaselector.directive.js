(function() {
    'use strict';

    angular
        .module('media')
        .directive('mediaselector', mediaselector);

    /** @ngInject */
    function mediaselector() {
        var directive = {
            restrict: 'E',
            templateUrl: function($element, $attrs) {
                if ($attrs.mode === 'inplace') {
	            	return 'app/components/media/inplaceselector.view.html';
                }
            	return 'app/components/media/mediaselector.view.html';
            },
            replace: true,
            scope: {},
            link: function(scope, elem, attrs) {
                elem.ready(function() {
				    try {
						scope.model = scope.$eval(attrs.model);
						if(scope.model) {										
				       		scope.mediasource = scope.model[scope.field];
						}
				    } catch (ex) {}	
                })
            },
            transclude: true,
            controller: MediaSelectCtrl,
            controllerAs: 'mediaselect',
            bindToController: true
        };

        return directive;

        /** @ngInject */
        function MediaSelectCtrl($scope, $element, $attrs, $http, $modal) {
			$scope.model = {};
	        $scope.field = "";
	        if ($attrs.mediatype) {
	            $scope.mediatype = $attrs.mediatype;
	        } else {
	            $scope.mediatype = 'image';
	        }
			var defaultimage = $attrs.defaultimage;
			$attrs.$observe('mediatype', function(newVal) {
				$scope.mediatype = newVal;
			})
	        var allowupload = "false";
	        $attrs.$observe('allowupload', function(newVal) { 
	            allowupload = newVal;
	        });
	        var allowlibrary = "false";
	        $attrs.$observe('allowlibrary', function(newVal) { 
	            allowlibrary = newVal;
	        });
	        var allowexternalurl = "false";
	        $attrs.$observe('allowexternalurl', function(newVal) { 
	            allowexternalurl = newVal;
	        });
			$attrs.$observe('model', function(passedVal) {
				$scope.$watchCollection(passedVal, function(newmodelVal) {
					$scope.model = newmodelVal; //$scope.$eval($attrs.model);	
					if($scope.model) {				
			       		$scope.mediasource = $scope.model[$scope.field];
					}
				});					
			});
			
	        $scope.label = $attrs.label;
			$attrs.$observe('label', function(newVal) {
				$scope.label = newVal;
			})
	        $scope.field = $attrs.field;
			$attrs.$observe('field', function(newVal) {
				$scope.field = newVal;
			})
	        $scope.chooseMedia = function(field, type) {
	            var modalOptions = {
	                backdrop: true,
	                keyboard: true,
	                modalFade: true,
	                templateUrl: 'app/components/media/modalchooser.view.html',
	                controller: modalDialogController,
	                closeButtonText: 'Close',
	                resolve: {
	                    mediatype: function() {
	                        return $scope.mediatype;
	                    },
	                    allowupload: function() {
	                        return allowupload;
	                    },
	                    allowlibrary: function() {
	                        return allowlibrary;
	                    },
	                    allowexternalurl: function() {
	                        return allowexternalurl;
	                    }
	                },
	                actionButtonText: 'OK'
	            };
	            //$scope.mediatype = type;
	            var dialog = $modal.open(modalOptions);
	            dialog.result.then(function(result) {
					var model = $scope.$eval($attrs.model);
	                model[$scope.field] = result;
	                $scope.mediasource = result;
	            });
	        };
	        $scope.removeMedia = function(field) {
				var model = $scope.$eval($attrs.model);
	            model[$scope.field] = "";
				if(defaultimage) {
		            $scope.mediasource = defaultimage;					
				} else {
		            $scope.mediasource = "";					
				}
	        };
	    }
	
	    /** @ngInject */
	    function modalDialogController($scope, $modalInstance, FileUploader, mediatype, allowupload, allowlibrary, allowexternalurl) {
	        $scope.mediatype = mediatype;
	        $scope.allowupload = allowupload;
	        $scope.allowlibrary = allowlibrary;
	        $scope.allowexternalurl = allowexternalurl;
	        var uploader = $scope.uploader = new FileUploader({
	            url: '/upload',
	            queueLimit: 1
	        });
	        uploader.onCompleteItem = function(fileItem, response, status, headers) {
	            if (status == 200 && response.length > 0) {
	                $modalInstance.close(response[0]);
	            }
	        };
	        $scope.closeDialog = function() {
	            $modalInstance.dismiss("closed");
	        }
	        $scope.fileSelected = function(url) {
	            if (url.length != 0) {
	                $modalInstance.close(url);
	            }
	        }
	        $scope.uploadFile = function() {
	            var selectedFile = uploader.queue[0];
	            selectedFile.upload();
	        }
	    }
    }

})();
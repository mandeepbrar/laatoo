(function() {
  'use strict';

  var mediapp = angular
    .module('media')
    .config(config);

  /** @ngInject */
  function config($sceDelegateProvider, formlyConfigProvider) {
	$sceDelegateProvider.resourceUrlWhitelist([
	    'self',
	    'http://youtube.com/**',
	    'http://www.youtube.com/**',
	    'https://youtube.com/**',
	    'https://www.youtube.com/**',
  	]); 
	formlyConfigProvider.setType({
	  name: 'media',
      templateUrl: 'app/components/media/mediaselector.view.html',
	  controller: mediaEditcontroller
	});	
	formlyConfigProvider.templateManipulators.preWrapper.push(function(template, options, scope) {
		if(options.type == "media") {
			scope.haslabel = false;
			scope.field = options.key;
			scope.allowupload = options.templateOptions.allowupload;
			scope.$watch('model', function (newValue) {
					try {
						scope.mediasource = newValue[options.key];
						scope.label = options.templateOptions.label;
					}catch(ex){}
		    });
		}
		return template;
	});	
  }
  
  /** @ngInject */
  function mediaEditcontroller($scope, $modal) {
		$scope.field = "";
		if($scope.options.templateOptions.mediatype) {
			$scope.mediatype = $scope.options.templateOptions.mediatype;
		} else {
			$scope.mediatype = 'image';
		}
		
		console.log($scope);
		var allowupload = "false";
		if($scope.options.templateOptions.allowupload) {
			allowupload = $scope.options.templateOptions.allowupload;
			console.log(allowupload);
		}
		
		$scope.chooseMedia = function(field, type){
			var modalOptions = {
	        		backdrop: true,
	        		keyboard: true,
	        		modalFade: true,
	        		templateUrl: 'app/components/media/modalchooser.view.html',
				controller: modalDialogController,
	        		closeButtonText: 'Close',
				resolve: {
         			mediatype: function () {	
         			  	return type;
         			},
					allowupload: function() {
						console.log("return "+allowupload);
						return allowupload;
					}
       			},
	            actionButtonText: 'OK'
	        };
			//$scope.mediatype = type;
			var dialog = $modal.open(modalOptions);
			dialog.result.then(function (result) {
					$scope.model[$scope.field] = result;
					$scope.mediasource = result;
		        });
	    };	
		$scope.removeMedia = function(field) {
			$scope.model[field] = "";
			$scope.mediasource = "";
		};	
  }

  /** @ngInject */
  function modalDialogController($scope, $modalInstance, FileUploader, mediatype, allowupload) {
	$scope.mediatype = mediatype;
	$scope.allowupload = allowupload;
    var uploader = $scope.uploader = new FileUploader({
        url: '/upload',
		queueLimit: 1
    });
	uploader.onCompleteItem = function(fileItem, response, status, headers) {
		if(status == 200 && response.length >0) {
			$modalInstance.close(response[0]);
		}
    };
	$scope.closeDialog = function() {
		$modalInstance.dismiss("closed");
	}
	$scope.fileSelected = function(url) {
		if(url.length != 0) {
			$modalInstance.close(url);
		}
	}
	$scope.uploadFile = function() {
		var selectedFile = uploader.queue[0];
		selectedFile.upload();
	}
  }

})();



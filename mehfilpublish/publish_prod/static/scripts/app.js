!function(){"use strict";angular.module("view",["ngResource","ngStorage","ui.router","ui.bootstrap"])}(),function(){"use strict";function e(){function e(e,t,o){o.popupsearch&&console.log("popup search"),e.params={},e.applyFilter=function(){var t=e.$parent;t.params.args=e.params,t.refreshView()},e.resetFilter=function(){parent=e.$parent,parent.params.args=e.params={},parent.refreshView()}}var t={restrict:"E",template:"<div ng-transclude></div>",replace:!0,transclude:!0,controller:e,controllerAs:"viewsearch",bindToController:!0};return e.$inject=["$scope","$element","$attrs"],t}angular.module("view").directive("viewsearch",e)}(),function(){"use strict";function e(e){var t=e(pageConf.ViewsServer);return t}var t=angular.module("view");t.factory("ViewService",e),e.$inject=["$resource"]}(),function(){"use strict";function e(){function e(e,t,o,n,i){var a=o.name;if(e.params={},e.modelname="viewrows",o["class"]&&(e["class"]="class="+o["class"]),o.args&&(e.params=angular.fromJson(o.args)),o.modelname&&(e.modelname=o.modelname),o.viewserver&&(e.viewserver=o.viewserver),o.editable){e.editable="true"==o.editable,e.submitText="Save";var l="";o.action&&(l=o.action),e.onSubmit=function(){console.log("submit view"),i.put(l,e[e.modelname]).then(function(e){console.log(e)},function(e){console.log("error communicating with server")})}}e.params.viewname=a,e.refreshView=function(){var t=pageConf.ViewsServer;e.viewserver&&(t=e.viewserver),i.get(t,{params:e.params}).then(function(t){e[e.modelname]=t.data},function(e){401==e.status?window.location.href=window.pageConf.AuthPage:console.log("error communicating with server")})},e.refreshView()}var t={restrict:"E",templateUrl:function(e,t){return"ul"===t.viewtype?"app/components/links/ul.view.html":"table"===t.viewtype?"app/components/view/table.view.html":"app/components/view/view.view.html"},replace:!0,scope:{},transclude:!0,controller:e,controllerAs:"view",bindToController:!0};return e.$inject=["$scope","$element","$attrs","ViewService","$http"],t}angular.module("view").directive("view",e)}(),function(){"use strict";angular.module("uigrid",["ngResource","ui.bootstrap","formly","formlyBootstrap","smart-table"])}(),function(){"use strict";function e(e){e.setType({name:"ui-grid",templateUrl:"app/components/uigrid/multiselect.view.html",controller:t}),e.templateManipulators.preWrapper.push(function(e,t,o){return"ui-grid"==t.type&&(o.haslabel=!1,o.field=t.key,o.$watch("model",function(e){try{o.selected=e[t.key],o.griditems=t.templateOptions.griditems,o.status=[],o.label=t.templateOptions.label,o.columns=t.templateOptions.columns,o.valueField=t.templateOptions.value?t.templateOptions.valueField:t.key;for(var n in o.griditems){{var i=o.griditems[n];i[o.valueField]}o.status[i[o.valueField]]=o.selected.indexOf(i[o.valueField])>-1}}catch(a){}})),e})}function t(e){e.field="",e.mediatype=e.options.templateOptions.mediatype?e.options.templateOptions.mediatype:"image",console.log("setting up change method"),e.oncheckboxchange=function(t,o){console.log(e.model);var n=e.model[e.options.key];if(console.log("checkbox change"),console.log(n),n&&n instanceof Array||(console.log("initializing"),n=[]),console.log(n),t.target.checked)console.log("pushing "+o),n.push(o);else{var i=n.indexOf(o);i>-1&&n.splice(i,1)}e.model[e.options.key]=n,console.log(e.model)},console.log(e.oncheckboxchange)}angular.module("uigrid").config(e);e.$inject=["formlyConfigProvider"],t.$inject=["$scope"]}(),function(){"use strict";angular.module("media",["ngResource","ui.bootstrap","formly","formlyBootstrap","angularFileUpload","ngSanitize"])}(),function(){"use strict";function e(){function e(e,t,o,n){o.type&&(e.type=o.type),o["class"]&&(e["class"]="class="+o["class"]),o.height&&(e.height=o.height),o.width&&(e.width=o.width);o.source&&(e.source=o.source),o.$observe("source",function(t){t=t.replace("watch?v=","v/"),e.source=t}),o.$observe("type",function(t){e.type=t})}var t={restrict:"E",templateUrl:"app/components/media/media.view.html",replace:!0,scope:{},transclude:!0,controller:e,controllerAs:"media",bindToController:!0};return e.$inject=["$scope","$element","$attrs","$http"],t}angular.module("media").directive("media",e)}(),function(){"use strict";function e(e,o){e.resourceUrlWhitelist(["self","http://youtube.com/**","http://www.youtube.com/**","https://youtube.com/**","https://www.youtube.com/**"]),o.setType({name:"media",templateUrl:"app/components/media/mediaselector.view.html",controller:t}),o.templateManipulators.preWrapper.push(function(e,t,o){return"media"==t.type&&(o.haslabel=!1,o.field=t.key,o.$watch("model",function(e){try{o.mediasource=e[t.key],o.label=t.templateOptions.label}catch(n){}})),e})}function t(e,t){e.field="",e.mediatype=e.options.templateOptions.mediatype?e.options.templateOptions.mediatype:"image",e.chooseMedia=function(n,i){var a={backdrop:!0,keyboard:!0,modalFade:!0,templateUrl:"app/components/media/modalchooser.view.html",controller:o,closeButtonText:"Close",resolve:{mediatype:function(){return i}},actionButtonText:"OK"},l=t.open(a);l.result.then(function(t){e.model[e.field]=t,e.mediasource=t})},e.removeMedia=function(t){e.model[t]="",e.mediasource=""}}function o(e,t,o,n){e.mediatype=n;var i=e.uploader=new o({url:"/upload",queueLimit:1});i.onCompleteItem=function(e,o,n,i){200==n&&o.length>0&&t.close(o[0])},e.closeDialog=function(){t.dismiss("closed")},e.fileSelected=function(e){0!=e.length&&t.close(e)},e.uploadFile=function(){var e=i.queue[0];e.upload()}}angular.module("media").config(e);e.$inject=["$sceDelegateProvider","formlyConfigProvider"],t.$inject=["$scope","$modal"],o.$inject=["$scope","$modalInstance","FileUploader","mediatype"]}(),function(){"use strict";angular.module("login",["ngResource","ui.router","ui.bootstrap"])}(),function(){"use strict";function e(){function e(e,t,o){o.name;o["class"]&&(e["class"]="class="+o["class"]);o.social;o.social&&(e.social=!0);o.signup;o.signup&&(e.signup=!0)}var t={restrict:"E",templateUrl:"app/components/login/login.view.html",replace:!0,controller:e,controllerAs:"login",bindToController:!0};return e.$inject=["$scope","$element","$attrs"],t}angular.module("login").directive("login",e)}(),function(){"use strict";function e(e,t,o,n){e.login=function(){{var t=e.username,o=e.password;n.login({Id:t,Password:o},function(e,t){localStorage.auth=t(pageConf.AuthToken),window.location.href=pageConf.SuccessRedirect},function(e){console.log("err"+e)})}}}var t=angular.module("login");t.controller("LoginController",e),e.$inject=["$scope","$state","$location","AuthenticationService"]}(),function(){"use strict";function e(e){var t=e(pageConf.LocalAuthServer,{},{login:{method:"POST"}});return t}var t=angular.module("login");t.factory("AuthenticationService",e),e.$inject=["$resource"]}(),function(){"use strict";function e(e){e.setType({name:"ckeditor","extends":"textarea"}),e.templateManipulators.preWrapper.push(function(e,t,o){return"ckeditor"==t.type&&(t.ngModelAttrs={ckeditor:{attribute:"ckeditor"}},t.templateOptions||(t.templateOptions={}),t.templateOptions.ckeditor=""),e})}var t=angular.module("entity",["ngResource","ui.bootstrap","formly","formlyBootstrap","ckeditor"]);t.config(e),e.$inject=["formlyConfigProvider"]}(),function(){"use strict";function e(){function e(e,o,n,i){t=e,n["class"]&&(e["class"]="class="+n["class"]);var a;if(!n.name)throw new Error("Server error. Entity name missing");a=n.name;var l=window.pageConf.entities[a];e.entity=l,e.id=n.id,e.submitText="Submit",n.submitText&&(e.submitText=n.submitText);var r=null;n.id&&(r=n.id),e.onSubmit=function(){var t=e.id;null!=t?(e.entity.model.Id=t,i.put(l.url+"/"+t,e.entity.model).then(function(e){console.log(e)},function(e){console.log("error communicating with server")})):i.post(l.url,e.entity.model).then(function(e){console.log(e)},function(e){console.log(e),console.log("error communicating with server")})},null!=r&&n.$observe("id",function(t){return e.id=t,i.get(l.url+"/"+t).then(function(t){var o=t.data;e.entity.model=o},function(e){console.log("error communicating with server")})})}var t,o={restrict:"E",templateUrl:"app/components/entity/entityform.view.html",scope:{},replace:!0,controller:e,controllerAs:"entityform",bindToController:!0};return e.$inject=["$scope","$element","$attrs","$http"],o}angular.module("entity").directive("entityform",e)}(),function(){"use strict";function e(){function e(e,t,o,n){o["class"]&&(e["class"]="class="+o["class"]);var i;if(!o.name)throw new Error("Server error. Entity name missing");i=o.name;var a;if(!o.id)throw new Error("Server error. Entity id missing");a=o.id;var l=window.pageConf.entities[i];o.$observe("id",function(t){return n.get(l.url+"/"+t).then(function(t){e.entitydata=t.data},function(e){console.log("error communicating with server")})})}var t={restrict:"E",templateUrl:"app/components/entity/entity.view.html",scope:{},replace:!0,transclude:!0,controller:e,controllerAs:"entity",bindToController:!0};return e.$inject=["$scope","$element","$attrs","$http"],t}angular.module("entity").directive("entity",e)}(),function(){"use strict";angular.module("actions",["ngResource","ngStorage","ui.router","ui.bootstrap","ui.router.tabs"])}(),function(){"use strict";function e(){function e(e,t,o,n){var i=o.name;o["class"]&&(e["class"]="class="+o["class"]);var a=i;o.view&&(a=o.view);try{var l=new Array,r=window.pageConf.actionset[i];for(var s in r){var c=r[s],u=c.action,p=window.pageConf.actions[u],d={};d.heading=c.label,d.route=u,d.actiontype=p.actiontype,l.push(d)}e.items=l}catch(m){console.log(m)}}var t={restrict:"E",templateUrl:function(e,t){return"menu"===t.widget?"app/components/actions/menu.view.html":"tab"===t.widget?"app/components/actions/tabs.view.html":"pills"===t.widget?"app/components/actions/pills.view.html":"app/components/actions/menu.view.html"},replace:!0,scope:{},controller:e,controllerAs:"actionset",bindToController:!0};return e.$inject=["$scope","$element","$attrs","$templateCache"],t}var t=angular.module("actions");t.directive("actionset",e)}(),function(){"use strict";function e(){function e(e,t,o,n){var i=o.name;o["class"]&&(e["class"]="class="+o["class"]),o.label&&(e.label=o.label),e.route=i,e.action=window.pageConf.actions[i]}var t={restrict:"E",templateUrl:function(e,t){return"button"===t.widget,"app/components/actions/button.view.html"},replace:!0,scope:{},controller:e,controllerAs:"action",bindToController:!0};return e.$inject=["$scope","$element","$attrs","$templateCache"],t}var t=angular.module("actions");t.directive("action",e)}(),function(){"use strict";function e(){var e=angular.injector(["ng"]),o=e.get("$http");if(window.pageConf.AuthRequired){var n=localStorage.auth;n&&null!=n&&n.length>0?t(o,n):window.location.href=window.pageConf.AuthPage}else t(o,"")}function t(e,t){var o=document.location.href,n=o.indexOf("#");n>0&&(o=o.substring(0,n));var i=o+"/conf",a={};return a[window.pageConf.AuthToken]=t,e.get(i,{headers:a}).then(function(e){window.pageConf.partials=e.data.partials?e.data.partials:[],angular.element(document).ready(function(){angular.bootstrap(document,["main"])})},function(e){console.log(e),401==e.status&&(window.location.href=window.pageConf.AuthPage),console.log("error communicating with server")})}angular.module("main",["ngAnimate","ngCookies","ngTouch","ngSanitize","ngResource","ui.router","ui.bootstrap","login","view","actions","entity","media","smart-table","uigrid"]);e(),window.logout=function(){null!=localStorage.auth&&localStorage.auth.length>0&&(localStorage.auth="")}}(),function(){"use strict";function e(e){function t(e,t,o,n){e.params=n;o.name;e.confUrl=document.location.href+"/conf",o["class"]&&(e["class"]="class="+o["class"])}var o={restrict:"E",templateUrl:"app/page.view.html",transclude:!0,replace:!0,controller:t,controllerAs:"page",bindToController:!0};return t.$inject=["$scope","$element","$attrs","$stateParams"],o}angular.module("main").directive("page",e),e.$inject=["$http"]}(),function(){"use strict";function e(e,o){if(window.pageConf.partials)for(var n=0;n<window.pageConf.partials.length;n++){var i=window.pageConf.partials[n];o.put(i.Path,i.Template)}try{var a=window.pageConf.actions;for(var l in a){var r=a[l],s=r.view,c=r.url,u=r.templatepath,p=(r.actiontype,r.viewmode,{});p[s]={templateUrl:u},t.stateProvider.state(l,{url:c,views:p})}t.urlRouteProvider.otherwise("/")}catch(d){console.log(d)}e.debug("runBlock end")}var t=angular.module("main").run(e);e.$inject=["$log","$templateCache"]}(),function(){"use strict";function e(e,t){}angular.module("main").config(e),e.$inject=["$stateProvider","$urlRouterProvider"]}(),function(){"use strict";angular.module("main")}(),function(){"use strict";function e(e,o,n){t.stateProvider=e,t.urlRouteProvider=o;var i=localStorage.auth;i&&null!=i&&i.length>0&&(n.defaults.headers.common[pageConf.AuthToken]=i)}var t=angular.module("main").config(e);e.$inject=["$stateProvider","$urlRouterProvider","$httpProvider"]}(),angular.module("main").run(["$templateCache",function(e){e.put("app/page.view.html",'<div class="container-fluid" ng-transclude=""></div>'),e.put("app/components/actions/button.view.html",'<a class="btn btn-default" ui-sref="{{route}}" role="button">{{label}}</a>'),e.put("app/components/actions/menu.view.html",'<div class="navbar navbar-default"><ul class="nav navbar-nav"><li ng-repeat="item in items"><a ng-href="{{item.route}}" ng-if="item.actiontype!=\'openpartialpage\'">{{item.heading}}</a> <a ui-sref="{{item.heading}}" ng-if="item.actiontype===\'openpartialpage\'">{{item.heading}}</a></li></ul></div>'),e.put("app/components/actions/pills.view.html",'<div class="row"><tabs data="items" type="pills" vertical="true"></tabs></div>'),e.put("app/components/actions/tabs.view.html",'<div class="row"><tabs data="items" type="tabs"></tabs></div>'),e.put("app/components/entity/entity.view.html",'<div ng-transclude=""></div>'),e.put("app/components/entity/entityform.view.html",'<form><div class="row"><formly-form model="entity.model" fields="entity.fields" ng-if="entity.tabbed!=true"></formly-form><tabset ng-if="entity.tabbed==true"><tab ng-repeat="tab in entity.tabs" heading="{{tab.label}}"><formly-form model="entity.model" fields="tab.fields"></formly-form></tab></tabset></div><formly-form model="entity.model"><button type="submit" class="btn btn-default" ng-click="onSubmit()">{{submitText}}</button></formly-form></form>'),e.put("app/components/login/login.view.html",'<div class="container loginbox" ng-controller="LoginController"><div class="row"><div class="main"><h3>Please Log In<div style="display:inline" ng-if="signup">, or <a href="#">Sign Up</a></div></h3><div class="row" ng-if="social"><div class="col-xs-6 col-sm-6 col-md-6"><a href="#" class="btn btn-lg btn-primary btn-block">Facebook</a></div><div class="col-xs-6 col-sm-6 col-md-6"><a href="#" class="btn btn-lg btn-info btn-block">Google</a></div></div><div class="login-or" ng-if="social"><hr class="hr-or"><span class="span-or">or</span></div><form role="form"><div class="form-group"><label for="inputUsernameEmail">Username or email</label> <input type="text" class="form-control" id="inputUsernameEmail" ng-model="username"></div><div class="form-group"><a class="pull-right" href="#">Forgot password?</a> <label for="inputPassword">Password</label> <input type="password" class="form-control" id="inputPassword" ng-model="password"></div><div class="checkbox pull-right"><label><input type="checkbox"> Remember me</label></div><button type="submit" ng-click="login()" class="btn btn btn-primary">Log In</button></form></div></div></div>'),e.put("app/components/media/image.view.html",'<img src="{{source}}">'),e.put("app/components/media/media.view.html",'<div ng-switch="" on="type"><div ng-switch-when="image"><img ng-src="{{source}}" height="{{height}}" width="{{width}}"></div><div ng-switch-when="video"><iframe id="ytplayer" type="text/html" width="{{width}}" height="{{height}}" src="{{source}}" frameborder="0"></iframe></div><div></div></div>'),e.put("app/components/media/mediaselector.view.html",'<div class="clearfix"><label>{{label}}</label><div><media class="pull-left" source="{{mediasource}}" type="{{mediatype}}" height="90px" width="120px"></media><div class="pull-right"><button ng-click="removeMedia(\'{{field}}\')" class="btn btn-default">Remove</button> <button ng-click="chooseMedia(\'{{field}}\', \'{{mediatype}}\')" class="btn btn-default">Choose Media</button></div></div></div>'),e.put("app/components/media/modalchooser.view.html",'<div class="container" style="width:540px; height:500px;"><div class="row"><h2>Choose Media</h2></div><hr><tabset class="row" style="height:250px;"><tab heading="Library"><div style="height:260px"><media type="image" source="http://www.gettyimages.co.uk/gi-resources/images/Homepage/Category-Creative/UK/UK_Creative_462809583.jpg" height="70px"></media><media type="image" source="http://www.gettyimages.co.uk/gi-resources/images/Homepage/Category-Creative/UK/UK_Creative_462809583.jpg" height="70px"></media><media type="image" source="http://www.gettyimages.co.uk/gi-resources/images/Homepage/Category-Creative/UK/UK_Creative_462809583.jpg" height="70px"></media></div><hr><div class="row"><button class="btn btn-primary pull-right" ng-click="fileSelected(\'\')">Select</button> <button class="btn btn-default pull-right" ng-click="closeDialog()">Cancel</button></div></tab><tab heading="URL"><input class="form-control col-sm-4" type="text" placeholder="URL" ng-model="mediaurl"><div class="row col-xs-12" style="height:250px"><media source="{{mediaurl}}" type="{{mediatype}}" height="210px" width="280px" style="margin-top: 15px"></media></div><hr><div class="row"><button class="btn btn-primary pull-right" ng-click="fileSelected(mediaurl)">Select</button> <button class="btn btn-default pull-right" ng-click="closeDialog()">Cancel</button></div></tab><tab heading="Upload"><div class="row" style="height:260px"><div class="col-md-12"><h3>Upload files</h3><div ng-show="uploader.isHTML5"><div nv-file-drop="" uploader="uploader" class="well my-drop-zone">Drop files here</div></div><input type="file" nv-file-select="" uploader="uploader"><table class="table"><thead><tr><th width="40%">Name</th><th ng-show="uploader.isHTML5">Size</th><th ng-show="uploader.isHTML5">Progress</th><th>Status</th><th>Actions</th></tr></thead><tbody><tr ng-repeat="item in uploader.queue"><td><strong>{{ item.file.name }}</strong></td><td ng-show="uploader.isHTML5" nowrap="">{{ item.file.size/1024/1024|number:2 }} MB</td><td ng-show="uploader.isHTML5"><div class="progress" style="margin-bottom: 0;"><div class="progress-bar" role="progressbar" ng-style="{ \'width\': item.progress + \'%\' }"></div></div></td><td class="text-center"><span ng-show="item.isSuccess"><i class="glyphicon glyphicon-ok"></i></span> <span ng-show="item.isCancel"><i class="glyphicon glyphicon-ban-circle"></i></span> <span ng-show="item.isError"><i class="glyphicon glyphicon-remove"></i></span></td><td nowrap=""><button type="button" class="btn btn-warning btn-xs" ng-click="item.cancel()" ng-disabled="!item.isUploading"><span class="glyphicon glyphicon-ban-circle"></span> Cancel</button> <button type="button" class="btn btn-danger btn-xs" ng-click="item.remove()"><span class="glyphicon glyphicon-trash"></span> Remove</button></td></tr></tbody></table></div></div><hr><div class="row"><button class="btn btn-primary pull-right" ng-click="uploadFile()" ng-show="uploader.queue.length"><span class="glyphicon glyphicon-upload"></span> Upload</button> <button class="btn btn-default pull-right" ng-click="closeDialog()">Cancel</button></div></tab></tabset></div>'),e.put("app/components/media/video.view.html",'<ng-video-preview source="youtube" url="{{source}}"></ng-video-preview>'),e.put("app/components/uigrid/multiselect.view.html",'<div class="clearfix"><label>{{label}}</label><table st-table="griditems" class="table table-striped"><thead><tr><th></th><th ng-repeat="column in columns">{{column.name}}</th></tr></thead><tbody><tr ng-repeat="item in griditems"><td><input ng-click="oncheckboxchange($event, \'{{item[valueField]}}\');" ng-model="status[item[valueField]]" type="checkbox"></td><td ng-repeat="column in columns">{{item[column.name]}}</td></tr></tbody></table></div>'),e.put("app/components/view/table.view.html",'<div class="container"><table class="table col-xs-12" ng-transclude=""></table><div class="row col-xs-12" ng-show="editable"><button type="submit" class="btn btn-default" ng-click="onSubmit()">{{submitText}}</button></div></div>'),e.put("app/components/view/ul.view.html",'<div class="container"><ul ng-transclude=""></ul><div class="row col-xs-12" ng-show="editable"><button type="submit" class="btn btn-default" ng-click="onSubmit()">{{submitText}}</button></div></div>'),e.put("app/components/view/view.view.html",'<div><div class="row" ng-transclude=""></div><div class="row" ng-show="editable"><button type="submit" class="btn btn-default" ng-click="onSubmit()">{{submitText}}</button></div></div>')}]);
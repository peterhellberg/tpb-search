'use strict';

angular.module('tpbApp', [
  'ngRoute',
  'tpbApp.controllers',
  'expvar'
]).
config(['$compileProvider', '$routeProvider', '$locationProvider', function($compileProvider, $routeProvider, $locationProvider) {
  $compileProvider.aHrefSanitizationWhitelist(/^\s*(https?|magnet):/);
  $routeProvider.when('/search/', {templateUrl: '/static/partials/search/search.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/term/', {templateUrl: '/static/partials/search/term.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/match/', {templateUrl: '/static/partials/search/match.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/phrase/', {templateUrl: '/static/partials/search/phrase.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/match_phrase/', {templateUrl: '/static/partials/search/match_phrase.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/boolean/', {templateUrl: '/static/partials/search/boolean.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/prefix/', {templateUrl: '/static/partials/search/prefix.html', controller: 'SearchCtrl'});
  $routeProvider.when('/search/debug/', {templateUrl: '/static/partials/debug.html', controller: 'DebugCtrl'});
  $routeProvider.when('/metrics/', {templateUrl: '/static/partials/metrics.html', controller: 'MetricsCtrl'});
  $routeProvider.when('/about/', {templateUrl: '/static/partials/about.html', controller: 'AboutCtrl'});
  $routeProvider.otherwise({redirectTo: '/search/'});
  $locationProvider.html5Mode(true);
}]).filter('encode', function() {
  return window.encodeURIComponent;
});

angular.module('tpbApp.controllers', []).controller('AboutCtrl', ['$scope', function($scope) {}]);

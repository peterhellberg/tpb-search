'use strict';

angular.module('tpbApp', [
  'ngRoute',
  'tpbApp.filters',
  'tpbApp.controllers',
  'expvar'
]).
config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
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
}]);

angular.module('tpbApp.controllers', [])
  .controller('AboutCtrl', ['$scope', function($scope) {
  }]);

angular.module('tpbApp.filters', []).
  filter('interpolate', ['version', function(version) {
    return function(text) {
      return String(text).replace(/\%VERSION\%/mg, version);
    };
  }]);

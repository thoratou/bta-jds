SignCtrl.$inject = ['$scope', '$http', '$rootScope'];

function GameCtrl($scope, $http, $rootScope) {
  $rootScope.gamedata = {};

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    if($rootScope.loggedin){
      $http.get('/games/'+$rootScope.mail).
        success(function(data) {
          console.log($rootScope.gamedata);
          $rootScope.gamedata = data;
        }).error(logError);
    }
  };

  $scope.isHidden = function() {
    return !$rootScope.loggedin;
  }

  $scope.getFirstName = function(id) {
    return $rootScope.gamedata.players[id].firstname;
  }

  $scope.getLastName = function(id) {
    return $rootScope.gamedata.players[id].lastname;
  }

  $scope.getMail = function(id) {
    return $rootScope.gamedata.players[id].mail;
  }

  $scope.isPlayerInGame = function(id, game) {
    return game.players.hasOwnProperty(id);
  }

  $scope.addPlayerToGame = function(id, game) {
    $http.post('/addPlayerToGame', {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromGame = function(id, game) {
    $http.post('/removePlayerFromGame', {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitPlayerGameComment = function(id, game) {
    $http.post('/submitPlayerGameComment', {playerid: id, gameid: game.id, comment: game.players[id].comment}).
      success(function() {
        refresh()
      }).error(logError);
  }

}

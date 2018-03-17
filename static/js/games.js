SignCtrl.$inject = ['$scope', '$http', '$rootScope'];

function initGameData(data) {
  for(var gameid in data.games) {
    var game = data.games[gameid]
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (game.players[i].id === data.currentplayerid) {
        game.players[i].newComment = game.players[i].comment
      }
    }
    game.newpost = ""
  }
  for(var teamid in data.teams) {
    var team = data.teams[teamid]
    team.newName = team.name;
    team.newmanagerid = team.managerid
    team.newcomment = team.comment
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === data.currentplayerid) {
        team.players[i].newComment = team.players[i].comment
      }
    }
    team.newpost = ""
  }

  data.newFirstname = data.players[data.currentplayerid].firstname;
  data.newLastname = data.players[data.currentplayerid].lastname;
}

function GameCtrl($scope, $http, $rootScope) {
  $rootScope.gamedata = {};
  $rootScope.firstname = "";
  $rootScope.lastname = "";

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    if($rootScope.loggedin){
      $http.get('/games/'+$rootScope.mail+'?rnd='+new Date().getTime()).
        success(function(data) {
          $rootScope.gamedata = data;
          initGameData($rootScope.gamedata);
          console.log($rootScope.gamedata);
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

  $scope.getGameStyle = function(id, game) {
    if ($scope.isPlayerInGame(id,game)) {
      return "registeredgame"
    }
    var arrayLength = game.teams.length;
    for (var i = 0; i < arrayLength; i++) {
      if ($scope.isPlayerInTeam(id, game.teams[i])) {
        return "registeredgame";
      }
    }
    return "game";
  }

  $scope.isPlayerInGame = function(id, game) {
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (game.players[i].id === id)
      return true;
    }
    return false;
  }

  $scope.getPlayerInGame = function(id, game) {
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (game.players[i].id === id)
      return game.players[i];
    }
    return null;
  }

  $scope.getTeamStyle = function(id, teamid) {
    if ($scope.isPlayerInTeam(id, teamid)) {
      return "registeredteam";
    }
    return "team";
  }

  $scope.getTeam = function(teamid) {
    return $rootScope.gamedata.teams[teamid];
  }

  $scope.isPlayerInTeam = function(id, teamid) {
    var team = $scope.getTeam(teamid);
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === id)
      return true;
    }
    return false;
  }

  $scope.getPlayerInTeam = function(id, teamid) {
    var team = $scope.getTeam(teamid);
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === id)
      return team.players[i];
    }
    return null;
  }

  $scope.submitPlayerData = function(id, data) {
    if(data.newFirstname !== data.players[id].firstname || data.newLastname !== data.players[id].lastname) {
      $http.post('/submitPlayerData?rnd='+new Date().getTime(), {playerid: id, firstname: data.newFirstname, lastname: data.newLastname}).
        success(function() {
          refresh()
          $("#savetext").css("display","block");
          setTimeout(
            function() {
              $("#savetext").css("display","none");
            }, 4000);
        }).error(logError);
    }
  }

  $scope.addPlayerToGame = function(id, game) {
    $http.post('/addPlayerToGame?rnd='+new Date().getTime(), {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromGame = function(id, game) {
    $http.post('/removePlayerFromGame?rnd='+new Date().getTime(), {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitPlayerGameComment = function(id, game) {
    $http.post('/submitPlayerGameComment?rnd='+new Date().getTime(), {playerid: id, gameid: game.id, comment: $scope.getPlayerInGame(id, game).newComment}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.addTeamToGame = function(teamname, managerid, game) {
    if(teamname == undefined || teamname === "") {
      //todo error handling
      return;
    }
    $http.post('/addTeamToGame?rnd='+new Date().getTime(), {teamname: teamname, managerid: managerid, gameid: game.id}).
      success(function() {
        refresh()
        //reduce team creation panel
        $("#collapse_team_add_"+game.id).collapse("hide");
      }).error(logError);
  }

  $scope.removeTeamFromGame = function(teamid, managerid, game) {
    $http.post('/removeTeamFromGame?rnd='+new Date().getTime(), {teamid: teamid, managerid: managerid, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.changeTeamName = function(teamid, teamname, managerid) {
    if(teamname == undefined || teamname === "") {
      //todo error handling
      return;
    }
    $http.post('/changeTeamName?rnd='+new Date().getTime(), {teamid: teamid, teamname: teamname, managerid: managerid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.changeManager = function(teamid, oldmanagerid, newmanagerid) {
    if(oldmanagerid === newmanagerid) {
      //todo error handling
      return;
    }
    $http.post('/changeManager?rnd='+new Date().getTime(), {teamid: teamid, oldmanagerid: oldmanagerid, newmanagerid: newmanagerid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitTeamComment = function(teamid, newcomment, managerid) {
    $http.post('/submitTeamComment?rnd='+new Date().getTime(), {teamid: teamid, newcomment: newcomment, managerid: managerid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.addPlayerToTeam = function(id, teamid) {
    $http.post('/addPlayerToTeam?rnd='+new Date().getTime(), {playerid: id, teamid: teamid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromTeam = function(id, teamid) {
    $http.post('/removePlayerFromTeam?rnd='+new Date().getTime(), {playerid: id, teamid: teamid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitPlayerTeamComment = function(id, teamid) {
    $http.post('/submitPlayerTeamComment?rnd='+new Date().getTime(), {playerid: id, teamid: teamid, comment: $scope.getPlayerInTeam(id, teamid).newComment}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitTeamNewPost = function(teamid, newpost, playerid) {
    $http.post('/submitTeamNewPost?rnd='+new Date().getTime(), {teamid: teamid, newpost: newpost, playerid: playerid}).
      success(function() {
        refresh()
      }).error(logError);
  }

}

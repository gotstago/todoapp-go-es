angular.module('todoapp', ['ngRoute', 'ngWebSocket']).config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/todo/create', {
            templateUrl: 'scripts/todo/create.html',
            controller: ['$scope', '$http', '$location',
                function($scope, $http, $location) {
                    $scope.todo = {}
                    $scope.create = function() {
                        $http.post('/cmd/', {
                            name: 'createTodoItem',
                            data: $scope.todo,
                        }).success(function() {
                            $location.path('/todo')
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
        }).when('/todo/:todoid', {
            templateUrl: 'scripts/todo/show.html',
            controller: ['$scope', 'todo',
                function($scope, todo) {
                    $scope.todo = todo.data
                }
            ],
            resolve: {
                todo: ['$http', '$route',
                    function($http, $route) {
                        id = $route.current.params.todoid
                        return $http.get('/api/todo/' + id + '.json')
                    }
                ]
            }
        }).when('/todo/:todoid/edit', {
            templateUrl: 'scripts/todo/edit.html',
            controller: ['$scope', '$http', '$location', 'todo',
                function($scope, $http, $location, todo) {
                    $scope.todo = todo.data
                    $scope.edit = function() {
                        $http.post('/cmd/', {
                            name: 'updateTodoItem',
                            data: $scope.todo,
                        }).success(function() {
                            $location.path('/todo')
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
            resolve: {
                todo: ['$http', '$route',
                    function($http, $route) {
                        id = $route.current.params.todoid
                        return $http.get('/api/todo/' + id + '.json')
                    }
                ]
            }
        }).when('/todo', {
            templateUrl: 'scripts/todo/all.html',
            controller: ['$scope', '$http', '$location', 'todos', 'todoappws',
                function($scope, $http, $location, todos, todoappws) {
                    $scope.todos = todos.data
                    todoappws.on('todoItemCreated', function(event) {
                        $scope.todos[event.data.id] = event.data;
                    })
                    todoappws.on('todoItemUpdated', function(event) {
                        $scope.todos[event.data.id] = event.data
                    })
                    todoappws.on('todoItemRemoved', function(event) {
                        delete $scope.todos[event.data]
                    })
                    $scope.show = function(todo) {
                        $location.path('/todo/' + todo.id)
                    }
                    $scope.edit = function(todo) {
                        $location.path('/todo/' + todo.id + '/edit')
                    }
                    $scope.remove = function(todo) {
                        $http.post('/cmd/', {
                            name: 'removeTodoItem',
                            data: todo.id,
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                    $scope.toggleCompleted = function(todo) {
                        todo = angular.copy(todo)
                        todo.completed = !todo.completed
                        $http.post('/cmd/', {
                            name: 'updateTodoItem',
                            data: todo,
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
            resolve: {
                todos: ['$http',
                    function($http) {
                        return $http.get('/api/todo/all.json')
                    }
                ]
            }
        }).when('/game/create', {
            templateUrl: 'scripts/game/create.html',
            controller: ['$scope', '$http', '$location',
                function($scope, $http, $location) {
                    $scope.game = {}
                    $scope.create = function() {
                        $http.post('/cmd/', {
                            name: 'createGameItem',
                            data: $scope.game,
                        }).success(function() {
                            $location.path('/game')
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
        }).when('/game/:gameid', {
            templateUrl: 'scripts/game/show.html',
            controller: ['$scope', 'game',
                function($scope, game) {
                    $scope.game = game.data
                }
            ],
            resolve: {
                game: ['$http', '$route',
                    function($http, $route) {
                        id = $route.current.params.gameid
                        return $http.get('/api/game/' + id + '.json')
                    }
                ]
            }
        }).when('/game/:gameid/edit', {
            templateUrl: 'scripts/game/edit.html',
            controller: ['$scope', '$http', '$location', 'game',
                function($scope, $http, $location, game) {
                    $scope.game = game.data
                    $scope.edit = function() {
                        $http.post('/cmd/', {
                            name: 'updateGameItem',
                            data: $scope.game,
                        }).success(function() {
                            $location.path('/game')
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
            resolve: {
                game: ['$http', '$route',
                    function($http, $route) {
                        id = $route.current.params.gameid
                        return $http.get('/api/game/' + id + '.json')
                    }
                ]
            }
        }).when('/game', {
            templateUrl: 'scripts/game/all.html',
            controller: ['$scope', '$http', '$location', 'games', 'todoappws',
                function($scope, $http, $location, games, todoappws) {
                    $scope.games = games.data
                    todoappws.on('gameItemCreated', function(event) {
                        $scope.games[event.data.id] = event.data;
                    })
                    todoappws.on('gameItemUpdated', function(event) {
                        $scope.games[event.data.id] = event.data
                    })
                    todoappws.on('gameItemRemoved', function(event) {
                        delete $scope.games[event.data]
                    })
                    $scope.show = function(game) {
                        $location.path('/game/' + game.id)
                    }
                    $scope.edit = function(game) {
                        $location.path('/game/' + game.id + '/edit')
                    }
                    $scope.remove = function(game) {
                        $http.post('/cmd/', {
                            name: 'removeGameItem',
                            data: game.id,
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                    $scope.toggleCompleted = function(game) {
                        game = angular.copy(game)
                        game.completed = !game.completed
                        $http.post('/cmd/', {
                            name: 'updateGameItem',
                            data: game,
                        }).error(function() {
                            alert('Something wen\'t awry')
                        })
                    }
                }
            ],
            resolve: {
                games: ['$http',
                    function($http) {
                        return $http.get('/api/game/all.json')
                    }
                ]
            }
        })
    }
]).run(['todoappws','$rootScope',
    function(todoappws,$rootScope) {
        todoappws.on('newWebSocketClient', function(event) {
            $rootScope.numWsClients = event.data
        })
    }
]);
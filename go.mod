module github.com/hedwig100/todo-app

go 1.16

require github.com/hedwig100/todo-app/cmd/app v0.0.0-00010101000000-000000000000

replace github.com/hedwig100/todo-app/cmd/app => ./cmd/app

replace github.com/hedwig100/todo-app/internal/data => ./internal/data

replace github.com/hedwig100/todo-app/internal/route => ./internal/route

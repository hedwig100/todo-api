module github.com/hedwig100/todo-app/cmd/app

go 1.16

replace github.com/hedwig100/todo-app/internal/route => ../../internal/route

replace github.com/hedwig100/todo-app/internal/data => ../../internal/data

require github.com/hedwig100/todo-app/internal/route v0.0.0-00010101000000-000000000000

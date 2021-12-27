module github.com/hedwig100/todo-app/internal/route // REVIEW:Goのmoduleパスこれでいいのか

go 1.16

replace github.com/hedwig100/todo-app/internal/data => ../data

require github.com/hedwig100/todo-app/internal/data v0.0.0-00010101000000-000000000000

# API 

## users 

`/users` <br>

- POST
    - ユーザを登録する
    - 要求

    ```
    {
        "username":username, 
        "password":password
    }
    ```

`/users/{username}` <br>

- DELETE
    - usernameのuserを消す

`/users/{username}/login` <br> 

- POST 
    - ログインできたか
    - 要求

    ```
    {
        "username":username,
        "password":password
    }
    ```

    - レスポンス
        - ステータスコード
        ```

`/users/{username}/task-lists` <br> 

- GET
    - usernameのユーザのtaskListsの一覧
    - レスポンス

    ```
    {
        "taskLists" [
            {
                "icon":icon
                "listname":listname,
                "listId":listid,
            },
            ...
        ]
    }
    ```
    
## task-lists

`/task-lists/`

- POST
    - taskListを作成する
    - 要求

    ```
    {
        "username":username,
        "icon":icon,
        "listname":listname,
    }
    ```

    - レスポンス

    ```
    {
        "listId":listid,
    }
    ```

`/task-lists/{listId}`

- GET
    - listIdのリストを手に入れる
    - レスポンス

    ```
    {
        "username":username,
        "icon":icon,
        "listname":listname,
        "tasks" [
            {
                "taskname":taskname,
                "deadline":deadline,
                "isDone":false,
                "isImportant":true,
                "memo":memo
            },
            ...
        ]
    }
    ```

- PUT
    - listIdのリストの更新
    - 要求
    ```
    {
        "username":username,
        "icon":icon,
        "listname":listname,
    }
    ```

- DELETE
    - listIdのリストを消す

## tasks  

`/tasks` <br>

- POST
    - taskを追加する
    - 要求

    ```
    {
        "username":username,
        "listId":listid,
        "taskname":taskname,
        "deadline":deadline
    }
    ```

    - レスポンス

    ```
    {
        "taskId":taskid
    }
    ```

`/tasks/{taskId}` <br>

- GET
    - task_idのtaskの情報を手に入れる
    - レスポンス

    ```
    {
        "username":username,
        "listId":listid,
        "taskname":taskname,
        "deadline":deadline,
        "isDone":false,
        "isImportant":true,
        "memo":memo
    }
    ```

- PUT 
    - task_idのtaskの情報を更新する
    - 要求

    ```
    {
        "username":username,
        "listId":listid,
        "taskname":taskname,
        "deadline":deadline,
        "isDone":false,
        "isImportant":true,
        "memo":memo
    }
    ```

- DELETE
    - task_idのtaskを消す

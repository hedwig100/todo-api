# API 

## users 

`/users` <br>

- POST
    - ユーザを登録する
    - リクエスト

        ```
        {
            "username":username, 
            "password":password
        }
        ```

    - レスポンス
        - 成功したら201でbodyはなし
        - 失敗したら500エラー <!--REVIEW: 500エラーでいいか,RFCを読む?-->

            ```
            {
                "errorMessage"...,
            }
            ```

- DELETE
    - usernameのuserを消す
    - リクエスト

        ```
        {
            "username":username
            "password":password
        }
        ```

    - レスポンス
        - 成功したら200でbodyはなし
        - 失敗したら500エラー <!--REVIEW: 500エラーでいいか,RFCを読む?-->

            ```
            {
                "errorMessage"...,
            }
            ```

`/users/login` <br> 

- POST 
    - ログインできたか
    - リクエスト

        ```
        {
            "username":username,
            "password":password
        }
        ```

    - レスポンス
        - 201でログイン成功
        - 500でログイン失敗 <!--REVIEW: 500エラーでいいか,RFCを読む?-->
            ```
            {
                "errorMessage"...,
            }
            ```

`/users/task-lists/{username}` <br> 

- GET
    - usernameのユーザのtaskListsの一覧
    - リクエスト

        ```
        {
            "password": password
        }
        ```

    - レスポンス
        - 成功したらステータスコードは200で以下を返す
            ```
            {
                "taskLists" [
                    {
                        "username":userrname,
                        "icon":icon,
                        "listname":listname,
                        "listId":listid,
                    },
                    ...
                ]
            }
            ```
        - 失敗したら500エラー 

            ```
            {
                "errorMessage"...,
            }
            ```

    
## task-lists

`/task-lists`

- POST
    - taskListを作成する
    - リクエスト

        ```
        {
            "username":username,
            "password":password,
            "icon":icon,
            "listname":listname,
        }
        ```

    - レスポンス
        - 成功したら201をステータスコードとし、listIdを返す
            ```
            {
                "listId":listid,
            }
            ```
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
                {
                    "errorMessage"...,
                }
            ```

`/task-lists/{listId}`

- GET
    - listIdのリストを手に入れる
    - リクエスト

        ```
        {
            "username": username, 
            "password": password
        }
        ```

    - レスポンス
        - 成功したら200を返して, 以下のようなjsonを返す

            ```
            {
                "username":username,
                "icon":icon,
                "listname":listname,
                "tasks" [
                    {
                        "taskname":taskname,
                        "deadline":deadline,
                        "taskId":taskId,
                        "isDone":false,
                        "isImportant":true,
                        "memo":memo
                    },
                    ...
                ]
            }
            ```

        - 失敗したらステータスコードは500とし、エラーを返す
            ```
                {
                    "errorMessage"...,
                }
            ```


- PUT
    - listIdのリストの更新
    - リクエスト
        ```
        {
            "username":username,
            "password":password,
            "icon":icon,
            "listname":listname,
        }
        ```
    - レスポンス
        - 成功したら201でボディはなし
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
                {
                    "errorMessage"...,
                }
            ```

- DELETE
    - listIdのリストを消す
    - リクエスト
        ```
        {
            "username":username,
            "password":password
        }
        ```
    - レスポンス
        - 成功したら201でボディはなし
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
                {
                    "errorMessage"...,
                }
            ```

## tasks  

`/tasks` <br>

- POST
    - taskを追加する
    - リクエスト
        ```
        {
            "username":username,
            "password":password,
            "listId":listid,
            "taskname":taskname,
            "deadline":deadline
        }
        ```

    - レスポンス
        - 成功したら201でtaskIdを返す
            ```
            {
                "taskId":taskid
            }
            ```
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
            {
                "errorMessage"...,
            }
            ```

`/tasks/{taskId}` <br>

- GET
    - task_idのtaskの情報を手に入れる
    - リクエスト
        ```
        {
            "username": username,
            "password": password 
        }
        ```
    - レスポンス
        - 成功したら200で以下を返す
            ```
            {
                "listId":listid,
                "taskname":taskname,
                "deadline":deadline,
                "isDone":false,
                "isImportant":true,
                "memo":memo
            }
            ```
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
            {
                "errorMessage"...,
            }
            ```
    

- PUT 
    - task_idのtaskの情報を更新する
    - リクエスト
        ```
        {
            "username":username,
            "password":password,
            "listId":listid,
            "taskname":taskname,
            "deadline":deadline,
            "isDone":false,
            "isImportant":true,
            "memo":memo
        }
        ```
    - レスポンス
        - 成功したら201
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
            {
                "errorMessage"...,
            }
            ```

- DELETE
    - task_idのtaskを消す
    - リクエスト
        ```
        {
            "username":username,
            "password":password,
        }
        ```
    - レスポンス
        - 成功したら201
        - 失敗したらステータスコードは500とし、エラーを返す
            ```
            {
                "errorMessage"...,
            }
            ```
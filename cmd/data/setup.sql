DROP TABLE users CASCADE; 
DROP TABLE task_lists CASCADE;
DROP TABLE tasks CASCADE;

CREATE TABLE users (
    uuid SERIAL PRIMARY KEY, 
    username varchar(64) NOT NULL UNIQUE, 
    password bytea NOT NULL
); 

CREATE TABLE task_lists (
    list_id SERIAL PRIMARY KEY,
    username varchar(64) REFERENCES users (username) ON DELETE CASCADE,
    icon varchar(64) NOT NULL,
    listname varchar(64) NOT NULL
); 

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY, 
    username varchar(64) REFERENCES users (username) ON DELETE CASCADE,
    list_id int REFERENCES task_lists ON DELETE CASCADE,
    taskname varchar(64) NOT NULL, 
    deadline date, 
    is_done boolean DEFAULT false, 
    is_important boolean DEFAULT false,
    memo text DEFAULT ''
); 
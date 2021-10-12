drop table task; 

CREATE TABLE task (
    id serial primary key, 
    taskname varchar(255) NOT NULL,
    deadline timestamp, 
    isdone boolean,
    donetime timestamp
)
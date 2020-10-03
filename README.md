```sh
# from project root
cd chitchat
go run .
```

```sh
# on local machine (from project root)
cd .devcontainer
docker container exec -it <database_container_name> sh

# in the postgres container
psql --username=$POSTGRES_USER --dbname=$POSTGRES_DB

\l                      # list databases
\c <database_name>      # connect to a database
\dt                     # list tables in the database
\d <table_name>         # describe a table
\g                      # run the previous command
\timing                 # turn on timing
\h DROP TABLE           # get help of a statement e.g. DROP TABLE
```

```sql
create database chitchat;

create table users (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    name       varchar(255),
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp not null
);

create table sessions (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    email      varchar(255),
    user_id    integer references users(id),
    created_at timestamp not null
);

create table threads (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    topic      text,
    user_id    integer references users(id),
    created_at timestamp not null
);

create table posts (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    body       text,
    user_id    integer references users(id),
    thread_id  integer references threads(id),
    created_at timestamp not null
);
```

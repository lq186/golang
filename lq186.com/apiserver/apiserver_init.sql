/** create database */
create database golang_api character set utf8mb4 collate utf8mb4_general_ci;
/** create user */
create user 'golang_api'@'%' identified by 'golang_api_123';
grant all on golang_api.* to golang_api;
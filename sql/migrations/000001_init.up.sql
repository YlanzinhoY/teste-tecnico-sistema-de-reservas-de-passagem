create table if not exists management_route(
    id varchar primary key,
    route_name varchar,
    origin varchar(50),
    destination varchar(50)
);
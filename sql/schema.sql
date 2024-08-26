create table management_route(
    id varchar primary key,
    route_name varchar,
    origin varchar(50),
    destination varchar(50)
);

create table management_travel(
    management_travel_id uuid DEFAULT uuid_generate_v4(),
    management_routes_id varchar not null,
    ticket_price DOUBLE PRECISION not null,
    total_seats integer not null,
    travel_start timestamp not null,
    travel_finish timestamp not null,
    travel_company varchar(100) not null,

    PRIMARY KEY(management_travel_id),
    CONSTRAINT fk_management_routes
        FOREIGN KEY(management_routes_id)
            references management_route(id)
);
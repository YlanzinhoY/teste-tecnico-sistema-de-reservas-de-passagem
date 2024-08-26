create table if not exists management_travel(
    management_travel_id uuid DEFAULT uuid_generate_v4(),
    management_routes_id varchar,
    ticket_price DOUBLE PRECISION not null,
    total_seats integer not null,
    travel_start timestamp,
    travel_finish timestamp,
    travel_company varchar(100) not null,

PRIMARY KEY(management_travel_id),
CONSTRAINT fk_management_routes
    FOREIGN KEY(management_routes_id)
    references management_route(id)

);
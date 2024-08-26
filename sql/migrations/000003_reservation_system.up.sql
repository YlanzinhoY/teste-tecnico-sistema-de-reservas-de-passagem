CREATE TABLE IF NOT EXISTS reservation_system(
  reservation_id uuid default uuid_generate_v4(),
  management_travel_id uuid,
  passenger_name varchar(100) not null,
  seat_number int not null,

    PRIMARY KEY (reservation_id),
    CONSTRAINT reservation_system_unique_seat UNIQUE(reservation_id, seat_number),
    CONSTRAINT fk_management_travel
        FOREIGN KEY(management_travel_id) REFERENCES management_travel(management_travel_id)
);
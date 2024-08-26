-- name: GetRouteById :one
select * from management_route
where id = $1 limit 1;

-- name: GetManagementRouteAll :many
select * from management_route ORDER by id;

-- name: CreateManagementRoute :exec
INSERT INTO management_route (
    id,
    route_name,
    origin,
    destination
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING id, route_name, origin, destination;

-- name: UpdateManagementRoute :one
update management_route
    set route_name =$2, origin = $3,destination = $4
where id = $1
returning *;

-- name: DeleteManagementRoute :exec

delete from management_route
where id = $1;
-------------------------------------------------------------------------

-- name: CreateManagementTravel :exec
insert into management_travel(
    management_travel_id,
    management_routes_id,
    ticket_price,
    total_seats,
    travel_start,
    travel_finish,
    travel_company
) values (
$1,
$2,
$3,
$4,
$5,
$6,
          $7
) RETURNING management_routes_id, management_travel_id, ticket_price, total_seats, travel_start, travel_finish, travel_company;

-- name: GetManagementTravelById :one
SELECT * FROM management_travel WHERE management_travel_id = $1;

-- name: GetAllManagementTravel :many
SELECT * FROM management_travel;

-- name: PutManagementTravel :one
UPDATE management_travel
    SET ticket_price = $2, total_seats = $3, travel_start = $4, travel_finish = $4, travel_company = $5
WHERE management_travel_id = $1
RETURNING *;

-- name: DeleteManagementTravel :exec
DELETE from management_travel
WHERE management_travel_id = $1;
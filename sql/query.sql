-- name: GetRoute :one
select * from management_route
where id = $1 limit 1;

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

-- name: UpdateManagementRoute :exec
update management_route
    set route_name =$2, origin = $3,destination = $4
where id = $1;
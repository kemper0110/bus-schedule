BEGIN;

CREATE TABLE carriers(
    id SERIAL PRIMARY KEY,
    name varchar(50)
);

CREATE TABLE bus(
    id SERIAL PRIMARY KEY,
    model varchar(50),
    plate_number varchar(12),
    scheme varchar(100),
    carrier_id integer REFERENCES carriers(id) ON DELETE CASCADE
);

CREATE TABLE main_route(
    id SERIAL PRIMARY KEY,
    bus_id integer REFERENCES bus(id) ON DELETE CASCADE,
    carrier_id integer REFERENCES carriers(id) ON DELETE CASCADE,
    name varchar (60)
);

CREATE TABLE mid_route(
    id SERIAL PRIMARY KEY,
    from_time timestamp,
    to_time timestamp,
    price decimal,
    rating decimal,
    main_id integer REFERENCES main_route(id) ON DELETE CASCADE,
    from_city integer REFERENCES city(id) ON DELETE CASCADE,
    to_city integer REFERENCES city(id) ON DELETE CASCADE
);


COMMIT;




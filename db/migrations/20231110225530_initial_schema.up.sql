CREATE TABLE `space_crafts` (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    class varchar(255) NOT NULL,
    crew int unsigned not null,
    image varchar(255) NOT NULL,
    value numeric(20,5),
    status varchar(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
    updated_at timestamp NOT NULL DEFAULT current_timestamp(),
    is_deleted bool NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id)
);

CREATE TABLE `armaments` (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
    updated_at timestamp NOT NULL DEFAULT current_timestamp(),
    is_deleted bool NOT NULL DEFAULT FALSE,
    PRIMARY KEY(id)
);

CREATE TABLE `space_crafts_armaments` (
    id int NOT NULL AUTO_INCREMENT,
    space_craft_id int NOT NULL,
    armament_id int NOT NULL,
    qty int unsigned not null,
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
    updated_at timestamp NOT NULL DEFAULT current_timestamp(),
    is_deleted bool NOT NULL DEFAULT FALSE,
    PRIMARY KEY(id),
    FOREIGN KEY (space_craft_id) REFERENCES space_crafts(id),
    FOREIGN KEY (armament_id) REFERENCES armaments(id)
);


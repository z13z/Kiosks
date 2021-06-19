CREATE TYPE KioskImageState AS ENUM ('created', 'waiting', 'building', 'failed', 'done');
CREATE TYPE KioskUserPermission AS ENUM ('users', 'images', 'kiosks');

CREATE TABLE KioskImage
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128) UNIQUE NOT NULL,
    create_time TIMESTAMP           NOT NULL,
    script      VARCHAR             NOT NULL,
    state       KioskImageState,
    PRIMARY KEY (id)
);

CREATE TABLE Kiosk
(
    id             BIGINT GENERATED ALWAYS AS IDENTITY,
    address        VARCHAR(64) NOT NULL,
    kiosk_image_id BIGINT,
    last_online    TIMESTAMP   NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fkIdKioskImage_id FOREIGN KEY (kiosk_image_id) REFERENCES KioskImage (id)
);

CREATE TABLE KioskUser
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128) UNIQUE       NOT NULL,
    update_time TIMESTAMP                 NOT NULL,
    permissions KioskUserPermission ARRAY NOT NULL,
    password    VARCHAR(512),
    PRIMARY KEY (id)
);

INSERT INTO KioskUser(name, update_time, permissions, password)
VALUES ('zaza13', now(), '{kiosks,images,users}', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918')
CREATE TYPE KioskImageState AS ENUM ('created', 'waiting', 'done');
CREATE TYPE KioskUserPermission AS ENUM ('users', 'images', 'kiosks');

CREATE TABLE KioskImage
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128) NOT NULL,
    create_time TIMESTAMP    NOT NULL,
    script      VARCHAR      NOT NULL,
    state       KioskImageState,
    PRIMARY KEY (id)
);

CREATE TABLE Kiosk
(
    id             BIGINT GENERATED ALWAYS AS IDENTITY,
    name           VARCHAR(128) NOT NULL,
    kiosk_image_id BIGINT,
    create_time    TIMESTAMP    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fkIdKioskImage_id FOREIGN KEY (kiosk_image_id) REFERENCES KioskImage (id)
);

CREATE TABLE KioskUser
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(128)              NOT NULL,
    update_time TIMESTAMP                 NOT NULL,
    permissions KioskUserPermission ARRAY NOT NULL,
    password    VARCHAR(512),
    PRIMARY KEY (id)
);
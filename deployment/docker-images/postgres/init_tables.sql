CREATE TABLE KioskImage(
                           id BIGINT GENERATED ALWAYS AS IDENTITY,
                           name VARCHAR(128) NOT NULL,
                           create_time TIMESTAMP NOT NULL,
                           PRIMARY KEY(id)
);

CREATE TABLE Kiosk(
                      id BIGINT GENERATED ALWAYS AS IDENTITY,
                      name VARCHAR(128) NOT NULL,
                      kiosk_image_id BIGINT,
                      create_time TIMESTAMP NOT NULL,
                      PRIMARY KEY(id),
                      CONSTRAINT fkIdKioskImage_id FOREIGN KEY (kiosk_image_id) REFERENCES KioskImage(id)
);

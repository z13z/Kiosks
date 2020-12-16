CREATE TABLE KioskImage(
                           id BIGINT GENERATED ALWAYS AS IDENTITY,
                           name VARCHAR(128) NOT NULL,
                           PRIMARY KEY(id)
);

CREATE TABLE Kiosk(
                      id BIGINT GENERATED ALWAYS AS IDENTITY,
                      name VARCHAR(128) NOT NULL,
                      kiosk_image_id BIGINT,
                      PRIMARY KEY(id),
                      CONSTRAINT fkIdKioskImage_id FOREIGN KEY (kiosk_image_id) REFERENCES KioskImage(id)
);

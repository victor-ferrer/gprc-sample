CREATE TABLE Ticket (
    ID SERIAL PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    PurchaseDate TIMESTAMP NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    Currency VARCHAR(10) NOT NULL,
    Labels TEXT,
    File VARCHAR(255)
);
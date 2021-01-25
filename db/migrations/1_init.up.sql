CREATE TABLE Accounts (
                          Account_ID INTEGER auto_increment primary key,
                          Document_Number varchar(100) NOT NULL
);

CREATE TABLE OperationsTypes (
                                 OperationsType_ID INTEGER auto_increment primary key,
                                 Description varchar(200) NOT NULL,
                                 OperationsType varchar(200) NOT NULL
);

CREATE TABLE Transactions (
                              Transaction_ID INTEGER auto_increment primary key,
                              Account_ID INTEGER NOT NULL,
                              OperationsType_ID INTEGER NOT NULL,
                              Amount DOUBLE NOT NULL,
                              EventDate DATE
);

ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_Account FOREIGN KEY (Account_ID) REFERENCES Accounts(Account_ID);
ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_OperationsType FOREIGN KEY (OperationsType_ID) REFERENCES OperationsTypes(OperationsType_ID);
ALTER TABLE Accounts ADD CONSTRAINT Accounts_UN UNIQUE KEY (Document_Number);

INSERT INTO OperationsTypes (OperationsType_ID, Description, OperationsType) VALUES
(1, 'COMPRA A VISTA','DEBIT'),
(2, 'COMPRA PARCELADA','DEBIT'),
(3, 'SAQUE','DEBIT'),
(4, 'PAGAMENTO','CREDIT');
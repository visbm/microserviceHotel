ALTER TABLE SEAT
DROP COLUMN      rent_from   CASCADE,   
DROP COLUMN      rent_to     CASCADE,
ADD COLUMN       BOOLEAN     NOT NULL;
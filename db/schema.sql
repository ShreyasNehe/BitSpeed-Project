CREATE TABLE IF NOT EXISTS contact (
        id SERIAL PRIMARY KEY,
        phoneNumber VARCHAR(255),
        email VARCHAR(255),
        linkedId INT,
        linkPrecedence VARCHAR(10) CHECK (linkPrecedence IN ('primary', 'secondary')),
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deletedAt TIMESTAMP
    );

CREATE INDEX IF NOT EXISTS idx_primary_contact ON contact (linkPrecedence) WHERE linkPrecedence = 'primary';
CREATE INDEX IF NOT EXISTS idx_secondary_contact ON contact (linkPrecedence) WHERE linkPrecedence = 'secondary';

ALTER TABLE contact ADD CONSTRAINT fk_linked_contact FOREIGN KEY (linkedId) REFERENCES contact(id) ON DELETE CASCADE;

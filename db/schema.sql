CREATE TABLE IF NOT EXISTS contact (
        id SERIAL PRIMARY KEY,
        phone_number VARCHAR(255),
        email VARCHAR(255),
        linked_id INT,
        link_precedence VARCHAR(10) CHECK (link_precedence IN ('primary', 'secondary')),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP
    );

CREATE INDEX IF NOT EXISTS idx_primary_contact ON contact (linkPrecedence) WHERE linkPrecedence = 'primary';
CREATE INDEX IF NOT EXISTS idx_secondary_contact ON contact (linkPrecedence) WHERE linkPrecedence = 'secondary';

ALTER TABLE contact ADD CONSTRAINT fk_linked_contact FOREIGN KEY (linkedId) REFERENCES contact(id) ON DELETE CASCADE;

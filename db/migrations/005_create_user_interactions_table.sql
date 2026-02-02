CREATE TABLE IF NOT EXISTS user_interactions (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    book_id VARCHAR(255) NOT NULL,
    interaction_type VARCHAR(50) NOT NULL, -- e.g., 'view', 'like', 'read', 'borrow'
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_book
        FOREIGN KEY(book_id)
        REFERENCES books(id)
        ON DELETE CASCADE
);

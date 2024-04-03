CREATE TABLE IF NOT EXISTS project (
    project_id SERIAL PRIMARY KEY,
    project_name TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS project_member (
    member_id INT,
    project_id INT,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (member_id, project_id),
    FOREIGN KEY (member_id) REFERENCES member(member_id)
);
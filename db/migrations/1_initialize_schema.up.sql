CREATE TABLE IF NOT EXISTS user (
    user_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_credentials (
    user_id INT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    password_salt TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);

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

CREATE TABLE IF NOT EXISTS project_user (
    user_id INT,
    project_id INT,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, project_id)
    FOREIGN KEY (user_id) REFERENCES user(user_id),
);

CREATE TABLE IF NOT EXISTS tasks_group (
    tasks_group_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    group_name TEXT NOT NULL,
    description TEXT,
    Status       string    `validate:"required" json:"status" mapstructure:"status"`
	Priority     string    `validate:"required" json:"priority" mapstructure:"priority"`
	StartDate    time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"start_date" mapstructure:"start_date"`
	EndDate      time.Time `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"end_date" mapstructure:"end_date"`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (project_id) REFERENCES project(project_id)
);

CREATE TABLE IF NOT EXISTS task (
    task_id SERIAL PRIMARY KEY,
    tasks_group_id INT NOT NULL,
    task_name TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (task_group_id) REFERENCES tasks_group(task_group_id)
);

CREATE TABLE IF NOT EXISTS tasks_group_user (
    user_id INT,
    tasks_group_id INT,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, task_group_id)
    FOREIGN KEY (user_id) REFERENCES user(user_id),
);

CREATE TABLE IF NOT EXISTS task_user (
    user_id INT,
    task_id INT,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, task_id)
    FOREIGN KEY (user_id) REFERENCES user(user_id),
);
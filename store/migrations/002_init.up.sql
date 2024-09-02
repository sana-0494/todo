CREATE TABLE IF NOT EXISTS todo (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title text NOT NULL,
    status text NOT NULL,
    created_at timestamp with time zone default now()
    );




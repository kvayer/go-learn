CREATE TABLE tariffs (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    name TEXT NOT NULL,
    code INTEGER NOT NULL UNIQUE,
    price NUMERIC(12,2) CHECK (price > 0),
    speed_mbps INTEGER NOT NULL CHECK(speed_mbps > 0),
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE lessons (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    title TEXT NOT NULL,
    access_code TEXT NOT NULL UNIQUE,
    duration_minutes INTEGER NOT NULL CHECK (duration_minutes > 0),
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE departments (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE employees (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    name TEXT NOT NULL,
    department_id UUID NOT NULL REFERENCES departments(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE folders (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE files (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    name TEXT NOT NULL,
    folder_id UUID NOT NULL REFERENCES folders(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE support_agents (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE tickets (
    id UUID DEFAULT uuidv7() PRIMARY KEY,
    subject TEXT NOT NULL,
    support_agent_id UUID REFERENCES support_agents(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);
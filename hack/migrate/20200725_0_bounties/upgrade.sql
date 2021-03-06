CREATE TABLE bounties (
    id bigint DEFAULT comunion.id_generator() NOT NULL
        CONSTRAINT bounties_id_pk
            PRIMARY KEY,
    startup_id bigint NOT NULL,
    user_id bigint NOT NULL,
    title text  NOT NULL,
    type text NOT NULL,
    keywords text[] NOT NULL,
    contact_email text NOT NULL,
    intro text NOT NULL,
    description_addr text NOT NULL,
    description_file_addr text,
    duration smallint NOT NULL,
    expired_at timestamp with time zone NOT NULL,
    payments jsonb DEFAULT '[]'::jsonb NOT NULL,
    status text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE bounties_hunters_rel
(
    id bigint DEFAULT comunion.id_generator() NOT NULL
        CONSTRAINT bounties_hunters_rel_pk
            PRIMARY KEY,
    bounty_id bigint NOT NULL,
    hunter_id bigint NOT NULL,
    status text NOT NULL,
    started_at timestamp with time zone,
    submitted_at timestamp with time zone,
    quited_at timestamp with time zone,
    paid_at timestamp with time zone,
    paid_tokens jsonb DEFAULT '[]'::jsonb NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX bounties_hunters_rel_bounty_id_hunter_id ON bounties_hunters_rel(bounty_id,hunter_id);
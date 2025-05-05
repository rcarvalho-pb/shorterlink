CREATE TABLE IF NOT EXISTS tb_bloomfilter(
    id INTEGER,
    name VARCHAR(255) NOT NULL,
    size INTEGER NOT NULL,
    count INTEGER NOT NULL,
    data BLOB NOT NULL,
    CONSTRAINT bloomfilter_pk PRIMARY KEY(id),
    CONSTRAINT unique_name UNIQUE(name)
);

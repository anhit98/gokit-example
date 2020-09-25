# gokit-example
code to init todo table

```CREATE TABLE IF NOT EXISTS todo
(
    id character varying(254) NOT NULL,
    username character varying(254) NOT NULL,
    text text,
    created_at timestamp without time zone,
    CONSTRAINT todo_pkey PRIMARY KEY (id),
	 CONSTRAINT todo_id_key UNIQUE (id)
)```

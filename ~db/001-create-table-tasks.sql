-- Table: public.tasks

-- DROP TABLE IF EXISTS public.tasks;

CREATE TABLE IF NOT EXISTS public.tasks
(
    id serial NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    done boolean NOT NULL DEFAULT false,
    CONSTRAINT tasks_pkey PRIMARY KEY (id)
)

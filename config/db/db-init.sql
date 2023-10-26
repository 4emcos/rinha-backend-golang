
CREATE TABLE IF NOT EXISTS public.person
(
    nome text COLLATE pg_catalog."default" NOT NULL,
    apelido text COLLATE pg_catalog."default" NOT NULL,
    nascimento text COLLATE pg_catalog."default",
    id uuid NOT NULL,
    stack character varying[] COLLATE pg_catalog."default",
    CONSTRAINT person_pkey PRIMARY KEY (id)
);

create unique index person_nome_apelido_uindex on person (apelido)




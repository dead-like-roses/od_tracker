--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: activity; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.activity (
    device_id character(36),
    posted_at timestamp with time zone NOT NULL,
    app text
);


ALTER TABLE public.activity OWNER TO dbuser;

--
-- Name: devices; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.devices (
    device_id character(36) NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.devices OWNER TO dbuser;

--
-- Data for Name: activity; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.activity (device_id, posted_at, app) FROM stdin;
\.


--
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.devices (device_id, name) FROM stdin;
\.


--
-- Name: devices devices_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pkey PRIMARY KEY (device_id);


--
-- Name: activity activity_device_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_device_id_fkey FOREIGN KEY (device_id) REFERENCES public.devices(device_id);


--
-- PostgreSQL database dump complete
--


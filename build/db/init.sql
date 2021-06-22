--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2 (Debian 13.2-1.pgdg100+1)
-- Dumped by pg_dump version 13.2 (Debian 13.2-1.pgdg100+1)

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
-- Name: speedtest; Type: SCHEMA; Schema: -; Owner: speedtest_user
--

CREATE SCHEMA speedtest;


ALTER SCHEMA speedtest OWNER TO speedtest_user;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: results; Type: TABLE; Schema: speedtest; Owner: speedtest_user
--

CREATE TABLE speedtest.results (
    id bigint NOT NULL,
    latency bigint,
    download_speed numeric,
    upload_speed numeric,
    created_at timestamp with time zone
);


ALTER TABLE speedtest.results OWNER TO speedtest_user;

--
-- Name: results_id_seq; Type: SEQUENCE; Schema: speedtest; Owner: speedtest_user
--

CREATE SEQUENCE speedtest.results_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE speedtest.results_id_seq OWNER TO speedtest_user;

--
-- Name: results_id_seq; Type: SEQUENCE OWNED BY; Schema: speedtest; Owner: speedtest_user
--

ALTER SEQUENCE speedtest.results_id_seq OWNED BY speedtest.results.id;


--
-- Name: results id; Type: DEFAULT; Schema: speedtest; Owner: speedtest_user
--

ALTER TABLE ONLY speedtest.results ALTER COLUMN id SET DEFAULT nextval('speedtest.results_id_seq'::regclass);


--
-- Name: results results_pkey; Type: CONSTRAINT; Schema: speedtest; Owner: speedtest_user
--

ALTER TABLE ONLY speedtest.results
    ADD CONSTRAINT results_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

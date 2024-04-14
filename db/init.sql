--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

-- Started on 2024-04-14 11:37:58

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 25636)
-- Name: auth; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth (
    username character varying(64) NOT NULL,
    password character varying(64),
    isadmin boolean DEFAULT false
);


ALTER TABLE public.auth OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 25614)
-- Name: banners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banners (
    id bigint NOT NULL,
    feature_id bigint NOT NULL,
    tag bigint NOT NULL,
    title character varying(128),
    text character varying(512),
    url character varying(256),
    CONSTRAINT banners_feature_id_check CHECK ((feature_id > 0)),
    CONSTRAINT banners_tag_check CHECK ((tag > 0))
);


ALTER TABLE public.banners OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 25613)
-- Name: banners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banners_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.banners_id_seq OWNER TO postgres;

--
-- TOC entry 4858 (class 0 OID 0)
-- Dependencies: 215
-- Name: banners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banners_id_seq OWNED BY public.banners.id;


--
-- TOC entry 218 (class 1259 OID 25642)
-- Name: cache; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cache (
    feature_id bigint,
    tag bigint,
    CONSTRAINT cache_feature_id_check CHECK ((feature_id > 0)),
    CONSTRAINT cache_tag_check CHECK ((tag > 0))
);


ALTER TABLE public.cache OWNER TO postgres;

--
-- TOC entry 4696 (class 2604 OID 25617)
-- Name: banners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners ALTER COLUMN id SET DEFAULT nextval('public.banners_id_seq'::regclass);


--
-- TOC entry 4851 (class 0 OID 25636)
-- Dependencies: 217
-- Data for Name: auth; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth (username, password, isadmin) FROM stdin;
idkidk	idkidk	f
idkidk1	idkidk1	f
idkidkidk	idkidkidk	t
\.


--
-- TOC entry 4850 (class 0 OID 25614)
-- Dependencies: 216
-- Data for Name: banners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banners (id, feature_id, tag, title, text, url) FROM stdin;
6	267205879	11902357	Miller8308	paradigm	http://www.corporatecollaborative.io/world-class/cross-platform/enable/bricks-and-clicks
7	105474023	20949843	Streich9658	Innovative	https://www.legacysynergies.com/streamline/magnetic/brand/magnetic
9	12345	12345	First	this is the first sent message	go.dev
1	123456	123456	Second	this is the first message updated	go.dev/concurrecny
2	12345678	12345678	third	this is the third message updated	go.dev/garbage_collector
\.


--
-- TOC entry 4852 (class 0 OID 25642)
-- Dependencies: 218
-- Data for Name: cache; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cache (feature_id, tag) FROM stdin;
\.


--
-- TOC entry 4859 (class 0 OID 0)
-- Dependencies: 215
-- Name: banners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banners_id_seq', 9, true);


--
-- TOC entry 4705 (class 2606 OID 25641)
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (username);


--
-- TOC entry 4703 (class 2606 OID 25623)
-- Name: banners banners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_pkey PRIMARY KEY (feature_id, tag);


-- Completed on 2024-04-14 11:37:58

--
-- PostgreSQL database dump complete
--


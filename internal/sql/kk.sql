--
-- PostgreSQL database dump
--

-- Dumped from database version 14.12 (Ubuntu 14.12-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.12 (Ubuntu 14.12-1.pgdg20.04+1)

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
-- Name: cluster_dbs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cluster_dbs (
    cluster_id bigint NOT NULL,
    cluster_name character varying(255)
);


ALTER TABLE public.cluster_dbs OWNER TO root;

--
-- Name: COLUMN cluster_dbs.cluster_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_dbs.cluster_id IS '集群ID';


--
-- Name: COLUMN cluster_dbs.cluster_name; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_dbs.cluster_name IS '集群名称';


--
-- Name: cluster_dbs_cluster_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cluster_dbs_cluster_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cluster_dbs_cluster_id_seq OWNER TO root;

--
-- Name: cluster_dbs_cluster_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cluster_dbs_cluster_id_seq OWNED BY public.cluster_dbs.cluster_id;


--
-- Name: cluster_plugin_dbs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cluster_plugin_dbs (
    id bigint NOT NULL,
    cluster_id text NOT NULL,
    plugin_name character varying(40) NOT NULL,
    plugin_version character varying(10) NOT NULL,
    plugin_type character varying(30) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted boolean DEFAULT false,
    plugin_id bigint
);


ALTER TABLE public.cluster_plugin_dbs OWNER TO root;

--
-- Name: COLUMN cluster_plugin_dbs.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.id IS '集群使用的组件ID';


--
-- Name: COLUMN cluster_plugin_dbs.cluster_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.cluster_id IS '关联集群ID';


--
-- Name: COLUMN cluster_plugin_dbs.plugin_name; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.plugin_name IS '组件名称';


--
-- Name: COLUMN cluster_plugin_dbs.plugin_version; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.plugin_version IS '组件版本';


--
-- Name: COLUMN cluster_plugin_dbs.plugin_type; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.plugin_type IS '组件类型';


--
-- Name: COLUMN cluster_plugin_dbs.deleted; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.deleted IS '标记删除false|true';


--
-- Name: COLUMN cluster_plugin_dbs.plugin_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.cluster_plugin_dbs.plugin_id IS 'QKE支持的组件ID';


--
-- Name: cluster_plugin_dbs_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cluster_plugin_dbs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cluster_plugin_dbs_id_seq OWNER TO root;

--
-- Name: cluster_plugin_dbs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cluster_plugin_dbs_id_seq OWNED BY public.cluster_plugin_dbs.id;


--
-- Name: qke_plugin_dbs; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.qke_plugin_dbs (
    plugin_id bigint NOT NULL,
    plugin_name character varying(40) NOT NULL,
    plugin_version character varying(10) NOT NULL,
    plugin_type character varying(30) NOT NULL,
    install_type character varying(10),
    pre_versions text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted boolean DEFAULT false,
    k8s_versions text NOT NULL,
    description text,
    enable_installed boolean
);


ALTER TABLE public.qke_plugin_dbs OWNER TO root;

--
-- Name: COLUMN qke_plugin_dbs.plugin_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.plugin_id IS 'QKE支持的组件ID';


--
-- Name: COLUMN qke_plugin_dbs.plugin_name; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.plugin_name IS '组件名称';


--
-- Name: COLUMN qke_plugin_dbs.plugin_version; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.plugin_version IS '组件版本';


--
-- Name: COLUMN qke_plugin_dbs.plugin_type; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.plugin_type IS '组件类型';


--
-- Name: COLUMN qke_plugin_dbs.install_type; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.install_type IS '升级方式';


--
-- Name: COLUMN qke_plugin_dbs.pre_versions; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.pre_versions IS '可从低版本升级到当前版本，插入数据以逗号分割';


--
-- Name: COLUMN qke_plugin_dbs.deleted; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.deleted IS '标记删除false|true';


--
-- Name: COLUMN qke_plugin_dbs.k8s_versions; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.k8s_versions IS '支持的K8s版本';


--
-- Name: COLUMN qke_plugin_dbs.description; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.description IS '插件描述信息';


--
-- Name: COLUMN qke_plugin_dbs.enable_installed; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.qke_plugin_dbs.enable_installed IS '能否在集群创建后在选择安装, true支持, false不支持';


--
-- Name: qke_plugin_dbs_plugin_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.qke_plugin_dbs_plugin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qke_plugin_dbs_plugin_id_seq OWNER TO root;

--
-- Name: qke_plugin_dbs_plugin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.qke_plugin_dbs_plugin_id_seq OWNED BY public.qke_plugin_dbs.plugin_id;


--
-- Name: cluster_dbs cluster_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cluster_dbs ALTER COLUMN cluster_id SET DEFAULT nextval('public.cluster_dbs_cluster_id_seq'::regclass);


--
-- Name: cluster_plugin_dbs id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cluster_plugin_dbs ALTER COLUMN id SET DEFAULT nextval('public.cluster_plugin_dbs_id_seq'::regclass);


--
-- Name: qke_plugin_dbs plugin_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.qke_plugin_dbs ALTER COLUMN plugin_id SET DEFAULT nextval('public.qke_plugin_dbs_plugin_id_seq'::regclass);


--
-- Data for Name: cluster_dbs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cluster_dbs (cluster_id, cluster_name) FROM stdin;
\.


--
-- Data for Name: cluster_plugin_dbs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cluster_plugin_dbs (id, cluster_id, plugin_name, plugin_version, plugin_type, created_at, updated_at, deleted, plugin_id) FROM stdin;
\.


--
-- Data for Name: qke_plugin_dbs; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.qke_plugin_dbs (plugin_id, plugin_name, plugin_version, plugin_type, install_type, pre_versions, created_at, updated_at, deleted, k8s_versions, description, enable_installed) FROM stdin;
\.


--
-- Name: cluster_dbs_cluster_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cluster_dbs_cluster_id_seq', 1, false);


--
-- Name: cluster_plugin_dbs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cluster_plugin_dbs_id_seq', 1, false);


--
-- Name: qke_plugin_dbs_plugin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.qke_plugin_dbs_plugin_id_seq', 1, false);


--
-- Name: cluster_dbs cluster_dbs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cluster_dbs
    ADD CONSTRAINT cluster_dbs_pkey PRIMARY KEY (cluster_id);


--
-- Name: cluster_plugin_dbs cluster_plugin_dbs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cluster_plugin_dbs
    ADD CONSTRAINT cluster_plugin_dbs_pkey PRIMARY KEY (id);


--
-- Name: qke_plugin_dbs qke_plugin_dbs_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.qke_plugin_dbs
    ADD CONSTRAINT qke_plugin_dbs_pkey PRIMARY KEY (plugin_id);


--
-- PostgreSQL database dump complete
--


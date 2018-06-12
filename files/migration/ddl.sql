CREATE TABLE public.news (
  news_id serial NOT NULL,
  html_content text NOT NULL,
  status text NOT NULL DEFAULT 0,
  title text NOT NULL,
  is_valid int4 NOT NULL DEFAULT 1,
  CONSTRAINT news_pkey PRIMARY KEY (news_id)
)

CREATE INDEX index_status ON news USING btree (status) ;
CREATE INDEX index_title ON news USING btree (title) ;


CREATE TABLE public.topic (
  topic_id serial NOT NULL,
  "name" text NOT NULL,
  is_valid int4 NOT NULL DEFAULT 1,
  CONSTRAINT topic_pkey PRIMARY KEY (topic_id),
  CONSTRAINT unique_topic_name UNIQUE (name)
)

CREATE TABLE public.topic_news (
  topic_id int4 NOT NULL,
  news_id int4 NOT NULL,
  is_valid int4 NOT NULL DEFAULT 1,
  CONSTRAINT topic_news_pkey PRIMARY KEY (topic_id, news_id),
  CONSTRAINT topic_news_news_id_fkey FOREIGN KEY (news_id) REFERENCES news(news_id) ON UPDATE CASCADE,
  CONSTRAINT topic_news_topic_id_fkey FOREIGN KEY (topic_id) REFERENCES topic(topic_id) ON UPDATE CASCADE ON DELETE CASCADE
)
CREATE TABLE blog_posts (
    id serial PRIMARY KEY,
    slug varchar(96) UNIQUE NOT NULL,
    title text NOT NULL,
    body bytea NOT NULL,
    created timestamp(0) with time zone NOT NULL,
    updated timestamp(0) with time zone NOT NULL,
    is_code_snippet bool NOT NULL DEFAULT false
);

CREATE TABLE blog_tags (
    id serial PRIMARY KEY,
    name varchar(24) UNIQUE NOT NULL
);

CREATE TABLE post_tags (
  post_id int NOT NULL REFERENCES blog_posts(id) ON DELETE CASCADE,
  tag_id int NOT NULL REFERENCES blog_tags(id) ON DELETE CASCADE,
  PRIMARY KEY (post_id, tag_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS post_tags;
DROP TABLE IF EXISTS blog_posts;
DROP TABLE IF EXISTS blog_tags;

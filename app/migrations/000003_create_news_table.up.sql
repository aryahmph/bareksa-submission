CREATE TABLE IF NOT EXISTS news
(
    id           INT UNSIGNED                         NOT NULL PRIMARY KEY AUTO_INCREMENT,
    topic_name   VARCHAR(255)                         NOT NULL,
    title        VARCHAR(255)                         NOT NULL,
    short_desc   VARCHAR(255)                         NOT NULL,
    content      TEXT                                 NOT NULL,
    image_url    VARCHAR(255)                         NOT NULL UNIQUE,
    writer       VARCHAR(255)                         NOT NULL,
    status       ENUM ('draft', 'deleted', 'publish') NOT NULL,
    published_at DATE,
    created_at   TIMESTAMP                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP,
    deleted_at   TIMESTAMP,
    CONSTRAINT fk_news_topic
        FOREIGN KEY (topic_name) REFERENCES topics (name)
);
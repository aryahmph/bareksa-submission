CREATE TABLE IF NOT EXISTS news_tags
(
    id         INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_news    INT UNSIGNED NOT NULL,
    name_tag   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_news_tag_news
        FOREIGN KEY (id_news) REFERENCES news (id),
    CONSTRAINT fk_news_tag_tag
        FOREIGN KEY (name_tag) REFERENCES tags (name)
            ON DELETE CASCADE ON UPDATE CASCADE
);
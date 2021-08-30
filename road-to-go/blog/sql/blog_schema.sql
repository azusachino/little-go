CREATE DATABASE IF NOT EXISTS `blog_service` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8bm4_general_ci;

-- blog_tag
CREATE TABLE IF NOT EXISTS `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT 'tag name',
    `state`       tinyint(3) unsigned DEFAULT 1 COMMENT '0 forbidden, 1 open',
    `created_on`  int(10) unsigned    DEFAULT 0 COMMENT 'create time',
    `created_by`  varchar(100)        DEFAULT '' COMMENT 'created by someone',
    `modified_on` int(10) unsigned    DEFAULT 0 COMMENT 'modified time',
    `modified_by` varchar(100)        DEFAULT '' COMMENT 'modified by someone',
    `deleted_on`  int(10) unsigned    DEFAULT 0 COMMENT 'delete time',
    `is_del`      tinyint(3) unsigned DEFAULT 0 COMMENT '0 not delete, 1 deleted',
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8bm4 COMMENT 'tag management';

-- blog_article
CREATE TABLE IF NOT EXISTS `blog_article`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`           varchar(100)        DEFAULT '' COMMENT 'article title',
    `desc`            varchar(255)        DEFAULT '' COMMENT 'article desc',
    `cover_image_url` varchar(255)        DEFAULT '' COMMENT 'article cover image url',
    `content`         text                DEFAULT '' COMMENT 'article content',
    `created_on`      int(10) unsigned    DEFAULT 0 COMMENT 'create time',
    `created_by`      varchar(100)        DEFAULT '' COMMENT 'created by someone',
    `modified_on`     int(10) unsigned    DEFAULT 0 COMMENT 'modified time',
    `modified_by`     varchar(100)        DEFAULT '' COMMENT 'modified by someone',
    `deleted_on`      int(10) unsigned    DEFAULT 0 COMMENT 'delete time',
    `is_del`          tinyint(3) unsigned DEFAULT 0 COMMENT '0 not delete, 1 deleted',
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8bm4 COMMENT 'article management';

-- blog_article_tag
CREATE TABLE IF NOT EXISTS `blog_article_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`  int(10)          NOT NULL COMMENT 'tag name',
    `tag_id`      int(10)          NOT NULL COMMENT '0 forbidden, 1 open',
    `created_on`  int(10) unsigned    DEFAULT 0 COMMENT 'create time',
    `created_by`  varchar(100)        DEFAULT '' COMMENT 'created by someone',
    `modified_on` int(10) unsigned    DEFAULT 0 COMMENT 'modified time',
    `modified_by` varchar(100)        DEFAULT '' COMMENT 'modified by someone',
    `deleted_on`  int(10) unsigned    DEFAULT 0 COMMENT 'delete time',
    `is_del`      tinyint(3) unsigned DEFAULT 0 COMMENT '0 not delete, 1 deleted',
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8bm4 COMMENT 'article tag relation';


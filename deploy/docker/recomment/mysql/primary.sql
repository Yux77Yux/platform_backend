CREATE DATABASE IF NOT EXISTS db_comment_areas_1;
CREATE DATABASE IF NOT EXISTS db_comments_1;
CREATE DATABASE IF NOT EXISTS db_comment_content_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_comment_areas_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_comments_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_comment_content_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;


-- 使用 db_comment_areas_1 数据库
USE db_comment_areas_1;

CREATE TABLE IF NOT EXISTS Creation (
    creation_id BIGINT,                -- 作品ID，使用 BIGINT
    creation_author_id BIGINT NOT NULL,    -- 作者ID
    creation_arc VARCHAR(255),                   -- 用户头像，可以存储头像的 URL 或文件路径
    creation_title VARCHAR(255),                -- 作品标题，最大 255 个字符
    creation_bio VARCHAR(1000),     -- 作品简介，最大 1000 个字符
    creation_status ENUM('DRAFT', 'PENDING', 'PUBLISHED', 'REJECTED') NOT NULL DEFAULT 'DRAFT', -- 作品状态，草稿、审核中、已发布、已拒绝
    creation_duration_time INT DEFAULT 0,        -- 视频时长，单位秒
    creation_upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 默认当前时间

    PRIMARY KEY (creation_id),                   -- 使用 creation_id 作为主键
    INDEX idx_author (creation_author_id),       -- 按作者 ID 索引
    INDEX idx_title (creation_title),            -- 按标题索引
    INDEX idx_upload (creation_upload_time),            -- 按标题索引
    INDEX idx_creation_status (creation_status)  -- 按状态索引
);

-- 使用 db_comments_1 数据库
USE db_comments_1;

CREATE TABLE IF NOT EXISTS CreationEngagement (
    creation_id BIGINT NOT NULL,                 -- 作品ID，与 Creation 表的 creation_id 一对一
    creation_views INT DEFAULT 0,                 -- 播放数
    creation_likes INT DEFAULT 0,                 -- 点赞数
    creation_saves INT DEFAULT 0,                 -- 收藏数
    creation_publish_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作品发布时间

    PRIMARY KEY (creation_id),                   -- 使用 creation_id 作为主键
    INDEX idx_views (creation_views),    -- 按播放数索引
    INDEX idx_likes (creation_likes),    -- 按点赞数索引
    INDEX idx_saves (creation_saves),    -- 按收藏数索引
    INDEX idx_publish_time (creation_publish_time)    -- 按发布时间索引
);

-- 使用 db_comment_content_1 数据库
USE db_comment_content_1;

CREATE TABLE IF NOT EXISTS Category (
    category_id INT NOT NULL AUTO_INCREMENT,    -- 分类ID
    parent INT NOT NULL DEFAULT 0,              -- 父分类ID，0 表示大分区，非0表示二级分区
    name VARCHAR(255) NOT NULL,                 -- 分类名称
    description VARCHAR(1000),                  -- 分类描述

    PRIMARY KEY (category_id),                  -- 主键
    INDEX idx_parent (parent),                  -- 根据 parent 字段索引
    INDEX idx_name (name)                       -- 根据 name 字段索引
);
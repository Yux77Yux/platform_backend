CREATE DATABASE IF NOT EXISTS db_creation_1;
CREATE DATABASE IF NOT EXISTS db_creation_engagment_1;
CREATE DATABASE IF NOT EXISTS db_creation_category_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_engagment_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_category_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;


-- 使用 db_1 数据库
USE db_creation_1;

CREATE TABLE IF NOT EXISTS Creation (
    id BIGINT,                -- 作品ID，使用 BIGINT
    author_id BIGINT NOT NULL,    -- 作者ID
    arc TEXT,                   -- 用户头像，可以存储头像的 URL 或文件路径
    thumbnail TEXT NOT NULL,                   -- 封面
    title VARCHAR(255),                -- 作品标题，最大 255 个字符
    bio VARCHAR(1000),     -- 作品简介，最大 1000 个字符
    status ENUM('DRAFT', 'PENDING', 'PUBLISHED', 'REJECTED') NOT NULL DEFAULT 'DRAFT', -- 作品状态，草稿、审核中、已发布、已拒绝
    duration INT DEFAULT 0,        -- 视频时长，单位秒
    upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 默认当前时间

    PRIMARY KEY (id),                   -- 使用 id 作为主键
    INDEX idx_author (author_id),       -- 按作者 ID 索引 ，查询作者的作品
    FULLTEXT INDEX idx_title (title)            -- 按标题索引
);

-- 使用 db_engagment_1 数据库
USE db_creation_engagment_1;

CREATE TABLE IF NOT EXISTS CreationEngagement (
    creation_id BIGINT NOT NULL,                 -- 作品ID，与 Creation 表的 id 一对一
    views INT DEFAULT 0,                 -- 播放数
    likes INT DEFAULT 0,                 -- 点赞数
    saves INT DEFAULT 0,                 -- 收藏数
    publish_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作品发布时间

    PRIMARY KEY (creation_id),                   -- 使用 id 作为主键
    INDEX idx_saves (saves),    -- 按收藏数索引
    INDEX idx_publish_time (publish_time)    -- 按发布时间索引
);

-- 使用 db_category_1 数据库
USE db_creation_category_1;

CREATE TABLE IF NOT EXISTS Category (
    id INT NOT NULL AUTO_INCREMENT,    -- 分区ID
    creation_id INT NOT NULL,    -- 作品ID
    parent INT NOT NULL DEFAULT 0,              -- 父分区ID，0 表示大分区，非0表示二级分区
    name VARCHAR(255) NOT NULL,                 -- 分类名称
    description VARCHAR(1000),                  -- 分类描述

    PRIMARY KEY (id),                  -- 主键
    INDEX idx_parent (parent),                  -- 根据 parent 字段索引
    INDEX idx_name (name)                       -- 根据 name 字段索引
);
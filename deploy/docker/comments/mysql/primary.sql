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

CREATE TABLE IF NOT EXISTS CommentAreas (
    creation_id BIGINT PRIMARY KEY,
    total_comments INT DEFAULT 0,
    areas_status ENUM('ACTIVE','INACTIVE','HIDE') DEFAULT 'ACTIVE'
);

-- 使用 db_comments_1 数据库
USE db_comments_1;

CREATE TABLE IF NOT EXISTS Comments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,      -- 评论ID
    root BIGINT DEFAULT 0,                     -- 一级评论ID
    parent BIGINT DEFAULT 0,                   -- 回复对象所在ID
    dialog BIGINT DEFAULT 0,                   -- 二级评论ID
    user_id BIGINT NOT NULL,                   -- 发言的用户ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
    creation_id BIGINT NOT NULL,                -- 作品ID

    INDEX idx_root_parent_dialog(creation_id, root, parent)
);

-- 使用 db_comment_content_1 数据库
USE db_comment_content_1;

CREATE TABLE IF NOT EXISTS CommentContent (
    comment_id BIGINT PRIMARY KEY,                      -- 评论ID，外键关联评论表（假设 `comment_id` 是评论表的主键）
    content TEXT,                            -- 评论内容，TEXT类型
    media VARCHAR(255),                      -- 评论附件媒体文件，URL或文件路径（如果有）
);
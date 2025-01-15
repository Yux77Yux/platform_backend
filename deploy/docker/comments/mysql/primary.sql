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
    creation_id INT PRIMARY KEY,
    total_comments INT DEFAULT 0,
    areas_status ENUM('ACTIVE','INACTIVE','HIDE') DEFAULT 'ACTIVE'
);

-- 使用 db_comments_1 数据库
USE db_comments_1;

-- 评论表
CREATE TABLE IF NOT EXISTS Comments (
    id INT AUTO_INCREMENT PRIMARY KEY,            -- 评论ID
    root INT DEFAULT 0,                           -- 一级评论ID
    parent INT DEFAULT 0,                         -- 回复对象所在ID
    dialog INT DEFAULT 0,                         -- 二级评论ID
    user_id BIGINT NOT NULL,                         -- 发言的用户ID
    creation_id BIGINT NOT NULL,                     -- 作品ID
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
    status ENUM('PUBLISHED','DELETE') NOT NULL DEFAULT 'PUBLISHED',  -- 公告状态，默认是草稿

    INDEX idx_creation_root(creation_id, root),  -- 评论索引
    INDEX idx_status(status, created_at),        -- 主要用于清除DELETE状态行
    INDEX idx_user(user_id, created_at),  -- 与下配合
    INDEX idx_parent(parent, created_at)  -- 回复索引
);

-- 评论内容表
CREATE TABLE IF NOT EXISTS CommentContent (
    comment_id INT PRIMARY KEY,                   -- 评论ID，外键关联评论表
    content TEXT,                                    -- 评论内容，TEXT类型
    media TEXT,                                      -- 评论附件媒体文件，URL或文件路径（如果有）
    
    CONSTRAINT fk_comment FOREIGN KEY (comment_id) REFERENCES Comments(id)  -- 外键约束
);
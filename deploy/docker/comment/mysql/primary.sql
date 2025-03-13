CREATE DATABASE IF NOT EXISTS db_comment_area_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS db_comment_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_comment_area_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_comment_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

-- 使用 db_comment_area_1 数据库
USE db_comment_area_1;

CREATE TABLE IF NOT EXISTS CommentArea (
    creation_id BIGINT PRIMARY KEY,
    total_comments INT DEFAULT 0,
    areas_status ENUM('DEFAULT','HIDING','CLOSED') DEFAULT 'DEFAULT'
);

-- 使用 db_comment_1 数据库
USE db_comment_1;

-- 评论表
CREATE TABLE IF NOT EXISTS Comment (
    id INT AUTO_INCREMENT PRIMARY KEY,            -- 评论ID
    root INT DEFAULT 0,                           -- 一级评论ID
    parent INT DEFAULT 0,                         -- 回复对象所在ID
    dialog INT DEFAULT 0,                         -- 二级评论ID
    user_id BIGINT NOT NULL,                         -- 发言的用户ID
    creation_id BIGINT NOT NULL,                     -- 作品ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
    status ENUM('PUBLISHED','DELETED') NOT NULL DEFAULT 'PUBLISHED',  -- 评论状态

    INDEX idx_creation_root(creation_id, root),  -- 评论索引
    INDEX idx_user(user_id, created_at),  -- 与下配合  比如返回我说的，然后拿我说的id去请求回复id是我的其他id
    INDEX idx_parent(parent, created_at)  -- 回复索引
);

-- 评论内容表
CREATE TABLE IF NOT EXISTS CommentContent (
    comment_id INT PRIMARY KEY,                   -- 评论ID，外键关联评论表
    content TEXT,                                    -- 评论内容，TEXT类型
    media TEXT                                      -- 评论附件媒体文件，URL或文件路径（如果有）
);
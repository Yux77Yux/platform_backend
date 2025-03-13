CREATE DATABASE IF NOT EXISTS db_announcement_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_announcement_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

-- 使用 db_announcement_1 数据库
USE db_announcement_1;

CREATE TABLE IF NOT EXISTS Announcement (
    id INT AUTO_INCREMENT PRIMARY KEY,   -- 使用 AUTO_INCREMENT 来自动生成 id
    title NVARCHAR(255) NOT NULL,          -- 公告标题
    content TEXT NOT NULL,                -- 使用 TEXT 类型来存储公告内容
    publisher_id BIGINT NOT NULL,         -- 发布者 ID
    status ENUM('DRAFT', 'PUBLISHED','DELETE') NOT NULL DEFAULT 'DRAFT',  -- 公告状态，默认是草稿
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 公告创建时间，默认当前时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 公告更新时间，自动更新
);

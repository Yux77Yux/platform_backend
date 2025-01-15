CREATE DATABASE IF NOT EXISTS db_user_credentials_1;
CREATE DATABASE IF NOT EXISTS db_user_1;
CREATE DATABASE IF NOT EXISTS db_user_follow_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_user_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_user_credentials_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_user_follow_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

USE db_user_credentials_1;

CREATE TABLE IF NOT EXISTS UserCredentials (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,  -- 存储加密后的密码
    email VARCHAR(255) NULL,     -- 存储用户电子邮件
    user_id BIGINT NOT NULL,     -- id
    role ENUM('USER', 'ADMIN', 'SUPER_ADMIN') NOT NULL DEFAULT 'USER',-- 用户身份，使用枚举值，默认值为 USER (0)
    PRIMARY KEY (username),          -- 用户名作为主键
    UNIQUE (email),                   -- 电子邮件地址必须唯一
    UNIQUE (user_id)               -- id必须唯一
);

USE db_user_1;

CREATE TABLE IF NOT EXISTS User (
    id BIGINT NOT NULL,                -- id 长度固定为 36 个字符
    name VARCHAR(100) NOT NULL,            -- 用户名，最大 100 个字符
    avatar TEXT,                   -- 用户头像，可以存储头像的 URL 或文件路径
    bio VARCHAR(1000),                     -- 用户简介，最大 1000 个字符
    status ENUM('INACTIVE', 'ACTIVE', 'HIDING', 'LIMITED', 'DELETE') NOT NULL DEFAULT 'INACTIVE',         -- 用户状态，使用枚举值，默认值为 INACTIVE（1）
    gender ENUM('UNDEFINED', 'MALE', 'FEMALE') NOT NULL DEFAULT 'UNDEFINED',         -- 用户性别，使用枚举值，默认值为 UNDEFINED（0）
    bday DATE,                             -- 用户生日，DATE 类型
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间，默认当前时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间，自动更新时间
    PRIMARY KEY (id),                     -- 使用 id 作为主键
    INDEX idx_name (name)
);

USE db_user_follow_1;

CREATE TABLE IF NOT EXISTS Follow (
    follower_id BIGINT,                             -- 粉丝follower_id
    followee_id BIGINT,                             -- 作者followee_id
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间，默认当前时间 
    views INT DEFAULT 0,                            -- 访问次数
    
    PRIMARY KEY (follower_id,followee_id),          -- follower_id作为主键
    INDEX idx_views (views),
    INDEX idx_created (created_at)
);
CREATE DATABASE IF NOT EXISTS db_user_1;
CREATE DATABASE IF NOT EXISTS db_user_credentials_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_user_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_user_credentials_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

USE db_user_credentials_1;

CREATE TABLE IF NOT EXISTS UserCredentials (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,  -- 存储加密后的密码
    user_email VARCHAR(255) NULL,     -- 存储用户电子邮件
    user_id BIGINT NOT NULL,     -- user_id
    user_role ENUM('USER', 'ADMIN', 'SUPER_ADMIN') NOT NULL DEFAULT 'USER',-- 用户身份，使用枚举值，默认值为 USER (0)
    PRIMARY KEY (username),          -- 用户名作为主键
    UNIQUE (user_email),                   -- 电子邮件地址必须唯一
    UNIQUE (user_id),               -- user_id必须唯一
    INDEX idx_role (user_role)
);

USE db_user_1;

CREATE TABLE IF NOT EXISTS User (
    user_id BIGINT NOT NULL,                -- user_id 长度固定为 36 个字符
    user_name VARCHAR(100) NOT NULL,            -- 用户名，最大 100 个字符
    user_avator VARCHAR(255),                   -- 用户头像，可以存储头像的 URL 或文件路径
    user_bio VARCHAR(1000),                     -- 用户简介，最大 1000 个字符
    user_status ENUM('HIDING', 'INACTIVE', 'ACTIVE', 'LIMITED') NOT NULL DEFAULT 'INACTIVE',         -- 用户状态，使用枚举值，默认值为 INACTIVE（1）
    user_gender ENUM('UNDEFINED', 'MALE', 'FEMALE') NOT NULL DEFAULT 'UNDEFINED',         -- 用户性别，使用枚举值，默认值为 UNDEFINED（0）
    user_email VARCHAR(255) NULL,     -- 存储用户电子邮件
    user_bday DATE,                             -- 用户生日，DATE 类型
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间，默认当前时间
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间，自动更新时间
    user_role ENUM('USER', 'ADMIN', 'SUPER_ADMIN') NOT NULL DEFAULT 'USER',-- 用户身份，使用枚举值，默认值为 USER (0)
    PRIMARY KEY (user_id),                     -- 使用 user_id 作为主键
    UNIQUE (user_email),
    INDEX idx_role (user_role)
);
ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

USE db_user_1;

CREATE TABLE IF NOT EXISTS UserCredentials (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,  -- 存储加密后的密码
    email VARCHAR(255) NOT NULL,     -- 存储用户电子邮件
    PRIMARY KEY (username),          -- 用户名作为主键
    UNIQUE (email)                   -- 电子邮件地址必须唯一
);

CREATE TABLE IF NOT EXISTS User (
    user_uuid CHAR(36) NOT NULL,                -- UUID 长度固定为 36 个字符
    user_name VARCHAR(100) NOT NULL,            -- 用户名，最大 100 个字符
    user_avator VARCHAR(255),                   -- 用户头像，可以存储头像的 URL 或文件路径
    user_bio VARCHAR(1000),                     -- 用户简介，最大 1000 个字符
    user_status INT NOT NULL DEFAULT 1,         -- 用户状态，使用枚举值，默认值为 INACTIVE（1）
    user_gender INT NOT NULL DEFAULT 0,         -- 用户性别，使用枚举值，默认值为 UNDEFINED（0）
    user_bday DATE,                             -- 用户生日，DATE 类型
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间，默认当前时间
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间，自动更新时间
    PRIMARY KEY (user_uuid)                     -- 使用 UUID 作为主键
);
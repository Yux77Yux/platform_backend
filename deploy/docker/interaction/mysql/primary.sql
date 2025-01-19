CREATE DATABASE IF NOT EXISTS db_interaction_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_interaction_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;


-- 使用 db_interaction_1 数据库
USE db_interaction_1;

CREATE TABLE IF NOT EXISTS Interaction (
    user_id BIGINT,    -- 用户ID
    creation_id BIGINT,                -- 作品ID
    action_tag TINYINT DEFAULT 1,                   -- 动作记录，100表示收藏，010表示点赞，001表示观看
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 默认当前时间
    save_at TIMESTAMP DEFAULT NULL,                           -- 收藏的时间

    PRIMARY KEY (user_id,creation_id),                 
    INDEX idx_user_action_updated (user_id,action_tag,updated_at),       -- 历史记录
    INDEX idx_user_action_save (user_id,action_tag,save_at)       -- 收藏夹
);

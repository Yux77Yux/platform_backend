CREATE DATABASE IF NOT EXISTS db_review_1;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_review_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;


-- 使用 db_review_1 数据库
USE db_review_1;

CREATE TABLE IF NOT EXISTS Review (
    id BIGINT,                -- 审核信息ID，使用 BIGINT
    target_id BIGINT NOT NULL,    -- 审核的ID
    target_type ENUM('USER','COMMENT','CREATION') NOT NULL,          -- 审核类型
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    status ENUM('PENDING', 'APPROVED', 'REJECTED','DELETE') NOT NULL DEFAULT 'PENDING', -- 审核状态，
    remark varchar(255),       -- 审核备注，最大255字符
    reviewer_id BIGINT, -- 审核人ID

    PRIMARY KEY (id),                   -- 使用 id 作为主键
    INDEX idx_reviewer (reviewer_id,status),       -- 按作者 ID 索引
    INDEX idx_target (target_id, target_type)  -- 针对 target_id 和 target_type 的复合索引
);
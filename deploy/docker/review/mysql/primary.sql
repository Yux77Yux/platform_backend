CREATE DATABASE IF NOT EXISTS db_review_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_review_1.* TO 'yuxyuxx'@'%';

FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;

-- 使用 db_review_1 数据库
USE db_review_1;

CREATE TABLE IF NOT EXISTS Review (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,                -- 审核信息ID，使用 BIGINT
    target_id BIGINT NOT NULL,    -- 审核对象的ID
    target_type ENUM('USER','COMMENT','CREATION') NOT NULL,          -- 审核类型
    detail TEXT,                                     -- 举报理由
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    status ENUM('PENDING', 'APPROVED', 'REJECTED','DELETED') NOT NULL DEFAULT 'PENDING', -- 审核状态，
    remark TEXT,       -- 审核备注，最大255字符
    reviewer_id BIGINT, -- 审核人ID

    INDEX idx_reviewer (reviewer_id,target_type,status,created_at, id),       -- 按审核人 ID 索引
    INDEX idx_updated_at (updated_at DESC, id)       -- 按审核人 ID 索引
);
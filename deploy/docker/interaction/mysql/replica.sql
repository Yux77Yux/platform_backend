-- 在从数据库中更改主服务器配置
CHANGE MASTER TO 
    MASTER_HOST='mysql-primary-service', 
    MASTER_USER='yuxyuxx', 
    MASTER_PASSWORD='yuxyuxx',
    MASTER_LOG_FILE='primary-bin.000001', 
    MASTER_LOG_POS=4;

-- 启动复制进程
START REPLICA;

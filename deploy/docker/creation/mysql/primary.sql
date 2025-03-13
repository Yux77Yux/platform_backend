CREATE DATABASE IF NOT EXISTS db_creation_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS db_creation_engagment_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS db_creation_category_1 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';

GRANT ALL PRIVILEGES ON db_creation_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_creation_engagment_1.* TO 'yuxyuxx'@'%';
GRANT ALL PRIVILEGES ON db_creation_category_1.* TO 'yuxyuxx'@'%';


FLUSH PRIVILEGES;

ALTER USER 'yuxyuxx'@'%' IDENTIFIED WITH mysql_native_password BY 'yuxyuxx';
GRANT REPLICATION SLAVE ON *.* TO 'yuxyuxx'@'%';
FLUSH PRIVILEGES;


-- 使用 db_1 数据库
USE db_creation_1;

CREATE TABLE IF NOT EXISTS Creation (
    id BIGINT,                -- 作品ID，使用 BIGINT
    author_id BIGINT NOT NULL,    -- 作者ID
    src TEXT,                   -- 用户头像，可以存储头像的 URL 或文件路径
    thumbnail TEXT NOT NULL,                   -- 封面
    title VARCHAR(255) COLLATE utf8mb4_unicode_ci,              
    bio TEXT COLLATE utf8mb4_unicode_ci,         -- 作品简介
    status ENUM('DRAFT', 'PENDING', 'PUBLISHED', 'REJECTED','DELETE') NOT NULL DEFAULT 'DRAFT', -- 作品状态，草稿、审核中、已发布、已拒绝
    duration INT DEFAULT 0,        -- 视频时长，单位秒
    category_id INT NOT NULL,    -- 作品ID
    upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 默认当前时间

    PRIMARY KEY (id),                   -- 使用 id 作为主键
    INDEX idx_author (author_id,status)       -- 按作者 ID 索引 ，查询作者的作品
);

-- 使用 db_engagment_1 数据库
USE db_creation_engagment_1;

CREATE TABLE IF NOT EXISTS CreationEngagement (
    creation_id BIGINT NOT NULL,                 -- 作品ID，与 Creation 表的 id 一对一
    views INT DEFAULT 0,                 -- 播放数
    likes INT DEFAULT 0,                 -- 点赞数
    saves INT DEFAULT 0,                 -- 收藏数
    publish_time TIMESTAMP NULL, -- 作品发布时间

    PRIMARY KEY (creation_id),                   -- 使用 id 作为主键
    INDEX idx_creation_views (views),    -- 按 creation_id 和 views 联合索引，最多播放
    INDEX idx_creation_saves (saves),    -- 按 creation_id 和 saves 联合索引，最多收藏
    INDEX idx_creation_publish_time (publish_time) -- 按 creation_id 和 publish_time 联合索引，最新发布
);

-- 使用 db_category_1 数据库
USE db_creation_category_1;

CREATE TABLE IF NOT EXISTS Category (
    id INT,    -- 分区ID
    parent INT NOT NULL DEFAULT 0,              -- 父分区ID，0 表示大分区，非0表示二级分区
    name VARCHAR(30) NOT NULL,                 -- 分类名称
    description VARCHAR(80),                  -- 分类描述

    PRIMARY KEY (id),                  -- 主键
    INDEX idx_parent (parent)                  -- 根据 parent 字段索引
);

-- 动画分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(1, 0, 'Animation', '动画相关的内容，包括经典动画和现代动画'),
(2, 1, 'MAD·AMV', '音乐动画制作与混剪'),
(3, 1, 'MMD·3D', 'MikuMikuDance和3D动画'),
(4, 1, 'Fanworks', '同人作品及手书内容'),
(5, 1, 'Dubbing', '动画角色或创意内容的配音'),
(6, 1, 'AnimeTalks', '关于动漫的讨论和评论');

-- 游戏分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(7, 0, 'Gaming', '关于电子游戏的实况、评测和讨论'),
(8, 7, 'SinglePlayer', '单机游戏内容和玩法'),
(9, 7, 'Esports', '电竞比赛和选手相关内容'),
(10, 7, 'MobileGames', '移动端游戏相关内容'),
(11, 7, 'OnlineGames', '多人在线网络游戏'),
(12, 7, 'BoardGames', '桌面游戏和棋牌内容'),
(13, 7, 'MusicGames', '音乐游戏相关内容');

-- 音乐分区及其二级分区（新增音乐综合）
INSERT INTO Category (id, parent, name, description) VALUES
(14, 0, 'Music', '与音乐相关的内容，包括表演和教程'),
(15, 14, 'OriginalMusic', '原创音乐创作'),
(16, 14, 'Covers', '翻唱歌曲分享'),
(17, 14, 'Performances', '乐器演奏相关内容'),
(18, 14, 'MusicReviews', '乐评盘点和音乐内容分析'),
(19, 14, 'MusicGeneral', '综合音乐内容，含多样化主题');

-- 影视分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(20, 0, 'Movies', '与影视相关的内容，包括影评和电影解说'),
(21, 20, 'MovieTalks', '对影视剧的评论和分析'),
(22, 20, 'MovieEdits', '影视片段的剪辑和创意编辑');

-- 知识分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(23, 0, 'Knowledge', '关于各种主题的教育和知识内容'),
(24, 23, 'Science', '科学知识普及'),
(25, 23, 'SocialScience', '社会科学、法律、心理学内容'),
(26, 23, 'History', '历史、人文学科相关内容'),
(27, 23, 'Finance', '财经和商业相关知识');

-- 科技分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(28, 0, 'Technology', '与科技相关的话题，包括科技产品和创新'),
(29, 28, 'DigitalProducts', '数码产品评测与资讯'),
(30, 28, 'ComputerTech', '计算机技术与知识分享'),
(31, 28, 'GeekDIY', '极客精神下的创意DIY');

-- 美食分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(32, 0, 'Food', '美食的制作、食谱和烹饪技巧'),
(33, 32, 'FoodMaking', '美食的制作方法和教程'),
(34, 32, 'FoodDetective', '美食探索与发现'),
(35, 32, 'FoodReviews', '美食的测评和体验'),
(36, 32, 'FoodDiary', '记录日常美食和生活');

-- 动物圈分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(37, 0, 'Animals', '关于动物的相关内容'),
(38, 37, 'Cats', '关于猫的内容'),
(39, 37, 'Dogs', '关于狗的内容'),
(40, 37, 'ExoticPets', '关于奇特宠物的内容'),
(41, 37, 'WildAnimals', '关于野生动物的内容');

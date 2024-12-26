CREATE DATABASE IF NOT EXISTS db_creation_1;
CREATE DATABASE IF NOT EXISTS db_creation_engagment_1;
CREATE DATABASE IF NOT EXISTS db_creation_category_1;

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
    title VARCHAR(255),                -- 作品标题，最大 255 个字符
    bio VARCHAR(1000),     -- 作品简介，最大 1000 个字符
    status ENUM('DRAFT', 'PENDING', 'PUBLISHED', 'REJECTED') NOT NULL DEFAULT 'DRAFT', -- 作品状态，草稿、审核中、已发布、已拒绝
    duration INT DEFAULT 0,        -- 视频时长，单位秒
    category_id INT NOT NULL,    -- 作品ID
    upload_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 默认当前时间

    PRIMARY KEY (id),                   -- 使用 id 作为主键
    INDEX idx_author (author_id),       -- 按作者 ID 索引 ，查询作者的作品
    FULLTEXT INDEX idx_title (title)            -- 按标题索引
);

-- 使用 db_engagment_1 数据库
USE db_creation_engagment_1;

CREATE TABLE IF NOT EXISTS CreationEngagement (
    creation_id BIGINT NOT NULL,                 -- 作品ID，与 Creation 表的 id 一对一
    views INT DEFAULT 0,                 -- 播放数
    likes INT DEFAULT 0,                 -- 点赞数
    saves INT DEFAULT 0,                 -- 收藏数
    publish_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作品发布时间

    PRIMARY KEY (creation_id),                   -- 使用 id 作为主键
    INDEX idx_saves (saves),    -- 按收藏数索引
    INDEX idx_publish_time (publish_time)    -- 按发布时间索引
);

-- 使用 db_category_1 数据库
USE db_creation_category_1;

CREATE TABLE IF NOT EXISTS Category (
    id INT,    -- 分区ID
    parent INT NOT NULL DEFAULT 0,              -- 父分区ID，0 表示大分区，非0表示二级分区
    name VARCHAR(20) NOT NULL,                 -- 分类名称
    description VARCHAR(200),                  -- 分类描述

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

-- 音乐分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(14, 0, 'Music', '与音乐相关的内容，包括表演和教程'),
(15, 14, 'OriginalMusic', '原创音乐创作'),
(16, 14, 'Covers', '翻唱歌曲分享'),
(17, 14, 'Performances', '乐器演奏相关内容');

-- 影视分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(18, 0, 'Movies', '与影视相关的内容，包括影评和电影解说'),
(19, 18, 'MovieTalks', '对影视剧的评论和分析'),
(20, 18, 'MovieEdits', '影视片段的剪辑和创意编辑');

-- 知识分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(21, 0, 'Knowledge', '关于各种主题的教育和知识内容'),
(22, 21, 'Science', '科学知识普及'),
(23, 21, 'SocialScience', '社会科学、法律、心理学内容'),
(24, 21, 'History', '历史、人文学科相关内容'),
(25, 21, 'Finance', '财经和商业相关知识');

-- 科技分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(26, 0, 'Technology', '与科技相关的话题，包括科技产品和创新'),
(27, 26, 'DigitalProducts', '数码产品评测与资讯'),
(28, 26, 'ComputerTech', '计算机技术与知识分享'),
(29, 26, 'GeekDIY', '极客精神下的创意DIY');

-- 美食分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(30, 0, 'Food', '美食的制作、食谱和烹饪技巧'),
(31, 30, 'FoodMaking', '美食的制作方法和教程'),
(32, 30, 'FoodDetective', '美食探索与发现'),
(33, 30, 'FoodReviews', '美食的测评和体验'),
(34, 30, 'FoodDiary', '记录日常美食和生活');

-- 动物圈分区及其二级分区
INSERT INTO Category (id, parent, name, description) VALUES
(35, 0, 'Animals', '关于动物的相关内容'),
(36, 35, 'Cats', '关于猫的内容'),
(37, 35, 'Dogs', '关于狗的内容'),
(38, 35, 'ExoticPets', '关于奇特宠物的内容'),
(39, 35, 'WildAnimals', '关于野生动物的内容');

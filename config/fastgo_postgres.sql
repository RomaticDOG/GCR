
-- 删除并重建数据库（PostgreSQL无IF EXISTS + DROP DATABASE的复合语法，需分步骤）
DROP DATABASE IF EXISTS fastgo;
CREATE DATABASE fastgo WITH ENCODING 'UTF8' LC_COLLATE 'en_US.UTF-8' LC_CTYPE 'en_US.UTF-8';

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS post;
-- 1. 先创建表（不带COMMENT）
CREATE TABLE post (
  id BIGSERIAL PRIMARY KEY,
  "userID" VARCHAR(36) NOT NULL DEFAULT '',
  "postID" VARCHAR(35) NOT NULL DEFAULT '',
  title VARCHAR(256) NOT NULL DEFAULT '',
  content TEXT NOT NULL DEFAULT '',
  "createdAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT uk_post_postID UNIQUE ("postID")
);

-- 2. 单独添加字段注释
COMMENT ON COLUMN post."userID" IS '用户唯一 ID';
COMMENT ON COLUMN post."postID" IS '博文唯一 ID';
COMMENT ON COLUMN post.title IS '博文标题';
COMMENT ON COLUMN post.content IS '博文内容';
COMMENT ON COLUMN post."createdAt" IS '博文创建时间';
COMMENT ON COLUMN post."updatedAt" IS '博文最后修改时间';

-- 3. 给表本身添加注释（可选）
COMMENT ON TABLE post IS '博文表';

-- 创建索引（PostgreSQL索引命名更灵活）
CREATE INDEX idx_post_userID ON post ("userID");

-- 添加更新时间自动更新触发器（PostgreSQL无ON UPDATE语法，需触发器实现）
CREATE or replace function update_modified_column()
returns trigger as $$
begin
    new."updatedAt" = CURRENT_TIMESTAMP;
    return new;
end;
$$ language 'plpgsql';

create trigger update_post_updatedAt
before update on post
for each row
execute function update_modified_column();

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS "user";  -- user是关键字，需用双引号包裹
-- 创建用户表（不带字段内COMMENT）
CREATE TABLE "user" (
  id BIGSERIAL PRIMARY KEY,
  "userID" VARCHAR(36) NOT NULL DEFAULT '',
  username VARCHAR(255) NOT NULL DEFAULT '',
  password VARCHAR(255) NOT NULL DEFAULT '',
  nickname VARCHAR(30) NOT NULL DEFAULT '',
  email VARCHAR(256) NOT NULL DEFAULT '',
  phone VARCHAR(16) NOT NULL DEFAULT '',
  "createdAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT uk_user_userID UNIQUE ("userID"),
  CONSTRAINT uk_user_username UNIQUE (username),
  CONSTRAINT uk_user_phone UNIQUE (phone)
);

-- 单独添加字段注释
COMMENT ON COLUMN "user"."userID" IS '用户唯一 ID';
COMMENT ON COLUMN "user".username IS '用户名（唯一）';
COMMENT ON COLUMN "user".password IS '用户密码（加密后）';
COMMENT ON COLUMN "user".nickname IS '用户昵称';
COMMENT ON COLUMN "user".email IS '用户电子邮箱地址';
COMMENT ON COLUMN "user".phone IS '用户手机号';
COMMENT ON COLUMN "user"."createdAt" IS '用户创建时间';
COMMENT ON COLUMN "user"."updatedAt" IS '用户最后修改时间';

-- 表注释（可选）
COMMENT ON TABLE "user" IS '用户表';

-- 添加用户表更新时间触发器
create trigger update_user_updatedAt
before update on "user"
for each row
execute function update_modified_column();

--
-- Dumping data for table `user`
--

INSERT INTO "user" (
  id, "userID", username, password, nickname, 
  email, phone, "createdAt", "updatedAt"
) VALUES (
  96, 'user-000000', 'root', 
  '$2a$10$ctsFXEUAMd7rXXpmccNlO.ZRiYGYz0eOfj8EicPGWqiz64YBBgR1y', 
  'colin404', 'colin404@foxmail.com', '18110000000',
  '2024-12-12 03:55:25'::TIMESTAMPTZ, '2024-12-12 03:55:25'::TIMESTAMPTZ
);

-- 重置自增序列（因为手动插入了id=96，需调整序列起始值）
SELECT setval(pg_get_serial_sequence('"user"', 'id'), (SELECT MAX(id) FROM "user"));
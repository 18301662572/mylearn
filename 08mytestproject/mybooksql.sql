CREATE TABLE `book_category` (
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '书籍分类ID',
`book_category_name` varchar(64) NOT NULL COMMENT '书籍分类名称',
`book_category_no` int NOT NULL COMMENT '书籍分类排序',
`book_category_state` bit NOT NULL DEFAULT 0 COMMENT '书籍分类状态（0：存在 1：删除）',
`createtime` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updatetime` timestamp NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) ,
FULLTEXT INDEX `idx_book_category_name` (`book_category_name`) COMMENT '书籍分类名称索引'
);

CREATE TABLE `book` (
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
`book_name` varchar(64) NOT NULL COMMENT '书籍名称',
`content` longtext NULL COMMENT '书籍内容',
`book_category_id` bigint(20) NOT NULL COMMENT '书籍类型ID',
`create_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
`book_state` bit NOT NULL DEFAULT 0 COMMENT '书籍状态（0：存在 1：删除）',
`book_no` int NOT NULL COMMENT '书籍排序',
PRIMARY KEY (`id`) ,
FULLTEXT INDEX `idx_book_name` (`book_name`) COMMENT '书籍名称索引' ,
FULLTEXT INDEX `idx_content` (`content`) COMMENT '书籍内容索引' ,
INDEX `idx_book_category_id` (`book_category_id`) COMMENT '书籍分类id索引'
);

CREATE TABLE `user` (
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
`user_id` bigint(20) NOT NULL COMMENT '用户ID',
`user_name` varchar(64) NOT NULL COMMENT '登录名称',
`password` varchar(64) NOT NULL COMMENT '密码',
`create_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` timestamp NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
`user_state` bit NOT NULL DEFAULT 0 COMMENT '用户状态（0：存在 1：删除）',
`nick_name` varchar(64) NULL COMMENT '昵称',
PRIMARY KEY (`id`) ,
UNIQUE INDEX `idx_user_id` (`user_id`) COMMENT '用户ID 唯一索引' ,
UNIQUE INDEX `idx_user_name` (`user_name`) COMMENT '登录名称唯一索引'
);


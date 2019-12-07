create table users
(
    id         varchar(100) comment 'ユーザーID',
    email      varchar(255) not null comment 'メールアドレス',
    name       varchar(255) null comment 'ユーザー名',
    image_url  varchar(255) null comment '画像パス',
    created_at timestamp    null comment '作成日',
    updated_at timestamp    null comment '更新日',
    primary key (id)
) default charset = utf8 comment 'ユーザーテーブル';

create index idx_users_id on users (id);

create table attendances
(
    id         int unsigned auto_increment comment '勤怠ID',
    user_id    varchar(100) comment 'ユーザーID',
    kind       tinyint unsigned not null comment '勤怠区分',
    remark     varchar(1000)    null comment '備考',
    created_at timestamp        null comment '作成日',
    updated_at timestamp        null comment '更新日',
    primary key (id)
) default charset = utf8 comment '勤怠テーブル';

create index idx_attendances_id on attendances (id);

alter table attendances
    add foreign key fk_attendances_user_id (user_id) references users (id)

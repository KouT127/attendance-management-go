create table working_hours
(
    id            int unsigned auto_increment comment '勤務時刻ID',
    month         date      not null comment '月',
    working_hours float     not null comment '月の所定労働時間',
    created_at    timestamp null comment '作成日',
    updated_at    timestamp null comment '更新日',
    primary key (id)
) default charset = utf8 comment '月の所定労働時間テーブル';
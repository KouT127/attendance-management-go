create table attendances_time
(
    id         int unsigned auto_increment,
    created_at datetime     null comment '作成日',
    updated_at datetime     null comment '更新日',
    pushed_at  datetime     not null comment '打刻時間',
    remark     varchar(250) null comment '備考',
    constraint attendances_time_pk
        unique (id)
);

create index attendances_time_id_index
    on attendances_time (id);


alter table attendances
    drop column kind;

alter table attendances
    drop column pushed_at;

alter table attendances
    drop column remark;

alter table attendances
    add clocked_in_id int unsigned null comment '出勤時間ID';

alter table attendances
    add clocked_out_id int unsigned null comment '退勤時間ID';

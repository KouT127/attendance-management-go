alter table attendances
    add pushed_at datetime null comment '出勤・退勤時間' after kind;

alter table attendances
    modify created_at datetime null comment '作成日';

alter table attendances
    modify updated_at datetime null comment '更新日';


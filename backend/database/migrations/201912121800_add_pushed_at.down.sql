alter table attendances
    drop column pushed_at;

alter table attendances
    modify created_at timestamp null comment '作成日';

alter table attendances
    modify updated_at timestamp null comment '更新日';


drop table attendances_time;

alter table attendances
    add kind tinyint(3) unsigned;

alter table attendances
    add pushed_at datetime not null;

alter table attendances
    add remark varchar(1000);

alter table attendances
    drop column clocked_in_id;

alter table attendances
    drop column clocked_out_id;

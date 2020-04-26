alter table attendances_time
    add attendance_id int unsigned null;

create index attendances_time_index_attendance_id
    on attendances_time (attendance_id desc);

alter table attendances_time
    add attendance_kind_id int null;

alter table attendances
    drop column clocked_in_id;

alter table attendances
    drop column clocked_out_id;

alter table attendances
    add constraint attendances_foreign_key_user_id_1
        foreign key (user_id) references users (id);
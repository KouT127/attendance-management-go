drop index attendances_time_index_attendance_id on attendances_time;

alter table attendances_time
    drop column attendance_id;

alter table attendances
    add clocked_in_id int unsigned null comment '出勤ID';

alter table attendances
    add clocked_out_id int unsigned null comment '退勤ID';
alter table attendances_time
    add constraint attendances_time_pk
        primary key (id);

alter table attendances_time
    drop key attendances_time_pk;
alter table attendances
    add constraint attendances_ibfk_1 foreign key attendances (user_id) references users (id);
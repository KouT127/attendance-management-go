import React from "react";
import "./AttendanceUserListItem.sass"

export const AttendanceUserListItem = () => {
    return (
        <li className='attendance-user-item'>
            <div className='attendance-user-item__left'>
                <h3 className='attendance-user-item__left-name'>name</h3>
                <p className='attendance-user-item__left-kind'>出勤</p>
            </div>
            <p className='attendance-user-item__right'>0:02:28</p>
        </li>
    );
};

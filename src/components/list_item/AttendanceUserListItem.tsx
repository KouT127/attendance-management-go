import React from "react";
import "./AttendanceUserListItem.sass"

export const AttendanceUserListItem = () => {
    return (
        <li className='attendance-list-item'>
            <div className='attendance-list-item-left'>
                <h3 className='attendance-list-item-left-name'>name</h3>
                <p className='attendance-list-item-left-kind'>出勤</p>
            </div>
            <p className='attendance-list-item-right'>0:02:28</p>
        </li>
    );
};

import React from "react";
import "./AttendanceUserListItem.sass"
import {AttendanceKind} from "../../domains/attendance/attendance";

interface Props {
    name: string
    attendanceKind: AttendanceKind
    submittedAt: string
}

export const AttendanceUserListItem = (props: Props) => {

    return (
        <li className='attendance-user-item'>
            <div className='attendance-user-item__left'>
                <h3 className='attendance-user-item__left-name'>
                    {props.name}
                </h3>
                <p className='attendance-user-item__left-kind'>
                    {AttendanceKind.toString(props.attendanceKind.kind)}
                </p>
            </div>
            <p className='attendance-user-item__right'>
                {props.submittedAt}
            </p>
        </li>
    );
};

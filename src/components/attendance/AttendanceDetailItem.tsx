import React from "react";
import "./AttendanceDetailItem.sass"
import {AttendanceKind} from "../../domains/attendance/attendance";

interface Props {
    name: string
    attendanceKind: AttendanceKind
    submittedAt: string
}

export const AttendanceDetailItem = (props: Props) => {

    return (
        <li className='attendance-detail-item'>
            <div className='attendance-detail-item__left'>
                <h3 className='attendance-detail-item__left-name'>
                    {props.name}
                </h3>
                <p className='attendance-detail-item__left-kind'>
                    {AttendanceKind.toString(props.attendanceKind.kind)}
                </p>
            </div>
            <p className='attendance-detail-item__right'>
                {props.submittedAt}
            </p>
        </li>
    );
};

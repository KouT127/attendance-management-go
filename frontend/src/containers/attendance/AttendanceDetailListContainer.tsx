import React from "react";
import {Attendance, AttendanceKind} from "../../domains/attendance/Attendance";
import {AttendanceDetailItem} from "../../components/attendance/AttendanceDetailItem";
import {useUserSelector} from "../../hooks/auth";

type Props = {
    attendances: Array<Attendance>
}

export const AttendanceDetailListContainer = (props: Props) => {
    const {user} = useUserSelector();
    if (!props.attendances) {
        return <div>Loading...</div>
    }
    return (
        <ol className='attendance-list'>
            {
                props.attendances.map((doc, index) => {
                    const data = doc;
                    const attendance: Attendance = {
                        userId: data.userId,
                        kind: data.kind,
                        remark: data.remark,
                        createdAt: data.createdAt,
                        updatedAt: data.updatedAt,
                    };
                    return (
                        <AttendanceDetailItem
                            key={'attendance-user-list-item' + index}
                            name={user.username || 'No Name'}
                            attendanceKind={new AttendanceKind(attendance.kind)}
                            submittedAt={attendance.createdAt || ''}
                        />
                    )
                })
            }
        </ol>
    )
};

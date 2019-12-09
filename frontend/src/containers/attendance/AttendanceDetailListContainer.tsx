import React from "react";
import {UserState} from "../../redux/states/UserState";
import {Attendance, AttendanceKind} from "../../domains/attendance/attendance";
import * as moment from "moment";
import {AttendanceDetailItem} from "../../components/attendance/AttendanceDetailItem";

type Props = {
    user: UserState
    attendances: Array<Attendance>
}

export const AttendanceDetailListContainer = (props: Props) => {
    if (!props.attendances){
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
                    const timestamp = attendance && attendance.createdAt;
                    const unix = timestamp && timestamp.seconds;
                    const formattedTime = unix === null || unix === undefined ?
                        '' : moment.unix(unix).format('YYYY/MM/DD hh:mm:ss');
                    return (
                        <AttendanceDetailItem
                            key={'attendance-user-list-item' + index}
                            name={props.user.username || 'username'}
                            attendanceKind={new AttendanceKind(attendance.kind)}
                            submittedAt={formattedTime}
                        />
                    )
                })
            }
        </ol>
    )
};

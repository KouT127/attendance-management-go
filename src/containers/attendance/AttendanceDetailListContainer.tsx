import React from "react";
import {IUserState} from "../../redux/states/UserState";
import {AttendanceKind, Attendance} from "../../domains/attendance/attendance";
import * as firebase from "firebase";
import * as moment from "moment";
import {AttendanceDetailItem} from "../../components/attendance/AttendanceDetailItem";

type Props = {
    user: IUserState
    attendances: firebase.firestore.QueryDocumentSnapshot[]
}

export const AttendanceDetailListContainer = (props: Props) => {
    return (
        <ol className='attendance-list'>
            {
                props.attendances.map((doc, index) => {
                    const data = doc.data();
                    const attendance: Attendance = {
                        type: data.type,
                        content: data.content,
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
                            name={props.user.name || 'name'}
                            attendanceKind={new AttendanceKind(attendance.type)}
                            submittedAt={formattedTime}
                        />
                    )
                })
            }
        </ol>
    )
};

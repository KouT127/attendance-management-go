import React, {useEffect} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInformationHeader";
import {AttendanceKind, IAttendance} from "../../domains/attendance/attendance";
import * as firebase from "firebase";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {useUserSelector} from "../../hooks/auth";
import {IUserState} from "../../redux/states/UserState";
import moment from "moment";
import {AttendanceFormContainer} from "../../components/form/AttendanceForm";
import {useAttendanceDocuments} from "../../hooks/firestore";

export const AttendanceUser = () => {
    console.log('AttendanceUser render');
    const {user} = useUserSelector();
    const {documents, observeAttendance} = useAttendanceDocuments();
    useEffect(() => {
        observeAttendance();
    }, []);
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <AttendanceFormContainer
                documents={documents}
            />
            <AttendanceUserList user={user} attendances={documents}/>
        </div>
    );
};

type ListProps = {
    user: IUserState
    attendances: firebase.firestore.QueryDocumentSnapshot[]
}

const AttendanceUserList = (props: ListProps) => {
    console.log('AttendanceUserList render');
    console.log(props.user);
    return <ol className='attendance-list'>
        {props.attendances.map((doc, index) => {
            const data = doc.data();
            const attendance: IAttendance = {
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
                <AttendanceUserListItem
                    key={'attendance-user-list-item' + index}
                    name={props.user.name || 'name'}
                    attendanceKind={new AttendanceKind(attendance.type)}
                    submittedAt={formattedTime}
                />
            )
        })}
    </ol>
};

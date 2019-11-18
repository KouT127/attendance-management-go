import React, {useCallback, useEffect, useState} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInformationHeader";
import {TimerSection} from "../../components/section/TimerSection";
import {AttendanceKind, IAttendance} from "../../domains/attendance/model";
import {firebaseApp} from "../../lib/firebase";
import * as firebase from "firebase";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {useUserSelector} from "../../hooks/auth";
import {IUserState} from "../../redux/states/UserState";
import moment from "moment";
import {AttendanceForm} from "../../components/form/AttendanceForm";

export const useAttendanceDocuments = () => {
    const [documents, setDocuments] = useState<firebase.firestore.QueryDocumentSnapshot[]>([]);

    const observeAttendance = useCallback(async () => {
        firebaseApp
            .firestore()
            .collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances')
            .orderBy('createdAt', 'desc')
            .limit(5)
            .onSnapshot((snapshot) => {
                const documents = snapshot.docs;
                setDocuments(documents)
            });

    }, []);

    return {
        observeAttendance,
        documents
    }
};

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
            <AttendanceForm
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
            return (<AttendanceUserListItem
                key={'attendance-user-list-item' + index}
                name={props.user.name || 'No Name'}
                attendanceKind={new AttendanceKind(attendance.type)}
                submittedAt={formattedTime}
            />)
        })}
    </ol>
};

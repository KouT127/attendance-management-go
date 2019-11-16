import React, {useCallback, useEffect, useState} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInformationHeader";
import {TimerSection} from "../../components/section/TimerSection";
import {AttendanceKind, IAttendance} from "../../domains/attendance/model";
import {firebaseApp} from "../../lib/firebase";
import * as firebase from "firebase";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {useAuthUser} from "../../hooks/auth";
import {IUserState} from "../../redux/states/UserState";

const useAttendanceDocuments = () => {
    const [documents, setDocuments] = useState<firebase.firestore.QueryDocumentSnapshot[]>([]);

    const getAttendance = useCallback(async () => {
        const snapshot = await firebaseApp
            .firestore()
            .collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances').get();
        const documents = snapshot.docs;
        setDocuments(documents)
    }, []);

    return {
        getAttendance,
        documents
    }
};

export const AttendanceUser = () => {
    const {user} = useAuthUser();
    const {documents, getAttendance} = useAttendanceDocuments();
    useEffect(() => {
        getAttendance();
    }, []);
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <TimerSection/>
            <AttendanceUserList user={user} attendances={documents}/>
        </div>
    );
};

type ListProps = {
    user: IUserState
    attendances: firebase.firestore.QueryDocumentSnapshot[]
}

const AttendanceUserList = (props: ListProps) => {
    return <ol className='attendance-list'>
        {props.attendances.map((doc, index) => {
            const data = doc.data();
            const attendance: IAttendance = {
                type: data.type,
                content: data.content,
                createdAt: data.createdAt,
                updatedAt: data.updatedAt,
            };
            return (<AttendanceUserListItem
                key={'attendance-user-list-item' + index}
                name={props.user.name || 'No Name'}
                attendanceKind={new AttendanceKind(attendance.type)}
                submittedAt={''}
            />)
        })}
    </ol>
};

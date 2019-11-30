import React, {useEffect} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/attendance/AttendanceUserInformationHeader";
import {useUserSelector} from "../../hooks/auth";
import {useAttendanceDocuments} from "../../hooks/firestore";
import {AttendanceFormContainer} from "../../containers/attendance/AttendanceFormContainer";
import {AttendanceDetailListContainer} from "../../containers/attendance/AttendanceDetailListContainer";

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
            <AttendanceDetailListContainer
                user={user}
                attendances={documents}/>
        </div>
    );
};

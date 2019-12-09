import React, {useEffect} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/attendance/AttendanceUserInformationHeader";
import {useUserSelector} from "../../hooks/auth";
import {useAttendanceDocuments} from "../../hooks/firestore";
import {AttendanceFormContainer} from "../../containers/attendance/AttendanceFormContainer";
import {AttendanceDetailListContainer} from "../../containers/attendance/AttendanceDetailListContainer";

export const AttendanceUser = () => {
    const {user} = useUserSelector();
    const {attendances, fetchAttendance} = useAttendanceDocuments();
    useEffect(() => {
        fetchAttendance();
    }, []);
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <AttendanceFormContainer
                attendances={attendances}
            />
            <AttendanceDetailListContainer
                user={user}
                attendances={attendances}/>
        </div>
    );
};

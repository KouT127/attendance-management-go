import React, {useEffect} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/attendance/AttendanceUserInformationHeader";
import {useAttendance} from "../../hooks/xhr";
import {AttendanceFormContainer} from "../../containers/attendance/AttendanceFormContainer";
import {AttendanceDetailListContainer} from "../../containers/attendance/AttendanceDetailListContainer";

export const AttendanceUser = () => {
    const {attendances, fetchAttendance} = useAttendance();
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
                attendances={attendances}/>
        </div>
    );
};

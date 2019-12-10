import React, {useEffect} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/attendance/AttendanceUserInformationHeader";
import {AttendanceContext, useAttendance} from "../../hooks/xhr";
import {AttendanceFormContainer} from "../../containers/attendance/AttendanceFormContainer";
import {AttendanceDetailListContainer} from "../../containers/attendance/AttendanceDetailListContainer";

export const AttendanceUser = () => {
    const {attendances, fetchAttendance} = useAttendance();
    useEffect(() => {
        fetchAttendance();
    }, []);
    return (
        <AttendanceContext.Provider value={{attendances: attendances}}>
            <div className='attendance'>
                <AttendanceUserInformationHeader/>
                <AttendanceFormContainer/>
                <AttendanceDetailListContainer/>
            </div>
        </AttendanceContext.Provider>
    );
};

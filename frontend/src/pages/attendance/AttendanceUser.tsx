import React from "react";
import "./AttendanceUser.sass";
import {AttendanceUserInformationHeader} from "../../components/attendance/AttendanceUserInformationHeader";

import {AttendanceFormContainer} from "../../containers/attendance/AttendanceFormContainer";
import {AttendanceDetailListContainer} from "../../containers/attendance/AttendanceDetailListContainer";
import {AttendanceProvider} from "../../hooks/attendance";

export const AttendanceUser = () => {
  return (
    <AttendanceProvider>
      <div className="attendance">
        <AttendanceUserInformationHeader />
        <AttendanceFormContainer />
        <AttendanceDetailListContainer />
      </div>
    </AttendanceProvider>
  );
};

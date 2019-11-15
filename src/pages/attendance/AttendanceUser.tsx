import React, {FC} from "react";
import "./AttendanceUser.sass";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInfomationHeader";
import {TimerSection} from "../../components/section/TimerSection";
import {AttendanceKind, AttendanceKindEnum} from "../../domains/attendance/model";

export const AttendanceUser: FC = () => {
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <TimerSection/>
            <ol className='attendance-list'>
                <AttendanceUserListItem
                    name={'kou'}
                    attendanceKind={new AttendanceKind(AttendanceKindEnum.GO_TO_WORK)}
                    submittedAt={'2019/10/10'}
                />
                <AttendanceUserListItem
                    name={'kou'}
                    attendanceKind={new AttendanceKind(AttendanceKindEnum.LEAVE_WORK)}
                    submittedAt={'2019/10/10'}
                />
                <AttendanceUserListItem
                    name={'kou'}
                    attendanceKind={new AttendanceKind(AttendanceKindEnum.GO_TO_WORK)}
                    submittedAt={'2019/10/10'}
                />
            </ol>
        </div>
    );
};

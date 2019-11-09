import React from "react";
import "./AttendanceUser.sass";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInfomationHeader";
import {TimerSection} from "../../components/section/TimerSection";

export const AttendanceUser = () => {
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <TimerSection/>
            <ol className='attendance-list'>
                <AttendanceUserListItem/>
            </ol>
        </div>
    );
};

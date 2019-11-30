import React from "react";
import {RoundedButton} from "../common/RoundedButton";
import {AttendanceTimerContainer} from "../../containers/attendance/AttendanceTimerContainer";

type Props = {
    buttonTitle: string
    onChangeTextArea: (event: React.ChangeEvent<HTMLTextAreaElement>) => void
    onClickButton: () => void
};

export const AttendanceForm = (props: Props) => {
    return (
        <section className='timer-section'>
            <AttendanceTimerContainer/>
            <textarea
                name='content'
                className='timer-section__textarea'
                onChange={props.onChangeTextArea}/>
            <RoundedButton
                title={props.buttonTitle}
                appearance={"black"}
                onClick={props.onClickButton}/>
        </section>
    )
};

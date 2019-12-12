import React from "react";
import {AttendanceTimerContainer} from "../../containers/attendance/AttendanceTimerContainer";
import {RoundedSubmitButton} from "../common/RoundedSubmitButton";

type Props = {
    buttonTitle: string
    register: any
    onClickButton: (e: React.BaseSyntheticEvent<object, any, any>) => Promise<void>
};

export const AttendanceForm = (props: Props) => {
    return (
        <section className='timer-section'>
            <form onSubmit={props.onClickButton}>
                <AttendanceTimerContainer/>
                <textarea
                    name='remark'
                    className='timer-section__textarea'
                    ref={props.register({
                        max_length: 100,
                    })}
                />
                <RoundedSubmitButton className={'timer-section__button-section'} title={props.buttonTitle}/>
            </form>
        </section>
    )
};

import React from "react";
import {RoundedButton} from "../common/RoundedButton";
import {AttendanceTimerContainer} from "../../containers/attendance/AttendanceTimerContainer";
import {OnSubmit} from "react-hook-form/dist/types";
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
                    name='content'
                    className='timer-section__textarea'
                    ref={props.register({
                        required: 'Required',
                        max_length: 100,
                    })}
                />
                <RoundedSubmitButton className={'timer-section__button-section'} title={props.buttonTitle}/>
            </form>
        </section>
    )
};

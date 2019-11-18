import React, {useState} from "react";
import {AttendanceKindEnum, IAttendance} from "../../domains/attendance/model";
import {db} from "../../lib/firebase";
import {RoundedButton} from "../button/RoundedButton";
import * as firebase from "firebase";
import {TimerSection} from "../section/TimerSection";

export const AttendanceForm = () => {
    const [attendance, setAttendance] = useState<IAttendance>({
        type: AttendanceKindEnum.GO_TO_WORK,
        content: '',
        createdAt: undefined,
        updatedAt: undefined,
    });
    // const handleChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     const target = event.target;
    //     const value = target.type === 'checkbox' ? target.checked : target.value;
    //     setAttendance({
    //         ...attendance,
    //         [event.target.name]: event.target.value
    //     });
    // };

    const handleChangeTextareaText = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        const target = event.target;
        setAttendance({
            ...attendance,
            [target.name]: target.value
        });
    };

    const addAttendance = () => {
        db.collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances')
            .add({
                ...attendance,
                createdAt: firebase.firestore.FieldValue.serverTimestamp(),
                updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
            })
    };

    return (
        <section className='timer-section'>
            <TimerSection/>
            <textarea
                name='content'
                className='timer-section__textarea'
                onChange={handleChangeTextareaText}/>
            <RoundedButton
                title={"出勤する"}
                appearance={"black"}
                onClick={addAttendance}/>
        </section>
    )
};

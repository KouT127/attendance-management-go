import React, {useEffect, useState} from "react";
import {AttendanceKindEnum, IAttendance} from "../../domains/attendance/model";
import {db} from "../../lib/firebase";
import {RoundedButton} from "../button/RoundedButton";
import * as firebase from "firebase";
import {TimerSection} from "../section/TimerSection";

type Props = {
    documents: firebase.firestore.QueryDocumentSnapshot[]
}

export const AttendanceFormContainer = (props: Props) => {
    const [title, setTitle] = useState('');
    const [attendance, setAttendance] = useState<IAttendance>({
        type: AttendanceKindEnum.GO_TO_WORK,
        content: '',
        createdAt: undefined,
        updatedAt: undefined,
    });

    useEffect(() => {
        const latestAttendance = props.documents.length > 0 ? props.documents[0] : undefined;
        const latestKindType: AttendanceKindEnum = latestAttendance && latestAttendance.data().type;
        const kindType = latestKindType === AttendanceKindEnum.GO_TO_WORK ? AttendanceKindEnum.LEAVE_WORK : AttendanceKindEnum.GO_TO_WORK
        const buttonTitle = kindType === AttendanceKindEnum.GO_TO_WORK ? '出勤する' : '退勤する';
        setTitle(buttonTitle);
        setAttendance({
            type: kindType,
        })
    }, [props.documents]);

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
        AttendanceForm({
            buttonTitle: title,
            onClickButton: addAttendance,
            onChangeTextArea: handleChangeTextareaText
        })
    )
};
type AttendanceFormProps = {
    buttonTitle: string
    onChangeTextArea: (event: React.ChangeEvent<HTMLTextAreaElement>) => void
    onClickButton: () => void
}

export const AttendanceForm = (props: AttendanceFormProps) => {


    // const handleChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     const target = event.target;
    //     const value = target.type === 'checkbox' ? target.checked : target.value;
    //     setAttendance({
    //         ...attendance,
    //         [event.target.name]: event.target.value
    //     });
    // };

    return (
        <section className='timer-section'>
            <TimerSection/>
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

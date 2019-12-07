import React, {useCallback, useEffect, useState} from "react";
import {AttendanceKindEnum, Attendance} from "../../domains/attendance/attendance";
import {db} from "../../lib/firebase";
import * as firebase from "firebase";
import {AttendanceForm} from "../../components/attendance/AttendanceForm";
import {useUserSelector} from "../../hooks/auth";
import useForm from "react-hook-form";

type Props = {
    documents: firebase.firestore.QueryDocumentSnapshot[]
}

export const AttendanceFormContainer = (props: Props) => {
    const [title, setTitle] = useState('');
    const { handleSubmit, register, errors, reset } = useForm();

    const {user} = useUserSelector();
    const [attendance, setAttendance] = useState<Attendance>({
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


    const handleClickButton = useCallback(() => {
        db.collection('users')
            .doc(user.id)
            .collection('attendances')
            .add({
                ...attendance,
                createdAt: firebase.firestore.FieldValue.serverTimestamp(),
                updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
            });
        reset();
    }, []);

    return (
        AttendanceForm({
            buttonTitle: title,
            register: register,
            onClickButton: handleSubmit(handleClickButton),
        })
    )
};

import React, {useCallback, useContext, useEffect, useState} from "react";
import {Attendance, AttendanceKindEnum} from "../../domains/attendance/Attendance";
import {firebaseApp} from "../../lib/firebase";
import {AttendanceForm} from "../../components/attendance/AttendanceForm";
import {useUserSelector} from "../../hooks/auth";
import useForm from "react-hook-form";
import axios from "axios";
import {AttendanceContext} from "../../hooks/xhr";

export const AttendanceFormContainer = () => {
  const [title, setTitle] = useState("");
  const { handleSubmit, register, errors, reset } = useForm();
  const { attendances } = useContext(AttendanceContext);

  const { user } = useUserSelector();
  const [attendance, setAttendance] = useState<Attendance>({
    userId: user.id,
    kind: AttendanceKindEnum.GO_TO_WORK,
    remark: "",
    createdAt: undefined,
    updatedAt: undefined
  });

  useEffect(() => {
    const latestAttendance =
      attendances.length > 0 ? attendances[0] : undefined;
    const latestKindType: AttendanceKindEnum =
      (latestAttendance && latestAttendance.kind) ||
      AttendanceKindEnum.GO_TO_WORK;
    const kindType =
      latestKindType === AttendanceKindEnum.GO_TO_WORK
        ? AttendanceKindEnum.LEAVE_WORK
        : AttendanceKindEnum.GO_TO_WORK;
    const buttonTitle =
      kindType === AttendanceKindEnum.GO_TO_WORK ? "出勤する" : "退勤する";
    setTitle(buttonTitle);
    setAttendance({
      ...attendance,
      kind: kindType
    });
  }, [attendances]);

  const handleClickButton = useCallback(async () => {
    const currentUser = firebaseApp.auth().currentUser;
    if (!currentUser) {
      return;
    }
    const token = await currentUser.getIdToken();
    const response = await axios.post(
      `http://localhost:8080/v1/attendances`,
      {
        ...attendance
      },
      {
        headers: { authorization: token }
      }
    );
    reset();
  }, []);

  return AttendanceForm({
    buttonTitle: title,
    register: register,
    onClickButton: handleSubmit(handleClickButton)
  });
};

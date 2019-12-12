import React, {useCallback, useEffect, useState} from "react";
import {Attendance, AttendanceKindEnum} from "../../domains/attendance/Attendance";
import {AttendanceForm} from "../../components/attendance/AttendanceForm";
import {useUserSelector} from "../../hooks/auth";
import useForm from "react-hook-form";
import {useAttendanceStore} from "../../hooks/attendance";

export const AttendanceFormContainer = () => {
  const [buttonTitle, setButtonTitle] = useState("");
  const { handleSubmit, register, errors, reset } = useForm();
  const { createAttendance, latestKindType } = useAttendanceStore();
  const { user } = useUserSelector();
  const [attendance, setAttendance] = useState<Attendance>({
    userId: user.id,
    kind: AttendanceKindEnum.GO_TO_WORK,
    remark: "",
    createdAt: undefined,
    updatedAt: undefined
  });

  useEffect(() => {
    const buttonTitle =
      latestKindType === AttendanceKindEnum.GO_TO_WORK
        ? "出勤する"
        : "退勤する";
    setButtonTitle(buttonTitle);
    setAttendance({
      ...attendance,
      kind: latestKindType
    });
  }, [latestKindType]);

  const handleClickButton = useCallback(
    async (value: any) => {
      setAttendance({
        ...attendance,
        remark: value.remark || ""
      });
      await createAttendance(attendance);
      reset();
    },
    [attendance]
  );

  return AttendanceForm({
    buttonTitle: buttonTitle,
    register: register,
    onClickButton: handleSubmit(handleClickButton)
  });
};

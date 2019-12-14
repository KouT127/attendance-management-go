import React, {useCallback, useEffect, useState} from "react";
import {AttendanceKindEnum} from "../../domains/attendance/Attendance";
import {AttendanceForm} from "../../components/attendance/AttendanceForm";
import {useAuth, useUserSelector} from "../../hooks/auth";
import useForm from "react-hook-form";
import {useAttendanceStore} from "../../hooks/attendance";

export const AttendanceFormContainer = () => {
  const [buttonTitle, setButtonTitle] = useState("");
  const { handleSubmit, register, errors, reset } = useForm();
  const { attendances, setCreateEvent, latestKindType } = useAttendanceStore();
  const { user } = useUserSelector();
  const { getToken } = useAuth();

  useEffect(() => {
    const buttonTitle =
      latestKindType === AttendanceKindEnum.GO_TO_WORK
        ? "出勤する"
        : "退勤する";
    setButtonTitle(buttonTitle);
  }, [latestKindType]);

  const handleClickButton = useCallback(
    async (value: any) => {
      const token = await getToken();
      await setCreateEvent({
        token: token,
        userId: user.id,
        remark: value.remark,
        kind: latestKindType
      });
      reset();
    },
    [latestKindType]
  );

  return AttendanceForm({
    buttonTitle: buttonTitle,
    register: register,
    onClickButton: handleSubmit(handleClickButton)
  });
};

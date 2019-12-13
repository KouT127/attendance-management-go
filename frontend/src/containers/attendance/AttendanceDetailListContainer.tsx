import React, {useEffect} from "react";
import {Attendance, AttendanceKind} from "../../domains/attendance/Attendance";
import {AttendanceDetailItem} from "../../components/attendance/AttendanceDetailItem";
import {useAuth, useUserSelector} from "../../hooks/auth";
import {useAttendanceStore} from "../../hooks/attendance";

export const AttendanceDetailListContainer = () => {
  const { user } = useUserSelector();
  const { getToken } = useAuth();
  const { attendances, fetchAttendance } = useAttendanceStore();
  useEffect(() => {
    const fetch = async () => {
      const token = await getToken();
      fetchAttendance(token);
    };
    fetch();
  }, []);

  if (!attendances) {
    return <div>Loading...</div>;
  }

  return (
    <ol className="attendance-list">
      {attendances.map((doc, index) => {
        const data = doc;
        const attendance: Attendance = {
          userId: data.userId,
          kind: data.kind,
          remark: data.remark,
          createdAt: data.createdAt,
          updatedAt: data.updatedAt
        };
        return (
          <AttendanceDetailItem
            key={"attendance-user-list-item" + index}
            name={user.name || "No Name"}
            attendanceKind={new AttendanceKind(attendance.kind)}
            submittedAt={attendance.createdAt || ""}
          />
        );
      })}
    </ol>
  );
};

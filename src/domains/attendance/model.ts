import * as firebase from "firebase";

interface IDocument {
    createdAt?: firebase.firestore.FieldValue
    updatedAt?: firebase.firestore.FieldValue
}

export enum AttendanceKindEnum {
    GO_TO_WORK = 10,
    LEAVE_WORK = 20,
}

export class AttendanceKind {
    kind: AttendanceKindEnum;

    constructor(kind: AttendanceKindEnum) {
        this.kind = kind
    }

    public static toString(kind: AttendanceKindEnum) {
        switch (kind) {
            case AttendanceKindEnum.GO_TO_WORK:
                return '出勤';
            case AttendanceKindEnum.LEAVE_WORK:
                return '退勤';
            default:
                return '';
        }
    }
}


export interface IAttendance extends IDocument {
    type: AttendanceKindEnum
    content?: string
}

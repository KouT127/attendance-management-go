import * as firebase from "firebase";

interface IDocument {
    createdAt?: firebase.firestore.FieldValue
    updatedAt?: firebase.firestore.FieldValue
}

export enum AttendanceType {
    GO_TO_WORK = 10,
    LEAVE_WORK = 20,
}

export interface IAttendance extends IDocument {
    type: AttendanceType
    content?: string
}

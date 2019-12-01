import firebase from "firebase";

export interface Document {
    createdAt?: firebase.firestore.Timestamp
    updatedAt?: firebase.firestore.Timestamp
}

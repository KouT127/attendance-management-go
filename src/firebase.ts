import { firebaseConfig } from "./lib/config";
import * as firebase from "firebase";

export const firebaseApp = firebase.initializeApp(firebaseConfig);
export const firestoreApp = firebaseApp.firestore();

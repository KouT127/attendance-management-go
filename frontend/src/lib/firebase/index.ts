import {firebaseConfig} from './config';
import * as firebase from "firebase";

export const firebaseApp = firebase.initializeApp(firebaseConfig);
export const db = firebaseApp.firestore();

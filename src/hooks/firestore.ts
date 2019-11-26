import {useCallback, useState} from "react";
import {firebaseApp} from "../lib/firebase";
import {useUserSelector} from "./auth";

export const useAttendanceDocuments = () => {
    const [documents, setDocuments] = useState<firebase.firestore.QueryDocumentSnapshot[]>([]);

    const observeAttendance = useCallback(async () => {
        firebaseApp
            .firestore()
            .collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances')
            .orderBy('createdAt', 'desc')
            .limit(5)
            .onSnapshot((snapshot) => {
                const documents = snapshot.docs;
                setDocuments(documents)
            });

    }, []);

    return {
        observeAttendance,
        documents
    }
};

export const useUserDocuments = () => {
    const {user} = useUserSelector();
    const [documents, setDocuments] = useState<firebase.firestore.QueryDocumentSnapshot[]>([]);

    const setUserDocument = useCallback(async (name: string) => {
        console.log(name)
        await firebaseApp
            .firestore()
            .collection('users')
            .doc(user.id)
            .set({
                name: name
            })
    }, []);

    return {
        setUserDocument,
        documents
    }
};

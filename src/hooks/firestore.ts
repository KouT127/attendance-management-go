import {useCallback, useState} from "react";
import {firebaseApp} from "../lib/firebase";

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

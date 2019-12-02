import {useCallback, useEffect, useState} from "react";
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
    const [userDocumentRef, setUserDocumentRef] = useState<firebase.firestore.DocumentReference | undefined>();

    useEffect(() => {
        const reference = firebaseApp
            .firestore()
            .collection('users')
            .doc(user.id);
        setUserDocumentRef(reference)
    }, []);

    const setUserDocument = useCallback(async (name: string) => {
        if (!userDocumentRef) {
            return;
        }
        await userDocumentRef
            .update({
                username: name
            })
    }, [userDocumentRef]);

    return {
        userDocumentRef,
        setUserDocument,
        documents
    }
};

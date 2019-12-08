import {useCallback, useState} from "react";
import {firebaseApp} from "../lib/firebase";
import {useUserSelector} from "./auth";
import axios from "axios";
import {useDispatch} from "react-redux";
import {actionCreator} from "../store";

export const useAttendanceDocuments = () => {
    const {user} = useUserSelector();
    const [documents, setDocuments] = useState<firebase.firestore.QueryDocumentSnapshot[]>([]);

    const observeAttendance = useCallback(async () => {
        firebaseApp
            .firestore()
            .collection('users')
            .doc(user.id)
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
    const dispatch = useDispatch();

    const setUserData = useCallback(async (name: string) => {
        const currentUser = firebaseApp.auth().currentUser;
        if (!currentUser) {
            return
        }
        const token = await currentUser.getIdToken();
        console.log(name, user.email);
        const response = await axios.put(
            `http://localhost:8080/v1/users/${currentUser.uid}`,

            {name: name, email: user.email, imageUrl: user.imageUrl,},
            {headers: {'authorization': token}});
        const userData = response.data.user;
        dispatch(actionCreator.userActionCreator.loadedUser({initialLoaded: true, userState: {...userData,}}))
    }, []);

    return {
        setUserData
    }
};

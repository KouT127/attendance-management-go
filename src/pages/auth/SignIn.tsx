import React, {FC, useEffect, useState} from "react";
import {RoundedButton} from "../../components/button/RoundedButton";
import "./SignIn.sass";
import "../../styles/base.sass";
import {firebaseApp} from "../../lib/firebase";
import {useUserSelector} from "../../hooks/auth";
import {useHistory} from "react-router";


const signUp = async (email: string, password: string) => {
    try {
        const context = await firebaseApp.auth().signInWithEmailAndPassword(email, password);
        console.log(context);
        return context.user;
    } catch (e) {
        console.error('auth', e)
    }
};

interface IInputInformation {
    email?: string
    password?: string
}

export const SignIn: FC = () => {
    const history = useHistory();
    const {user, isAuthenticated} = useUserSelector();
    const [input, setInput] = useState<IInputInformation>({});
    useEffect(() => {
        if (!isAuthenticated) {
            return;
        }
        history.replace('/home');
    }, [user]);


    const handleChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInput({
            ...input,
            [event.target.name]: event.target.value,
        });
    };
    const handleSubmit = () => {
        if (input.email === undefined || input.password === undefined) {
            return;
        }
        signUp(input.email, input.password)
    };
    return (
        <>
            <div className='container'>
                <div className='container-header'>
                    <h2 className='container-header-title'>
                        新規ユーザー登録
                    </h2>
                </div>
                <div className='container-input-text'>
                    <label className='container-input-label'>メールアドレス</label>
                    <div className='text'>
                        <input className='container-input-text-area'
                               name='email'
                               type='email'
                               onChange={handleChangeInput}/>
                    </div>
                </div>
                <div className='container-input-text'>
                    <label className='container-input-label'>password</label>
                    <div className='text'>
                        <input className='container-input-text-area'
                               name='password'
                               type='password'
                               onChange={handleChangeInput}/>
                    </div>
                </div>
                <div className='container-input-button-section'>
                    <RoundedButton
                        title='登録する'
                        appearance='navy'
                        onClick={handleSubmit}/>
                </div>
            </div>
        </>
    );
};

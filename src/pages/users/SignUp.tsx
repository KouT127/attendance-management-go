import React, {useState} from "react";
// @ts-ignore
import {RoundedButton} from "../../components/button/RoundedButton";
import "./SignUp.sass";
import "../../base/base.sass";
import firebase from "firebase";


const signUp = async (email: string, password: string) => {
    try {
        const context = await firebase.auth().signInWithEmailAndPassword(email, password);
        console.log(context)
        return context.user;
    } catch (e) {
        console.error('auth', e)
    }

};

interface IInputInformation {
    email?: string
    password?: string
}

export const SignUp = () => {
    const [input, setInput] = useState<IInputInformation>({});
    const handleChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInput({
            ...input,
            [event.target.name]: event.target.value,
        });
    };
    const handleSubmit = () => {
        if (input.email == undefined || input.password == undefined) {
            return;
        }
        signUp(input.email, input.password)
    };
    return (
        <div>
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
        </div>
    );
};

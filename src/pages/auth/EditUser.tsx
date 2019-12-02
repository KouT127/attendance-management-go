import React, {useCallback, useState} from "react"
import './EditUser.sass'
import {useUserDocuments} from "../../hooks/firestore";
import useForm from "react-hook-form";
import { useHistory } from 'react-router-dom'

type FormData = {
    username: string
}

export const EditUser = () => {
    const {handleSubmit, setValue, register, errors} = useForm<FormData>();
    const {setUserDocument} = useUserDocuments();
    const history = useHistory();

    const onSubmit = handleSubmit(async ({username}) => {
        await setUserDocument(username);
        history.push('/home')
    });

    console.log(errors)
    return (
        <div>
            <div>

            </div>
            <form onSubmit={onSubmit}>
                <section className='edit-user__header'>
                    <h1 className='edit-user__header-title'>ユーザー作成</h1>
                </section>
                <div className='edit-user__form-section'>
                    <label>ユーザー名</label>
                    <input className='edit-user__text-input'
                           type={'text'}
                           name={'username'}
                           ref={register({ required: true, maxLength: 50 })}/>
                    {errors.username && <p className='edit-user__error-message'>ユーザー名は必須です</p>}
                </div>

                <input className='edit-user__button'
                       type='submit'
                       name='enter'
                       value='登録'/>
            </form>
        </div>
    )
};

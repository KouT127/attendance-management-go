import React from "react";
import {FieldError, FieldName} from "react-hook-form/dist/types";
import {CreateFormData} from "../../pages/auth/CreateUser";

type Props<T> = {
    register: {
        (validateRule: Partial<any>): (Ref: HTMLInputElement) => void,
    },
    errors: Partial<Record<FieldName<T>, FieldError>>,
}

export const CreateUserForm = (props: Props<CreateFormData>) => {
    return (
        <>
            <section className='create-user__header'>
                <h1 className='create-user__header-title'>ユーザー作成</h1>
            </section>
            <div className='create-user__form-section'>
                <label>ユーザー名</label>
                <input className='create-user__text-input'
                       type={'text'}
                       name={'name'}
                       ref={props.register({required: true, maxLength: 50})}/>
                {props.errors.name && <p className='create-user__error-message'>ユーザー名は必須です</p>}
            </div>

            <input className='create-user__button'
                   type='submit'
                   name='enter'
                   value='登録'/>
        </>
    )
};

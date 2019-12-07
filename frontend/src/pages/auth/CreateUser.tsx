import React from "react"
import './CreateUser.sass'
import {useUserDocuments} from "../../hooks/firestore";
import useForm from "react-hook-form";
import {useHistory} from 'react-router-dom'
import {CreateUserForm} from "../../containers/user/CreateUserForm";

export type CreateFormData = {
    username: string
}

export const CreateUser = () => {
    const {handleSubmit, register, errors} = useForm<CreateFormData>();
    const {setUserDocument} = useUserDocuments();
    const history = useHistory();
    const onSubmit = handleSubmit(async ({username}) => {
        await setUserDocument(username);
        history.push('/home')
    });
    return (
        <>
            <form onSubmit={onSubmit}>
                <CreateUserForm
                    register={register}
                    errors={errors}/>
            </form>
        </>
    )
};

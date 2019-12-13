import React from "react"
import './CreateUser.sass'

import useForm from "react-hook-form";
import {useHistory} from 'react-router-dom'
import {CreateUserForm} from "../../containers/user/CreateUserForm";
import {useUserDetail} from "../../hooks/user";

export type CreateFormData = {
    name: string
}

export const CreateUser = () => {
    const {handleSubmit, register, errors} = useForm<CreateFormData>();
    const {setUserData} = useUserDetail();
    const history = useHistory();
    const onSubmit = handleSubmit(async ({name}) => {
        await setUserData(name);
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

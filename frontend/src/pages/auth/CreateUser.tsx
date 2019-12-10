import React from "react"
import './CreateUser.sass'
import {useUserDetail} from "../../hooks/xhr";
import useForm from "react-hook-form";
import {useHistory} from 'react-router-dom'
import {CreateUserForm} from "../../containers/user/CreateUserForm";

export type CreateFormData = {
    username: string
}

export const CreateUser = () => {
    const {handleSubmit, register, errors} = useForm<CreateFormData>();
    const {setUserData} = useUserDetail();
    const history = useHistory();
    const onSubmit = handleSubmit(async ({username}) => {
        await setUserData(username);
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

import React, {Component, Suspense} from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {SignUp} from "./pages/users/SignUp";
import {DefaultLayout} from "./layouts/DefaultLayout";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useDispatch} from "react-redux";
import {actionCreator} from "./store";
import {useAuthUser} from "./hooks/auth";
import {Header} from "./components/header/Header";

;

interface AuthParameters {
    isAuthenticated: boolean;
    children: any
}

const Auth = (props: AuthParameters) => {
    return props.isAuthenticated ? props.children : <Redirect to={'/users/new'}/>
};

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Routes = () => {
    const {isAuthenticated} = useAuthUser();
    // const user = DefaultLayout(AttendanceUser)
    return (
        <>
            <Header title='Time'/>
            <Router>
                <div className='contents'>
                    <Switch>
                        <UsersRoute/>
                        <Auth isAuthenticated={isAuthenticated}>
                            <Route path="/attendance" component={AttendanceUser}/>
                            <Route path="/attendance/scan" exact component={AttendanceScan}/>
                        </Auth>
                    </Switch>
                </div>
            </Router>
        </>
    );
};

const UsersRoute = () => {

    return (
        <>
            <Route path="/users/new" exact component={SignUp}/>
        </>
    );
};


export default App;

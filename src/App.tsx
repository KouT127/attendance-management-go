import React, {Component, useEffect} from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {SignIn} from "./pages/auth/SignIn";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useDispatch} from "react-redux";
import {actionCreator} from "./store";
import {useAuthUser} from "./hooks/auth";
import {Header} from "./components/header/Header";
import {useApplication} from "./hooks/application";
import {PulseLoader} from "react-spinners";
import {Splash} from "./pages/common/Splash";
import {NotFound} from "./pages/common/NotFound";


export type HeaderProps = {
    children: any
}
export type Props = {
    component: () => object
    path: string
}

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Auth: React.FC<HeaderProps> = (props) => {
    const {isAuthenticated} = useAuthUser();
    return isAuthenticated ? props.children : <Redirect to={'/signin'}/>
};

const ProtectedRoute = (props: Props) => {
    const component = props.component();
    return (
        <Route exact path={props.path} render={() => {
            return (
                <Auth>
                    {component}
                </Auth>
            )
        }}/>
    );
};
export const Loading = (props: HeaderProps) => {
    const {initialLoaded} = useApplication();
    const {isAuthenticated} = useAuthUser();
    useEffect(() => {
        if (!initialLoaded) {
            return;
        }
    }, [initialLoaded, isAuthenticated]);
    if (!initialLoaded) {
        return (
            <div className='initial-loading__section'>
                <PulseLoader
                    sizeUnit={"px"}
                    size={30}
                    color={'black'}
                    loading={true}
                />
            </div>
        );
    }
    return props.children
};

const Routes = () => {
    return (
        <>
            <Header title='Time'/>
            <main className={'contents'}>
                <Router>
                    <Loading>
                        <Switch>
                            <Route exact path="/" component={Splash}/>
                            <Route exact path="/signin" component={SignIn}/>
                            <ProtectedRoute component={AttendanceUser} path="/home"/>
                            <ProtectedRoute component={AttendanceScan} path="/scan"/>
                            <Route path={'*'} component={NotFound}/>
                        </Switch>
                    </Loading>
                </Router>
            </main>
        </>
    );
};

export default App;

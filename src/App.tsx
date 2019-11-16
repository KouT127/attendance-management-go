import React, {useEffect} from "react";
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

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Auth: React.FC<HeaderProps> = (props) => {
    const {isAuthenticated} = useAuthUser();
    return isAuthenticated ? props.children : <Redirect to={'/signin'}/>
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
                    color={'#123abc'}
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
                    <Switch>
                        <Loading>
                            <Auth>
                                <Route exact path="/home" component={AttendanceUser}/>
                                <Route exact path="/scan" component={AttendanceScan}/>
                            </Auth>
                            <AuthRoute/>
                        </Loading>
                    </Switch>
                </Router>
            </main>
        </>
    );
};

const AuthRoute = () => {
    return (
        <>
            <Route exact path="/" component={Splash}/>
            <Route exact path="/signin" component={SignIn}/>
            <Route component={NotFound}/>
        </>
    );
};


export default App;

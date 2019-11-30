import React, {Component, Suspense, useEffect} from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useAuth, useUserSelector} from "./hooks/auth";
import {Header} from "./components/common/Header";
import {useApplication} from "./hooks/application";
import {PulseLoader} from "react-spinners";
import {NotFound} from "./pages/common/NotFound";
import {Splash} from "./pages/common/Splash";
import {EditUser} from "./pages/auth/EditUser";

const SignIn = React.lazy(() => import('./pages/auth/SignIn')
    .then(importedModule => ({
        default: importedModule.SignIn
    }))
);


const AttendanceUser = React.lazy(() => import("./pages/attendance/AttendanceUser")
    .then(importedModule => ({
        default: importedModule.AttendanceUser
    }))
);

export type HeaderProps = {
    children: any
}
export type Props = {
    component: React.ComponentType<any>
    path: string
}

const App: React.FC = () => {
    const {observeAuth} = useAuth();
    observeAuth();
    return (<Routes/>)
};

const Auth: React.FC<HeaderProps> = (props) => {
    const {isAuthenticated} = useUserSelector();
    return isAuthenticated ? props.children : <Redirect to={'/signin'}/>
};

interface IRouteProps {
    exact?: boolean;
    path: string;
    component: React.ComponentType<any>;
}

const ProtectedRoute = ({component: Component, ...otherProps}: IRouteProps) => {
    return (
        <Route render={() => {
            return (
                <Auth>
                    <Component {...otherProps} />
                </Auth>
            )
        }}/>
    );
};
export const Loading = (props: HeaderProps) => {
    const {initialLoaded} = useApplication();
    const {isAuthenticated} = useUserSelector();
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
                <Suspense fallback={<div>Loading</div>}>
                    <Router>
                        <Loading>
                            <Switch>
                                <Route exact path="/" component={Splash}/>
                                <Route exact path="/signin" component={SignIn}/>
                                <Route exact path="/users/new" component={EditUser}/>
                                <ProtectedRoute component={AttendanceUser} path="/home"/>
                                <ProtectedRoute component={AttendanceScan} path="/scan"/>
                                <Route path={'*'} component={NotFound}/>
                            </Switch>
                        </Loading>
                    </Router>
                </Suspense>
            </main>
        </>
    );
};

export default App;

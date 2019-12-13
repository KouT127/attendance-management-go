import React, {Component, Suspense, useEffect} from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";
import {useAuth, useUserSelector} from "./hooks/auth";
import {Header} from "./components/common/Header";
import {useApplication} from "./hooks/application";
import {PulseLoader} from "react-spinners";
import {Splash} from "./pages/common/Splash";


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

const CreateUser = React.lazy(() => import("./pages/auth/CreateUser")
    .then(importedModule => ({
        default: importedModule.CreateUser
    }))
);

const NotFound = React.lazy(() => import('./pages/common/NotFound')
    .then(importedModule => ({
        default: importedModule.NotFound
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
    const {subscribeAuth} = useAuth();
    subscribeAuth();
    return (<Routes/>)
};

const Auth: React.FC<HeaderProps> = (props) => {
    const {isAuthenticated, shouldEdit} = useUserSelector();
    if (shouldEdit) {
        return <Redirect to={'/users/new'}/>;
    }
    if (!isAuthenticated) {
        return <Redirect to={'/signin'}/>
    }
    return props.children
};

interface RouteProps {
    exact?: boolean;
    path: string;
    component: React.ComponentType<any>;
}

const ProtectedRoute = ({component: Component, ...otherProps}: RouteProps) => {
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
                <Suspense fallback={<div/>}>
                    <Router>
                        <Loading>
                            <Switch>
                                <Route exact path="/" component={Splash}/>
                                <Route exact path="/signin" component={SignIn}/>
                                <Route exact path="/users/new" component={CreateUser}/>
                                <ProtectedRoute path="/home" component={AttendanceUser}/>
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

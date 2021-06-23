import './App.css';
import { Component, React } from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import NavigationWidget from './components/navigation/NavigationWidget'
import KiosksPage from "./components/pages/kiosks/KiosksPage";
import ImagesPage from "./components/pages/images/ImagesPage";
import UsersPage from "./components/pages/users/UsersPage";
import StatisticsPage from "./components/pages/statistics/StatisticsPage";
import LoginPage from "./components/pages/login/LoginPage";
import { JWT_TOKEN_KEY } from './Constants'

import logo from './components/icons/800x800.png'


class App extends Component {

    render() {
        if (localStorage.getItem(JWT_TOKEN_KEY)) {
            return (
                <BrowserRouter>
                    <div className="BaseContainer">
                        <div className='LogoDiv ChildElem'>
                            <img src={logo} alt="kiosk logo" />
                            <h1>Kiosk Management System</h1>
                        </div>
                        <NavigationWidget />
                    </div>
                    <Switch>
                        <Route exact path="/kiosks" component={KiosksPage} />
                        <Route exact path="/images" component={ImagesPage} />
                        <Route exact path={"/users"} component={UsersPage} />
                        <Route path="/" component={StatisticsPage} />
                    </Switch>

                </BrowserRouter>
            );
        } else {
            return (
                <LoginPage />
            )
        }
    }
}

export default App;
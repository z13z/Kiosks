import './App.css';
import {Component, React} from 'react'
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import NavigationWidget from './components/navigation/NavigationWidget'
import KiosksPage from "./components/pages/kiosks/KiosksPage";
import ImagesPage from "./components/pages/images/ImagesPage";
import UsersPage from "./components/pages/users/UsersPage";
import StatisticsPage from "./components/pages/statistics/StatisticsPage";
import LoginPage from "./components/pages/login/LoginPage";
import {JWT_TOKEN_KEY} from './Constants'


class App extends Component {

    render() {
        if (localStorage.getItem(JWT_TOKEN_KEY)) {
            return (
                <BrowserRouter>
                    <div style={{display: "flex"}}>
                        <NavigationWidget/>
                        <Switch>
                            <Route exact path="/kiosks" component={KiosksPage}/>
                            <Route exact path="/images" component={ImagesPage}/>
                            <Route exact path={"/users"} component={UsersPage}/>
                            <Route path="/" component={StatisticsPage}/>
                        </Switch>
                    </div>
                </BrowserRouter>
            );
        } else {
            return (
                <LoginPage/>
            )
        }
    }
}

export default App;
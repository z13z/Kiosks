import './App.css';
import {React, Component} from 'react'
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import NavigationWidget from './components/navigation/NavigationWidget'
import KiosksPage from "./components/pages/kiosks/KiosksPage";
import ImagesPage from "./components/pages/images/ImagesPage";
import UsersPage from "./components/pages/users/UsersPage";
import StatisticsPage from "./components/pages/statistics/StatisticsPage";

class App extends Component {

    render() {
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
    }
}

export default App;

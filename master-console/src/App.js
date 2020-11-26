import './App.css';
import {React, Component} from 'react'
import {BrowserRouter, Route} from 'react-router-dom'
import NavigationWidget from './components/navigation/NavigationWidget'
import KiosksPage from "./components/pages/kiosks/KiosksPage";
import ImagesPage from "./components/pages/images/ImagesPage";

class App extends Component {

    render() {
        return (
            <BrowserRouter>
                <div style={{display: "flex"}}>
                    <NavigationWidget/>
                    <Route exact path="/kiosks" component={KiosksPage}/>
                    <Route exact path="/images" component={ImagesPage}/>
                </div>
            </BrowserRouter>
        );
    }

}

export default App;

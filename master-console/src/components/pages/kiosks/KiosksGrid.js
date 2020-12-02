import React, {Component} from 'react'
import './KiosksPage.css'
import axios from "axios"

class kiosksGrid extends Component {
    state = {
        kiosks: [
            {
                name: "test kiosk 1"
            }
        ]
    }

    componentDidMount() {
        axios.get("/kiosks/list").then(response => {
            this.setState(response)
        })
    }

    render() {
        return <div className="KiosksGridDiv">
            <ul>
                {this.state.kiosks.map(kiosk => <li>{kiosk.name}</li>)}
            </ul>
        </div>
    }
}

export default kiosksGrid
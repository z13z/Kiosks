import React, {Component} from 'react'
import './KiosksPage.css'
import axios from "axios"

class kiosksGrid extends Component {
    state = {
        kiosks: []
    }

    componentDidMount() {
        this.loadKiosks()
    }

    shouldComponentUpdate(nextProps, nextState, nextContext) {
        return Object.is(this.state, nextState)
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        this.loadKiosks()
    }

    render() {
        return <div className="KiosksGridDiv">
            <ul>
                {this.state.kiosks.map(kiosk => <li key={kiosk.id}>{kiosk.name}</li>)}
            </ul>
        </div>
    }

    loadKiosks() {
        let queryParams = {params: {}}
        if (this.props.kioskId != null && this.props.kioskId !== "") {
            queryParams.params.id = this.props.kioskId
        }
        if (this.props.kioskName != null && this.props.kioskName !== "") {
            queryParams.params.name = this.props.kioskName
        }
        axios.get("/kiosk", queryParams).then(response => {
            console.log(response.data)
            this.setState({
                kiosks: response.data
            })
        })
    }
}

export default kiosksGrid
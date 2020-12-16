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

    stateDidChange(props, state) {
        return JSON.stringify(this.state) !== JSON.stringify(state) || JSON.stringify(this.props) !== JSON.stringify(props)
    }

    shouldComponentUpdate(nextProps, nextState, nextContext) {
        return this.stateDidChange(nextProps, nextState)
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.stateDidChange(prevProps, prevState)) {
            this.loadKiosks();
        }
    }

    render() {
        return <div className="KiosksGridDiv">
            <ul>
                {this.state.kiosks != null ? this.state.kiosks.map(kiosk => <li
                    key={kiosk.id}>{kiosk.name}</li>) : null}
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
            this.setState({
                kiosks: response.data
            })
        })
    }
}

export default kiosksGrid
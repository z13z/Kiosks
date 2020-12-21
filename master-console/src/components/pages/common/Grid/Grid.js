import React, {Component} from 'react'
import './Grid.css'
import axios from "axios";

class Grid extends Component {
    state = {
        data: []
    }

    componentDidMount() {
        this.loadData()
    }

    stateDidChange(props, state) {
        return JSON.stringify(this.state) !== JSON.stringify(state) || JSON.stringify(this.props) !== JSON.stringify(props)
    }

    shouldComponentUpdate(nextProps, nextState, nextContext) {
        return this.stateDidChange(nextProps, nextState)
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.stateDidChange(prevProps, prevState)) {
            this.loadData();
        }
    }


    render() {
        const allColumns = this.getColumns()
        const header = allColumns.map(col => <th key={col} className="GridHeader">{col}</th>)
        return <table className="GridTable">
            <thead>
            <tr className="GridHeader">{header}</tr>
            </thead>
            <tbody>{this.state.data != null ? this.state.data.map(row => <tr className="GridRow"
                                                                             key={row.id}>{(allColumns.map(col =>
                <td key={row.id + '_' + col} className="GridColumn">{row[col]}</td>))}</tr>) : null}
            </tbody>
        </table>
    }


    //abstract methods
    getColumns() {
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }

    loadData() {
        let queryParams = {params: {}}
        this.getSearchProps().forEach(prop =>{
            queryParams[prop] = this.props[prop]
        })
        // noinspection JSCheckFunctionSignatures
        axios.get(this.getQueryUrl(), queryParams).then(response => {
            this.setState({
                data: response.data
            })
        })
    }

    getSearchProps(){
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }

    getQueryUrl(){
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }
}

export default Grid
import React, {Component} from 'react'
import './Grid.css'
import axios from "axios";
import './Pagination'
import GridPagination from "./Pagination";
import {JWT_TOKEN_KEY} from '../../../../Constants'

const ROWS_COUNT_ON_PAGE = 10

class Grid extends Component {

    state = {
        data: [],
        currentPage: 1,
        pagesCount: 1
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
        return (
            <div>
                <div className="GridTableDiv">
                    <table className="GridTable">
                        <thead>
                        <tr className="GridHeader">{header}</tr>
                        </thead>
                        <tbody>{this.state.data != null ? this.state.data.map(row => <tr className="GridRow"
                                                                                         key={row.id}>{(allColumns.map(col =>
                            <td key={row.id + '_' + col} className="GridColumn">{row[col]}</td>))}</tr>) : null}
                        </tbody>
                    </table>
                </div>
                <GridPagination currentPage={this.state.currentPage} pagesCount={this.state.pagesCount}
                                changePage={this.changePage.bind(this)}/>
            </div>
        )
    }

    changePage(page) {
        if (page > 0 && page <= this.state.pagesCount) {
            this.setState({
                currentPage: page
            });
        }
    }

    //abstract methods
    getColumns() {
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }

    loadData() {
        const axiosHeader = {Authorization: localStorage.getItem(JWT_TOKEN_KEY)}
        let queryParams = {params: {}}
        this.getSearchProps().forEach(prop => {
            queryParams[prop] = this.props[prop]
        })
        queryParams['offset'] = (this.state.currentPage - 1) * ROWS_COUNT_ON_PAGE
        queryParams['limit'] = ROWS_COUNT_ON_PAGE
        // noinspection JSCheckFunctionSignatures
        axios.get(this.getQueryUrl(), {params: queryParams, header: axiosHeader}).then(response => {
            this.setState({
                data: response.data.rows,
                pagesCount: Math.ceil(response.data.rowsCount / ROWS_COUNT_ON_PAGE)
            })
        })
    }

    getSearchProps() {
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }

    getQueryUrl() {
        console.error("getColumns isn't implemented in Grid")
        throw new TypeError("getColumns isn't implemented in Grid")
    }
}

export default Grid
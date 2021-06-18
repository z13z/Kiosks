import './ImagesPage.css'
import Grid from "../common/Grid/Grid";
import axios from "axios";
import {JWT_TOKEN_KEY} from "../../../Constants";
import {Button} from "reactstrap";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faEdit, faPlay, faTrash} from "@fortawesome/free-solid-svg-icons";

class ImagesGrid extends Grid {

    getColumns() {
        return ["id", "name", "state"]
    }

    getActionColumnsHeader() {
        return ["build", "edit", "delete"]
    }

    getSearchProps() {
        return ["id", "name"]
    }

    getQueryUrl() {
        return "/image"
    }

    buildImageAction(row) {
        let queryParams = {}
        queryParams['id'] = row.id
        queryParams['name'] = row.name
        queryParams['script'] = row.script
        queryParams['state'] = "waiting"

        axios.post('/image', queryParams, {headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)}}).then(() => {
            this.props.successfullyUpdated()
        }).catch(error => {
            if (error.response.status === 401) {
                localStorage.removeItem(JWT_TOKEN_KEY)
                window.location.reload();
            } else if (error.response.status === 403) {
                alert("action is forbidden")
            } else if (error.response.status === 400) {
                alert("image can't be built")
            } else {
                throw error;
            }
            this.props.onClose()
        })
    }

    deleteAction(rowId) {
        axios.delete(this.getQueryUrl(), {
            headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)},
            data: {rowId}
        }).then(() => {
            this.props.successfullyUpdated()
        }).catch(error => {
            if (error.response !== undefined) {
                if (error.response.status === 401) {
                    localStorage.removeItem(JWT_TOKEN_KEY)
                    window.location.reload();
                } else if (error.response.status === 403) {
                    alert("action is forbidden")
                } else if (error.response.status === 400) {
                    alert("bar request")
                } else {
                    throw error;
                }
            } else {
                throw error;
            }
        })
    }

    getActionColumns(row) {
        return (
            <>
                <td key={row.id + "_build"} className="GridColumn" style={{width: '41px'}}>
                    <Button style={row.state !== 'created' ? {opacity: '30%'} : null}
                            disabled={row.state !== 'created'}>
                        <FontAwesomeIcon icon={faPlay}
                                         onClick={() => this.buildImageAction(row)}/>
                    </Button>
                </td>
                <td key={row.id + "_edit"} className="GridColumn" style={{width: '41px'}}>
                    <Button>
                        <FontAwesomeIcon icon={faEdit}
                                         onClick={() => this.props.editImageAction(row.id, row.name, row.script)}/>
                    </Button>
                </td>
                <td key={row.id + "_delete"} className="GridColumn" style={{width: '41px'}}>
                    <Button onClick={() => this.deleteAction(row.id)}>
                        <FontAwesomeIcon icon={faTrash}/>
                    </Button>
                </td>
            </>
        )
    }
}

export default ImagesGrid
import './UsersPage.css'
import Grid from "../common/Grid/Grid";
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faEdit, faTrash} from '@fortawesome/free-solid-svg-icons'
import {Button} from "reactstrap";
import axios from "axios";
import {JWT_TOKEN_KEY} from "../../../Constants";

class UsersGrid extends Grid {

    getColumns() {
        return ["id", "name", "permissions"]
    }

    getActionColumnsHeader() {
        return ["edit", "delete"]
    }

    getSearchProps() {
        return ["id", "name"]
    }

    getQueryUrl() {
        return "/users"
    }

    deleteAction(rowId) {
        axios.delete(this.getQueryUrl(), {
            headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)},
            data: {rowId}
        }).then(() => {
            this.props.successfullyUpdated()
        }).catch(error => {
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
        })
    }

    getActionColumns(row) {
        return (
            <>
                <td key={row.id + "_edit"} className="GridColumn" style={{width: '41px'}}>
                    <Button>
                        <FontAwesomeIcon icon={faEdit}
                                         onClick={() => this.props.editUserAction(row.id, row.name, row.permissions)}/>
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

export default UsersGrid
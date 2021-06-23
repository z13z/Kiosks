import Grid from "../common/Grid/Grid";
import {Button} from "reactstrap";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faFileImage, faPlayCircle} from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import {JWT_TOKEN_KEY} from "../../../Constants";
import fileDownload from "js-file-download";

class KiosksGrid extends Grid {

    getColumns() {
        return ["id", "address", "kioskImageId", "lastOnline"]
    }

    getActionColumnsHeader() {
        return ["screenshot", "command"]
    }


    getSearchProps() {
        return ["id", "kioskAddress"]
    }

    getQueryUrl() {
        return "/kiosk"
    }

    // used code from https://github.com/kennethjiang/js-file-download
    getScreenshot(kioskId) {
        axios.get('/kiosksCommander', {
            params: {id: kioskId},
            headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)},
            responseType: 'blob',
        }).then(res => {
            fileDownload(res.data, "screenshot.png");
        }).catch(error => {
            if (error.response.status === 401) {
                localStorage.removeItem(JWT_TOKEN_KEY)
                window.location.reload();
            } else if (error.response.status === 400) {
                alert("Bad Request, " + error.response.data)
            } else if (error.response.status === 403) {
                alert("action is forbidden")
            } else if (error.response.status === 503) {
                alert("Kiosk is inactive")
            } else {
                throw error;
            }
        });
    }

    getActionColumns(row) {
        return (
            <>
                <td key={row.id + "_screenshot"} className="GridColumn" style={{width: '41px'}}>
                    <Button>
                        <FontAwesomeIcon icon={faFileImage}
                                         onClick={() => this.getScreenshot(row.id)}/>
                    </Button>
                </td>
                <td key={row.id + "_command"} className="GridColumn" style={{width: '41px'}}>
                    <Button onClick={() => this.props.sendCommandAction(row.id)}>
                        <FontAwesomeIcon icon={faPlayCircle}/>
                    </Button>
                </td>
            </>
        )
    }
}

export default KiosksGrid
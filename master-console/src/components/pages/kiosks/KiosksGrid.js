import './KiosksPage.css'
import Grid from "../common/Grid/Grid";

class KiosksGrid extends Grid {

    getColumns() {
        return ["id", "name"]
    }

    getSearchProps() {
        return ["id", "name"]
    }

    getQueryUrl() {
        return "/kiosk"
    }
}

export default KiosksGrid
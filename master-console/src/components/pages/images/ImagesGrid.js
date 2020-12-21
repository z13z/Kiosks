import './ImagesPage.css'
import Grid from "../common/Grid/Grid";

class ImagesGrid extends Grid {

    getColumns() {
        return ["id", "name"]
    }

    getSearchProps(){
        return ["id", "name"]
    }

    getQueryUrl(){
        return "/image"
    }
}

export default ImagesGrid
import './UsersPage.css'
import Grid from "../common/Grid/Grid";

class UsersGrid extends Grid {

    getColumns() {
        return ["id", "name"]
    }

    getSearchProps(){
        return ["id", "name"]
    }

    getQueryUrl(){
        return "/users"
    }
}

export default UsersGrid
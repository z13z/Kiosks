import './UsersPage.css'
import Grid from "../common/Grid/Grid";

class UsersGrid extends Grid {

    getColumns() {
        return ["id", "name", "permissions"]
    }

    getSearchProps(){
        return ["id", "name"]
    }

    getQueryUrl(){
        return "/users"
    }
}

export default UsersGrid
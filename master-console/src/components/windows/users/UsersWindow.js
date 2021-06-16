import {React} from 'react'
import PopUpWindow from '../common/PopUpWindow'
import {Button, FormGroup, Input, Label} from 'reactstrap';
import {ALL_USER_PERMISSIONS} from '../../../Constants'

const usersWindow = (props) => {
    const fieldValues = {
        username: props.username !== undefined ? props.username : "",
        password: "",
        repassword: "",
        permissions: []
    };
    return (
        <PopUpWindow {...props}>
            <FormGroup key='usernameGroupKey'>
                <Label for="usernameField">Username</Label>
                <Input type="text" name="username" id="usernameField" value={fieldValues.username} required/>
            </FormGroup>
            <FormGroup key='passwordFieldKey'>
                <Label for="passwordField">Password</Label>
                <Input type="password" name="password" id="passwordField" value={fieldValues.password}/>
            </FormGroup>
            <FormGroup key='rePasswordFieldKey'>
                <Label for="rePasswordField">Reenter Password</Label>
                <Input type="rePassword" name="rePassword" id="rePasswordField" value={fieldValues.password}/>
            </FormGroup>
            <FormGroup key='permissionsFieldKey'>
                <Label for="permissionsField">Password</Label>
                <Input type={'select'} name='permissions' id='permissionsField' multiple>
                    {ALL_USER_PERMISSIONS.map((permissionName,) => {
                        return <option
                            value={permissionName}
                            selected={fieldValues.permissions.includes(permissionName)}>{permissionName}</option>
                    })
                    }
                </Input>
            </FormGroup>
            <Button variant="primary" type="submit" key="submitUserFromField">
                Submit
            </Button>
        </PopUpWindow>
    )
}

export default usersWindow;
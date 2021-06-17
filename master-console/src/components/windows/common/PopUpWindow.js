import React from 'react'
import Modal from 'react-modal';
import {Button} from "reactstrap";

Modal.setAppElement('#root')

const popUpWindow = (props) => {
    const popUpWindowStyle = {
        overlay: {
            backgroundColor: 'lightsteelblue'
        },
        content: {
            color: 'blue'
        }
    }

    const contentDivStyle = {
        marginLeft: 'auto',
        marginRight: 'auto',
        width: '40%'
    }

    const fieldStyle = {
        float: 'right',
        paddingBottom: '20px'
    }

    return (
        <Modal style={popUpWindowStyle}
               isOpen={true}>
            <div style={contentDivStyle}>
                {props.children.map((field,) => {
                    return (<div style={fieldStyle}>{field}</div>)
                })}
            </div>
            <Button color="success" onClick={props.onSubmit}>
                Submit
            </Button>
            <Button color="warning" onClick={props.onClose}>
                close
            </Button>
        </Modal>
    )
}

export default popUpWindow
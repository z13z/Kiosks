import React from 'react'
import {Button, Modal, ModalBody} from 'reactstrap';

const popUpWindow = (props) => {
    return (
        <Modal
            {...props}
            size="lg"
            aria-labelledby="contained-modal-title-vcenter"
            centered>
            <Modal.Header closeButton>
                <Modal.Title id="contained-modal-title-vcenter">
                    Modal heading
                </Modal.Title>
            </Modal.Header>
            <ModalBody>
                <div>
                    {props.children.map((field,) => {
                        return ({field})
                    })}
                </div>
            </ModalBody>
            <Modal.Footer>
                <Button onClick={() => {
                }}>Close</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default popUpWindow
import React from 'react';
import {Modal} from 'react-bootstrap';

class BambooModal extends React.PureComponent {
    static propTypes = {
        show: PropTypes.bool.isRequired,
    };
    render(){
        const { show } = this.props;
        const component = (
            <div>
                Bamboo HR employees
            </div>
        );
        return (
            <Modal
                dialogClassName='modal--scroll'
                show={show}
                onHide={this.handleClose}
                onExited={this.handleClose}
                bsSize='large'
                backdrop='static'
            >
                <Modal.Header closeButton={true}>
                    <Modal.Title>
                        {'Bamboo HR'}
                    </Modal.Title>
                </Modal.Header>
                <form
                    role='form'
                    onSubmit={this.handleCreate}
                >
                    <Modal.Body
                        ref='modalBody'
                    >
                        {component}
                    </Modal.Body>
                    <Modal.Footer>
                        Footer
                    </Modal.Footer>
                </form>
            </Modal>
        )
    }
}

export default BambooModal;
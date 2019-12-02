import React from 'react';
import {Modal} from 'react-bootstrap';

class BambooModal extends React.PureComponent {
    static propTypes = {
        show: PropTypes.bool.isRequired,
        employees: PropTypes.array.isRequired,
        getEmployees: PropTypes.func.isRequired,
    };

    componentDidMount() {
        const {getEmployees} = this.props;
        getEmployees();
    }

    render() {
        const {show, employees} = this.props;

        const renderEmployees = employees.map((employee) => <div>{employee.location}</div>);
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
                        {renderEmployees}
                    </Modal.Body>
                    <Modal.Footer>
                        Footer
                    </Modal.Footer>
                </form>
            </Modal>
        );
    }
}

export default BambooModal;
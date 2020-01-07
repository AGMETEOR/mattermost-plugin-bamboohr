import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {createEmployee} from '../../actions';

import BambooAction from './bamboo_action';

function mapStateToProps(state, ownProps) {
    const user = ownProps.user;
    return {
        user,
    };
}

const mapDispatchToProps = (dispatch) => bindActionCreators({
    createEmployee,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(BambooAction);
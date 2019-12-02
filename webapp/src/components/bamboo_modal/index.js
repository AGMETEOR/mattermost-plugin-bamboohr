import React from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {id as pluginId} from 'manifest';
import {getEmployees} from '../../actions';

import BambooModal from './bamboo_modal';

function mapStateToProps(state) {
    return {

        // Just get the modal to work at the moment
        show: true,
        employees: state[`plugins-${pluginId}`].employees,
    };
}

const mapDispatchToProps = (dispatch) => bindActionCreators({
    getEmployees,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(BambooModal);


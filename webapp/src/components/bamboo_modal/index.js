import React from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import BambooModal from './bamboo_modal';

function mapStateToProps() {
    return {
        // Just get the modal to work at the moment
        show: true,
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators({}, dispatch),
    };
}


export default connect(mapStateToProps, mapDispatchToProps)(BambooModal);


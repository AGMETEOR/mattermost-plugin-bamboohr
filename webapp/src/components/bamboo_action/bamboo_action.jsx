import React from 'react';

import PropTypes from 'prop-types';

class BambooAction extends React.PureComponent {
    static propTypes = {
        user: PropTypes.object.isRequired,
        createEmployee: PropTypes.func.isRequired,
    };

    render() {
        const {user, createEmployee} = this.props;
        // eslint-disable-next-line camelcase
        const {first_name, last_name} = user;
        return (
            <div
                data-toggle='tooltip'
                className='popover__row first'
            >
                <a
                    href='#'
                    className='text-nowrap'
                >
                    <button
                        className='style--none'
                        role='menuitem'
                        onClick={() => createEmployee({firstName: first_name, lastName: last_name})}
                    >
                        <i className='fa fa-bold'/>
                        <span>
                            {'Create user in Bamboo'}
                        </span>
                    </button>
                </a>
            </div>
        );
    }
}

export default BambooAction;
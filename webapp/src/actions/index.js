import Client from '../client';
import ActionTypes from '../actionTypes';
import Constants from '../constants';

import BambooModal from '../components/bamboo_modal';

export function getEmployees() {
    return async (dispatch) => {
        let data;
        try {
            data = await Client.getEmployees();
        } catch (error) {
            return {error};
        }

        dispatch({
            type: ActionTypes.RECEIVED_EMPLOYEES,
            data,
        });

        return {data};
    };
}

function openModal(modalData) {
    return (dispatch) => {
        const action = {
            type: 'MODAL_OPEN',
            modalId: modalData.modalId,
            dialogProps: modalData.dialogProps,
            dialogType: modalData.dialogType,
        };

        dispatch(action);
    };
}

export function openBambooModal(elements) {
    return (dispatch) => {
        const bambooModalData = {
            ModalId: Constants.BAMBOO_MODAL,
            dialogType: BambooModal,
            dialogProps: {
                elements,
            },
        };

        dispatch(openModal(bambooModalData));
    };
}

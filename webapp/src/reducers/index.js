import {combineReducers} from 'redux';
import ActionTypes from '../actionTypes';

function getEmployees(state = [], action) {
    switch (action.type) {
        case ActionTypes.RECEIVED_EMPLOYEES:
            return action.data;
        default:
            return state;
        }
}

export default combineReducers({
    getEmployees,
});
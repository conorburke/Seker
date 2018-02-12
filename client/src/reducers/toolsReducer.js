import { FETCH_TOOLS } from '../actions/types';

export default function(state = [], action) {
    // console.log(action);
    switch (action.type) {
        case FETCH_TOOLS:
            return action.payload || false;
        default:
            return state
    }
}